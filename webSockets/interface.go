package webSockets

import (
	"context"
	"io"

	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/multiversx/mx-chain-core-go/marshal"
	outportSenderData "github.com/multiversx/mx-chain-core-go/webSockets/data"
)

// Driver is an interface for saving node specific data to other storage.
// This could be an elastic search index, a MySql database or any other external services.
type Driver interface {
	SaveBlock(outportBlock *outport.OutportBlock) error
	RevertIndexedBlock(blockData *outport.BlockData) error
	SaveRoundsInfo(roundsInfos *outport.RoundsInfo) error
	SaveValidatorsPubKeys(validatorsPubKeys *outport.ValidatorsPubKeys) error
	SaveValidatorsRating(validatorsRating *outport.ValidatorsRating) error
	SaveAccounts(accounts *outport.Accounts) error
	FinalizedBlock(finalizedBlock *outport.FinalizedBlock) error
	GetMarshaller() marshal.Marshalizer
	Close() error
	IsInterfaceNil() bool
}

// WebSocketSenderHandler defines what the actions that a web socket sender should do
type WebSocketSenderHandler interface {
	Send(args outportSenderData.WsSendArgs) error
	Close() error
	IsInterfaceNil() bool
}

// HostWebSockets -
type HostWebSockets interface {
	Send(args outportSenderData.WsSendArgs) error
	SetPayloadHandler(handler PayloadHandler) error
	Start()
	Listen()
	Close() error
	IsInterfaceNil() bool
}

type PayloadHandler interface {
	ProcessPayload(payload []byte) error
	Close() error
}

// PayloadParser defines what a websocket payload parser should do
type PayloadParser interface {
	ExtractPayloadData(payload []byte) (*outportSenderData.PayloadData, error)
	ExtendPayloadWithCounter(payload []byte, counter uint64, withAcknowledge bool) []byte
	ExtendPayloadWithOperationType(payload []byte, operation outportSenderData.OperationType) []byte
	EncodeUint64(counter uint64) []byte
	DecodeCounter(payload []byte) (uint64, error)
	IsInterfaceNil() bool
}

// WSConClient defines what a web-sockets connection client should be able to do
type WSConClient interface {
	io.Closer
	OpenConnection(url string) error
	WriteMessage(messageType int, data []byte) error
	ReadMessage() (int, []byte, error)
	GetID() string
	IsInterfaceNil() bool
}

// HttpServerHandler defines the minimum behaviour of a http server
type HttpServerHandler interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

// Uint64ByteSliceConverter converts byte slice to/from uint64
type Uint64ByteSliceConverter interface {
	ToByteSlice(uint64) []byte
	ToUint64([]byte) (uint64, error)
	IsInterfaceNil() bool
}
