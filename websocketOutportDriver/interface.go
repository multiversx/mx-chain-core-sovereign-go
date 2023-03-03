package websocketOutportDriver

import (
	"github.com/multiversx/mx-chain-core-go/data/alteredAccount"
	"github.com/multiversx/mx-chain-core-go/data/outport"
	outportSenderData "github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
)

// Driver is an interface for saving node specific data to other storage.
// This could be an elastic search index, a MySql database or any other external services.
type Driver interface {
	SaveBlock(outportBlock *outport.OutportBlock) error
	RevertIndexedBlock(blockData *outport.BlockData) error
	SaveRoundsInfo(roundsInfos *outport.RoundsInfo) error
	SaveValidatorsPubKeys(validatorsPubKeys *outport.ValidatorsPubKeys) error
	SaveValidatorsRating(indexID string, infoRating []*outport.ValidatorRatingInfo) error
	SaveAccounts(blockTimestamp uint64, acc map[string]*alteredAccount.AlteredAccount, shardID uint32) error
	FinalizedBlock(headerHash []byte) error
	Close() error
	IsInterfaceNil() bool
}

// WebSocketSenderHandler defines what the actions that a web socket sender should do
type WebSocketSenderHandler interface {
	Send(args outportSenderData.WsSendArgs) error
	AddClient(wss outportSenderData.WSConn, remoteAddr string)
	Close() error
	IsInterfaceNil() bool
}

// Uint64ByteSliceConverter converts byte slice to/from uint64
type Uint64ByteSliceConverter interface {
	ToByteSlice(uint64) []byte
	ToUint64([]byte) (uint64, error)
	IsInterfaceNil() bool
}
