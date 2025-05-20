package sovereign

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/data/sovereign"
	"github.com/multiversx/mx-chain-core-go/data/sovereign/dto"
	"github.com/multiversx/mx-chain-core-go/marshal"
	testscommon "github.com/multiversx/mx-chain-core-go/testscommon/sovereign"
)

func TestNewHeadersNotifier(t *testing.T) {
	t.Parallel()

	t.Run("nil marshaller should fail", func(t *testing.T) {
		hn, err := NewHeadersNotifier(nil, &mock.HasherMock{})
		require.Equal(t, core.ErrNilMarshalizer, err)
		require.Nil(t, hn)
	})

	t.Run("nil hasher should fail", func(t *testing.T) {
		hn, err := NewHeadersNotifier(&marshal.GogoProtoMarshalizer{}, nil)
		require.Equal(t, core.ErrNilHasher, err)
		require.Nil(t, hn)
	})

	t.Run("valid arguments should succeed", func(t *testing.T) {
		hn, err := NewHeadersNotifier(&marshal.GogoProtoMarshalizer{}, &mock.HasherMock{})
		require.NoError(t, err)
		require.NotNil(t, hn)
		require.Empty(t, hn.subscribers)
	})
}

func TestHeadersNotifier_RegisterSubscriber(t *testing.T) {
	t.Parallel()

	hn, _ := NewHeadersNotifier(&marshal.GogoProtoMarshalizer{}, &mock.HasherMock{})

	t.Run("nil subscriber should fail", func(t *testing.T) {
		err := hn.RegisterSubscriber(nil)
		require.Equal(t, errNilHeaderSubscriber, err)
		require.Empty(t, hn.subscribers)
	})

	t.Run("multiple subscribers should work", func(t *testing.T) {
		hn.subscribers = nil // Reset
		sub1 := &testscommon.HeaderSubscriberMock{}
		sub2 := &testscommon.HeaderSubscriberMock{}
		err := hn.RegisterSubscriber(sub1)
		require.NoError(t, err)
		err = hn.RegisterSubscriber(sub2)
		require.NoError(t, err)
		require.Len(t, hn.subscribers, 2)
		require.Equal(t, []IncomingHeaderSubscriber{sub1, sub2}, hn.subscribers)
	})

	t.Run("concurrent register should be safe", func(t *testing.T) {
		hn.subscribers = nil // Reset
		var wg sync.WaitGroup
		numSubscribers := 100
		wg.Add(numSubscribers)

		for i := 0; i < numSubscribers; i++ {
			go func() {
				defer wg.Done()
				err := hn.RegisterSubscriber(&testscommon.HeaderSubscriberMock{})
				if err != nil {
					t.Error(err)
				}
			}()
		}

		wg.Wait()
		require.Len(t, hn.subscribers, numSubscribers)
	})
}

func TestHeadersNotifier_NotifyHeaderSubscribers(t *testing.T) {
	t.Parallel()

	marshaller := &marshal.GogoProtoMarshalizer{}
	hasher := &mock.HasherMock{}
	hn, _ := NewHeadersNotifier(marshaller, hasher)

	header := &sovereign.IncomingHeader{SourceChainID: dto.MVX, Nonce: 123}
	headerHash, _ := core.CalculateHash(marshaller, hasher, header)

	t.Run("nil header should fail", func(t *testing.T) {
		err := hn.NotifyHeaderSubscribers(nil)
		require.Equal(t, errNilIncomingHeader, err)
	})

	t.Run("no subscribers should succeed", func(t *testing.T) {
		hn.subscribers = nil // Reset
		err := hn.NotifyHeaderSubscribers(header)
		require.NoError(t, err)
	})

	t.Run("subscribers called correctly", func(t *testing.T) {
		hn.subscribers = nil // Reset
		calledCount := 0
		subscriber := &testscommon.HeaderSubscriberMock{
			AddHeaderCalled: func(hash []byte, h sovereign.IncomingHeaderHandler) error {
				calledCount++
				require.Equal(t, headerHash, hash)
				require.Equal(t, header, h)
				return nil
			},
		}

		err := hn.RegisterSubscriber(subscriber)
		require.Nil(t, err)

		err = hn.NotifyHeaderSubscribers(header)
		require.NoError(t, err)
		require.Equal(t, 1, calledCount)
	})

	t.Run("subscriber error should log but not fail", func(t *testing.T) {
		hn.subscribers = nil // Reset
		expectedErr := errors.New("subscriber error")
		subscriber := &testscommon.HeaderSubscriberMock{
			AddHeaderCalled: func(hash []byte, h sovereign.IncomingHeaderHandler) error {
				return expectedErr
			},
		}

		err := hn.RegisterSubscriber(subscriber)
		require.Nil(t, err)

		err = hn.NotifyHeaderSubscribers(header)
		require.NoError(t, err)
	})

	t.Run("concurrent notify should be safe", func(t *testing.T) {
		hn.subscribers = nil // Reset
		numNotifications := 100
		numSubscribers := 100
		var wg sync.WaitGroup
		wg.Add(numNotifications)

		go func() {
			for i := 0; i < numSubscribers; i++ {
				err := hn.RegisterSubscriber(&testscommon.HeaderSubscriberMock{
					AddHeaderCalled: func(hash []byte, h sovereign.IncomingHeaderHandler) error {
						return nil
					},
				})
				require.Nil(t, err)
			}
		}()

		go func() {
			for i := 0; i < numNotifications; i++ {
				go func() {
					defer wg.Done()
					err := hn.NotifyHeaderSubscribers(header)
					if err != nil {
						t.Error(err)
					}
				}()
			}
		}()

		wg.Wait()
	})
}
