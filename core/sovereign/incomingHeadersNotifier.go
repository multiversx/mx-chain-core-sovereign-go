package sovereign

import (
	"encoding/hex"
	"fmt"
	"sync"

	logger "github.com/multiversx/mx-chain-logger-go"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data/sovereign"
	"github.com/multiversx/mx-chain-core-go/hashing"
	"github.com/multiversx/mx-chain-core-go/marshal"
)

var log = logger.GetOrCreate("incoming-headers-notifier")

type headersNotifier struct {
	mutSubscribers sync.RWMutex
	subscribers    []IncomingHeaderSubscriber

	marshaller marshal.Marshalizer
	hasher     hashing.Hasher
}

// NewHeadersNotifier creates an incoming headers notifier
func NewHeadersNotifier(marshaller marshal.Marshalizer, hasher hashing.Hasher) (*headersNotifier, error) {
	if check.IfNil(marshaller) {
		return nil, core.ErrNilMarshalizer
	}
	if check.IfNil(hasher) {
		return nil, core.ErrNilHasher
	}

	return &headersNotifier{
		mutSubscribers: sync.RWMutex{},
		subscribers:    make([]IncomingHeaderSubscriber, 0),
		marshaller:     marshaller,
		hasher:         hasher,
	}, nil
}

// RegisterSubscriber will register an incoming header subscriber
func (hn *headersNotifier) RegisterSubscriber(handler IncomingHeaderSubscriber) error {
	if check.IfNil(handler) {
		return errNilHeaderSubscriber
	}

	hn.mutSubscribers.Lock()
	hn.subscribers = append(hn.subscribers, handler)
	hn.mutSubscribers.Unlock()

	return nil
}

// NotifyHeaderSubscribers will notify subscribed headers for an incoming header
func (hn *headersNotifier) NotifyHeaderSubscribers(header sovereign.IncomingHeaderHandler) error {
	if check.IfNil(header) {
		return errNilIncomingHeader
	}

	headerHash, err := core.CalculateHash(hn.marshaller, hn.hasher, header)
	if err != nil {
		return err
	}

	log.Debug("notifying incoming header", "chain", header.GetSourceChainID(), "nonce", header.GetNonce(), "hash", hex.EncodeToString(headerHash))

	hn.mutSubscribers.RLock()
	defer hn.mutSubscribers.RUnlock()

	for _, handler := range hn.subscribers {
		err = handler.AddHeader(headerHash, header)
		if err != nil {
			log.Error("headersNotifier.NotifyHeaderSubscribers", "handler", fmt.Sprintf("%T", handler), "error", err)
		}
	}

	return nil
}
