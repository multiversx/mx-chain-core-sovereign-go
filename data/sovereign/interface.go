package sovereign

import (
	"math/big"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/sovereign/dto"
)

// IncomingHeaderHandler defines the incoming header to a sovereign chain that is sent by a notifier
type IncomingHeaderHandler interface {
	GetProof() []byte
	GetNonceBI() *big.Int
	GetSourceChainID() dto.ChainID
	GetIncomingEventHandlers() []data.EventHandler
	IsInterfaceNil() bool
}

// Proof defines the proof of proposed mini blocks, using provided incoming header and events
type Proof interface {
	GetHeaderHandler() data.HeaderHandler
	GetIncomingEventHandlers() []data.EventHandler
	GetIncomingMiniBlockHandlers() []data.MiniBlockHandler
	IsInterfaceNil() bool
}
