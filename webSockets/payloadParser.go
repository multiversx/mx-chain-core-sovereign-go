package webSockets

import (
	"bytes"
	"fmt"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
)

const (
	withAcknowledgeNumBytes = 1
	uint64NumBytes          = 8
	uint32NumBytes          = 4
)

var (
	minBytesForCorrectPayload = withAcknowledgeNumBytes + uint64NumBytes + uint32NumBytes + uint32NumBytes

	prefixWithoutAck = []byte{0}
	prefixWithAck    = []byte{1}
)

type webSocketsPayloadConverter struct {
	uint64ByteSliceConverter Uint64ByteSliceConverter
}

// NewWebSocketPayloadParser returns a new instance of websocketPayloadParser
func NewWebSocketPayloadParser(uint64ByteSliceConverter Uint64ByteSliceConverter) (*webSocketsPayloadConverter, error) {
	if check.IfNil(uint64ByteSliceConverter) {
		return nil, data.ErrNilUint64ByteSliceConverter
	}

	return &webSocketsPayloadConverter{
		uint64ByteSliceConverter: uint64ByteSliceConverter,
	}, nil
}

// ExtractPayloadData will extract the data from the received payload
// It should have the following form:
// first byte - with acknowledge or not
// next 8 bytes - counter (uint64 big endian)
// next 4 bytes - operation type (uint32 big endian)
// next 4 bytes - message length (uint32 big endian)
// next X bytes - the actual data to parse
func (wpc *webSocketsPayloadConverter) ExtractPayloadData(payload []byte) (*data.PayloadData, error) {
	if len(payload) < minBytesForCorrectPayload {
		return nil, fmt.Errorf("invalid payload. minimum required length is %d bytes, but only provided %d",
			minBytesForCorrectPayload,
			len(payload))
	}

	var err error
	payloadData := &data.PayloadData{
		WithAcknowledge: false,
	}

	if payload[0] == byte(1) {
		payloadData.WithAcknowledge = true
	}
	payload = payload[withAcknowledgeNumBytes:]

	counterBytes := payload[:uint64NumBytes]
	payloadData.Counter, err = wpc.uint64ByteSliceConverter.ToUint64(counterBytes)
	if err != nil {
		return nil, fmt.Errorf("%w while extracting the counter from the payload", err)
	}
	payload = payload[uint64NumBytes:]

	operationTypeBytes := payload[:uint32NumBytes]
	var operationTypeUint64 uint64
	operationTypeUint64, err = wpc.uint64ByteSliceConverter.ToUint64(padUint32ByteSlice(operationTypeBytes))
	if err != nil {
		return nil, fmt.Errorf("%w while extracting the counter from the payload", err)
	}
	payloadData.OperationType = data.OperationTypeFromUint64(operationTypeUint64)
	payload = payload[uint32NumBytes:]

	var messageLen uint64
	messageLen, err = wpc.uint64ByteSliceConverter.ToUint64(padUint32ByteSlice(payload[:uint32NumBytes]))
	if err != nil {
		return nil, fmt.Errorf("%w while extracting the message length", err)
	}
	payload = payload[uint32NumBytes:]

	if messageLen != uint64(len(payload)) {
		return nil, fmt.Errorf("message counter is not equal to the actual payload. provided: %d, actual: %d",
			messageLen, len(payload))
	}

	payloadData.Payload = payload

	return payloadData, nil
}

// ExtendPayloadWithCounter will put in the provided payload the provided counter if needed
func (wpc *webSocketsPayloadConverter) ExtendPayloadWithCounter(payload []byte, counter uint64, withAcknowledge bool) []byte {
	ackData := prefixWithoutAck
	if withAcknowledge {
		ackData = prefixWithAck
	}

	newPayload := append(ackData, wpc.EncodeUint64(counter)...)
	newPayload = append(newPayload, payload...)
	return newPayload
}

// ExtendPayloadWithOperationType will extend the provided payload with the provided operation type
func (wpc *webSocketsPayloadConverter) ExtendPayloadWithOperationType(payload []byte, operation data.OperationType) []byte {
	opBytes := wpc.EncodeUint64(uint64(operation.Uint32()))
	opBytes = opBytes[uint32NumBytes:]

	messageLength := uint64(len(payload))
	messageLengthBytes := wpc.uint64ByteSliceConverter.ToByteSlice(messageLength)
	messageLengthBytes = messageLengthBytes[uint32NumBytes:]

	newPayload := append(opBytes, messageLengthBytes...)
	newPayload = append(newPayload, payload...)

	return newPayload
}

// EncodeUint64 will encode the provided counter in a byte array
func (wpc *webSocketsPayloadConverter) EncodeUint64(counter uint64) []byte {
	return wpc.uint64ByteSliceConverter.ToByteSlice(counter)
}

// DecodeCounter will decode the provided payload in a uint64 value
func (wpc *webSocketsPayloadConverter) DecodeCounter(payload []byte) (uint64, error) {
	return wpc.uint64ByteSliceConverter.ToUint64(payload)
}

func padUint32ByteSlice(initial []byte) []byte {
	padding := bytes.Repeat([]byte{0}, 4)
	return append(padding, initial...)
}

// IsInterfaceNil -
func (wpc *webSocketsPayloadConverter) IsInterfaceNil() bool {
	return wpc == nil
}
