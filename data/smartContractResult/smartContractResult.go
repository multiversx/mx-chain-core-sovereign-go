//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=$GOPATH/src smartContractResult.proto
package smartContractResult

import (
	"math/big"

	"github.com/multiversx/mx-chain-core-go/data"
)

var _ = data.TransactionHandler(&SmartContractResult{})

// IsInterfaceNil verifies if underlying object is nil
func (scr *SmartContractResult) IsInterfaceNil() bool {
	return scr == nil
}

// SetValue sets the value of the smart contract result
func (scr *SmartContractResult) SetValue(value *big.Int) {
	scr.Value = value
}

// SetData sets the data of the smart contract result
func (scr *SmartContractResult) SetData(data []byte) {
	scr.Data = data
}

// SetRcvAddr sets the receiver address of the smart contract result
func (scr *SmartContractResult) SetRcvAddr(addr []byte) {
	scr.RcvAddr = addr
}

// SetSndAddr sets the sender address of the smart contract result
func (scr *SmartContractResult) SetSndAddr(addr []byte) {
	scr.SndAddr = addr
}

// GetRcvUserName returns the receiver user name from the smart contract result
func (_ *SmartContractResult) GetRcvUserName() []byte {
	return nil
}

// TrimSlicePtr creates a copy of the provided slice without the excess capacity
func TrimSlicePtr(in []*SmartContractResult) []*SmartContractResult {
	if len(in) == 0 {
		return []*SmartContractResult{}
	}
	ret := make([]*SmartContractResult, len(in))
	copy(ret, in)
	return ret
}

// CheckIntegrity checks for not nil fields and negative value
func (scr *SmartContractResult) CheckIntegrity() error {
	if len(scr.RcvAddr) == 0 {
		return data.ErrNilRcvAddr
	}
	if len(scr.SndAddr) == 0 {
		return data.ErrNilSndAddr
	}
	if scr.Value == nil {
		return data.ErrNilValue
	}
	if scr.Value.Sign() < 0 {
		return data.ErrNegativeValue
	}
	if len(scr.PrevTxHash) == 0 {
		return data.ErrNilTxHash
	}

	return nil
}

// GetChainID returns the chain id from the smart contract result
func (scr *SmartContractResult) GetChainID() []byte {
	return nil
}

// GetSndUserName returns the sender username from the smart contract result
func (_ *SmartContractResult) GetSndUserName() []byte {
	return nil
}

// GetVersion returns the version of the smart contract result
func (_ *SmartContractResult) GetVersion() uint32 {
	return 0
}

// GetSignature returns the signature of the smart contract result
func (scr *SmartContractResult) GetSignature() []byte {
	return nil
}

// GetOptions returns the options of the smart contract result
func (_ *SmartContractResult) GetOptions() uint32 {
	return 0
}

// GetGuardianAddr returns the guardian address from the smart contract result
func (scr *SmartContractResult) GetGuardianAddr() []byte {
	return nil
}

// GetGuardianSignature returns the guardian signature of the smart contract result
func (scr *SmartContractResult) GetGuardianSignature() []byte {
	return nil
}

// GetRelayerSignature returns the relayer signature of the smart contract result
func (scr *SmartContractResult) GetRelayerSignature() []byte {
	return nil
}

// SetGasLimit sets the gas limit of the smart contract result
func (scr *SmartContractResult) SetGasLimit(gasLimit uint64) {
	scr.GasLimit = gasLimit
}

// SetGasPrice sets the gas price of the smart contract result
func (scr *SmartContractResult) SetGasPrice(gasPrice uint64) {
	scr.GasPrice = gasPrice
}

// SetSignature sets the signature of the smart contract result
func (scr *SmartContractResult) SetSignature(_ []byte) {
}

// SetGuardianAddr sets the guardian address of the smart contract result
func (scr *SmartContractResult) SetGuardianAddr(_ []byte) {
}

// SetGuardianSignature sets the guardian signature of the smart contract result
func (scr *SmartContractResult) SetGuardianSignature(_ []byte) {
}

// HasOptionGuardianSet checks if guardian is set for the smart contract result
func (scr *SmartContractResult) HasOptionGuardianSet() bool {
	return false
}
