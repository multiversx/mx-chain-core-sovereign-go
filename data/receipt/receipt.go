//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=$GOPATH/src receipt.proto
package receipt

import (
	"math/big"

	"github.com/multiversx/mx-chain-core-go/data"
)

var _ = data.TransactionHandler(&Receipt{})

// IsInterfaceNil verifies if underlying object is nil
func (rpt *Receipt) IsInterfaceNil() bool {
	return rpt == nil
}

// GetNonce returns the nonce of the receipt
func (rpt *Receipt) GetNonce() uint64 {
	return 0
}

// GetRcvAddr returns the receiver address from the receipt
func (rpt *Receipt) GetRcvAddr() []byte {
	return rpt.SndAddr
}

// GetGasLimit returns the gas limit of the receipt
func (rpt *Receipt) GetGasLimit() uint64 {
	return 0
}

// GetGasPrice returns the gas price of the receipt
func (rpt *Receipt) GetGasPrice() uint64 {
	return 0
}

// SetValue sets the value of the receipt
func (rpt *Receipt) SetValue(value *big.Int) {
	rpt.Value = value
}

// SetData sets the data of the receipt
func (rpt *Receipt) SetData(data []byte) {
	rpt.Data = data
}

// SetRcvAddr sets the receiver address of the receipt
func (rpt *Receipt) SetRcvAddr(_ []byte) {
}

// SetSndAddr sets the sender address of the receipt
func (rpt *Receipt) SetSndAddr(addr []byte) {
	rpt.SndAddr = addr
}

// GetRcvUserName returns the receiver user name from the receipt
func (_ *Receipt) GetRcvUserName() []byte {
	return nil
}

// CheckIntegrity checks for not nil fields and negative value
func (rpt *Receipt) CheckIntegrity() error {
	return nil
}

// GetChainID returns the chain id of the receipt
func (rpt *Receipt) GetChainID() []byte {
	return nil
}

// GetSndUserName returns the sender username from the receipt
func (rpt *Receipt) GetSndUserName() []byte {
	return nil
}

// GetVersion returns the version of the receipt
func (rpt *Receipt) GetVersion() uint32 {
	return 0
}

// GetSignature returns the signature of the receipt
func (rpt *Receipt) GetSignature() []byte {
	return nil
}

// GetOptions returns the options of the receipt
func (rpt *Receipt) GetOptions() uint32 {
	return 0
}

// GetGuardianAddr returns the guardian address from the receipt
func (rpt *Receipt) GetGuardianAddr() []byte {
	return nil
}

// GetGuardianSignature returns the guardian signature of the receipt
func (rpt *Receipt) GetGuardianSignature() []byte {
	return nil
}

// GetRelayerAddr returns the relayer address from the receipt
func (rpt *Receipt) GetRelayerAddr() []byte {
	return nil
}

// GetRelayerSignature returns the relayer signature of the receipt
func (rpt *Receipt) GetRelayerSignature() []byte {
	return nil
}

// SetGasLimit sets the gas limit of the receipt
func (rpt *Receipt) SetGasLimit(_ uint64) {
}

// SetGasPrice sets the gas price of the receipt
func (rpt *Receipt) SetGasPrice(_ uint64) {
}

// SetSignature sets the signature of the receipt
func (rpt *Receipt) SetSignature(_ []byte) {
}

// HasOptionGuardianSet checks if guardian is set for the receipt
func (rpt *Receipt) HasOptionGuardianSet() bool {
	return false
}
