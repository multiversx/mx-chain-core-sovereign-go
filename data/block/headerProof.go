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

// SetLeaderSignature will set the leader sig
func (x *ExtraSignatureData) SetLeaderSignature(leaderSig []byte) error {
	if x == nil {
		return data.ErrNilPointerReceiver
	}

	x.LeaderSignature = leaderSig
	return nil
}

// SetAggregatedSignature will set the aggregated sig
func (x *ExtraSignatureData) SetAggregatedSignature(aggregatedSig []byte) error {
	if x == nil {
		return data.ErrNilPointerReceiver
	}

	x.AggregatedSignature = aggregatedSig
	return nil
}
