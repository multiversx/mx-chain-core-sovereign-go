package block

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/transaction"
)

func TestShardHeaderExtended_SetIncomingEventHandlers(t *testing.T) {
	t.Parallel()

	hdr := &ShardHeaderExtended{
		IncomingEvents: []*transaction.Event{
			{
				Data: []byte("data1"),
			},
		},
	}

	events := []data.EventHandler{
		&transaction.Event{
			Data: []byte("data2"),
		},
	}
	err := hdr.SetIncomingEventHandlers(events)
	require.Nil(t, err)
	require.Equal(t, events, hdr.GetIncomingEventHandlers())

	err = hdr.SetIncomingEventHandlers(nil)
	require.Nil(t, err)
	require.Empty(t, hdr.GetIncomingEventHandlers())

	events = []data.EventHandler{
		&transaction.Event{
			Data: []byte("data3"),
		},
		&transaction.Event{
			Data: []byte("data4"),
		},
	}
	err = hdr.SetIncomingEventHandlers(events)
	require.Nil(t, err)
	require.Equal(t, events, hdr.GetIncomingEventHandlers())

	// check no internal data has been modified by outside pointer
	events[0].(*transaction.Event).Data = []byte("data5")
	require.Equal(t, []data.EventHandler{
		&transaction.Event{
			Data: []byte("data3"),
		},
		&transaction.Event{
			Data: []byte("data4"),
		},
	}, hdr.GetIncomingEventHandlers())
}
