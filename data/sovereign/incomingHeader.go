//go:generate protoc -I=. -I=$GOPATH/src/github.com/multiversx/mx-chain-core-go/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. incomingHeader.proto

package sovereign

import (
	"math/big"

	"github.com/multiversx/mx-chain-core-go/data"
)

// GetIncomingEventHandlers returns the incoming events as an array of event handlers
func (ih *IncomingHeader) GetIncomingEventHandlers() []data.EventHandler {
	if ih == nil {
		return nil
	}

	events := ih.GetIncomingEvents()
	logHandlers := make([]data.EventHandler, len(events))

	for i := range events {
		logHandlers[i] = events[i]
	}

	return logHandlers
}

// GetNonceBI returns the nonce as big.Int
func (ih *IncomingHeader) GetNonceBI() *big.Int {
	return ih.Nonce
}

// IsInterfaceNil checks if the underlying pointer is nil
func (ih *IncomingHeader) IsInterfaceNil() bool {
	return ih == nil
}
