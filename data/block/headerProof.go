//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. headerProof.proto
package block

import (
	"github.com/multiversx/mx-chain-core-go/data"
)

// IsInterfaceNil returns true if there is no value under the interface
func (x *HeaderProof) IsInterfaceNil() bool {
	return x == nil
}

// GetExtraSignatureHandlers returns internal extra signatures data as interfaces
func (x *HeaderProof) GetExtraSignatureHandlers() map[string]data.ExtraSignatureDataHandler {
	if x == nil {
		return nil
	}

	extraSigs := x.GetExtraSignatures()
	extraSigHandlers := make(map[string]data.ExtraSignatureDataHandler, len(extraSigs))

	for id, sigData := range extraSigs {
		copySigData := *sigData
		extraSigHandlers[id] = &copySigData
	}

	return extraSigHandlers
}
