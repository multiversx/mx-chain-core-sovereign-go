//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=$GOPATH/src rewardTx.proto
package rewardTx

import (
	"math/big"

	"github.com/multiversx/mx-chain-core-go/data"
)

var _ = data.TransactionHandler(&RewardTx{})

// IsInterfaceNil verifies if underlying object is nil
func (rtx *RewardTx) IsInterfaceNil() bool {
	return rtx == nil
}

// SetValue sets the value of the transaction
func (rtx *RewardTx) SetValue(value *big.Int) {
	rtx.Value = value
}

// GetNonce returns 0 as reward transactions do not have a nonce
func (rtx *RewardTx) GetNonce() uint64 {
	return 0
}

// GetData returns the data of the reward transaction
func (rtx *RewardTx) GetData() []byte {
	return []byte("")
}

// GetSndAddr returns the sender address from the reward transaction
func (rtx *RewardTx) GetSndAddr() []byte {
	return nil
}

// GetGasLimit returns the gas limit of the smart reward transaction
func (rtx *RewardTx) GetGasLimit() uint64 {
	return 0
}

// GetGasPrice returns the gas price of the smart reward transaction
func (rtx *RewardTx) GetGasPrice() uint64 {
	return 0
}

// SetData sets the data of the reward transaction
func (rtx *RewardTx) SetData(_ []byte) {
}

// SetRcvAddr sets the receiver address of the reward transaction
func (rtx *RewardTx) SetRcvAddr(addr []byte) {
	rtx.RcvAddr = addr
}

// SetSndAddr sets the sender address of the reward transaction
func (rtx *RewardTx) SetSndAddr(_ []byte) {
}

// GetRcvUserName returns the receiver user name from the reward transaction
func (rtx *RewardTx) GetRcvUserName() []byte {
	return nil
}

// CheckIntegrity checks for not nil fields and negative value
func (rtx *RewardTx) CheckIntegrity() error {
	if len(rtx.RcvAddr) == 0 {
		return data.ErrNilRcvAddr
	}
	if rtx.Value == nil {
		return data.ErrNilValue
	}
	if rtx.Value.Sign() < 0 {
		return data.ErrNegativeValue
	}

	return nil
}

// GetChainID returns the chain id from the reward transaction
func (rtx *RewardTx) GetChainID() []byte {
	return nil
}

// GetSndUserName returns the sender user name from the receipt
func (rtx *RewardTx) GetSndUserName() []byte {
	return nil
}

// GetVersion returns the version of the reward transaction
func (rtx *RewardTx) GetVersion() uint32 {
	return 0
}

// GetSignature returns the signature of the reward transaction
func (rtx *RewardTx) GetSignature() []byte {
	return nil
}

// GetOptions returns the options of the reward transaction
func (rtx *RewardTx) GetOptions() uint32 {
	return 0
}

// GetGuardianAddr returns the guardian address from the reward transaction
func (rtx *RewardTx) GetGuardianAddr() []byte {
	return nil
}

// GetGuardianSignature returns the guardian signature of the reward transaction
func (rtx *RewardTx) GetGuardianSignature() []byte {
	return nil
}

// GetRelayerAddr returns the relayer address from the reward transaction
func (rtx *RewardTx) GetRelayerAddr() []byte {
	return nil
}

// GetRelayerSignature returns the relayer signature of the reward transaction
func (rtx *RewardTx) GetRelayerSignature() []byte {
	return nil
}

// SetGasLimit sets the gas limit of the reward transaction
func (rtx *RewardTx) SetGasLimit(_ uint64) {
}

// SetGasPrice sets the gas price of the reward transaction
func (rtx *RewardTx) SetGasPrice(_ uint64) {
}

// SetSignature sets the signature of the reward transaction
func (rtx *RewardTx) SetSignature(_ []byte) {
}

// SetGuardianAddr sets the guardian address of the reward transaction
func (rtx *RewardTx) SetGuardianAddr(_ []byte) {
}

// SetGuardianSignature sets the guardian signature of the reward transaction
func (rtx *RewardTx) SetGuardianSignature(_ []byte) {
}

// HasOptionGuardianSet checks if guardian is set for the receipt
func (rtx *RewardTx) HasOptionGuardianSet() bool {
	return false
}
