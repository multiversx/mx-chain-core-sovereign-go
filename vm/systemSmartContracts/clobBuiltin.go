package systemSmartContracts

import (
	"fmt"

	"github.com/multiversx/mx-chain-core-go/core"
	vmcommon "github.com/multiversx/mx-chain-vm-common-go"
	"github.com/nikolaydubina/fpdecimal"
)

type clobBuiltin struct {
	sc *clobSC
}

func NewClobBuiltin(sc *clobSC) *clobBuiltin {
	return &clobBuiltin{
		sc: sc,
	}
}

func (cb *clobBuiltin) ProcessBuiltinFunction(
	acntSnd, acntDst vmcommon.UserAccountHandler,
	vmInput *vmcommon.ContractCallInput,
) (*vmcommon.VMOutput, error) {
	funcName := string(vmInput.Function)
	args := vmInput.Arguments

	switch funcName {
	case processOrderEndpoint:
		return cb.processOrder(args)
	case cancelOrderEndpoint:
		return cb.cancelOrder(args)
	case getOrderEndpoint:
		return cb.getOrder(args)
	case getDepthEndpoint:
		return cb.getDepth(args)
	default:
		return nil, fmt.Errorf("invalid function name: %s", funcName)
	}
}

func (cb *clobBuiltin) processOrder(args [][]byte) (*vmcommon.VMOutput, error) {
	// Simplified argument parsing for now. A real implementation would need a robust parser.
	orderID := string(args[0])
	side := Side(args[1][0])
	orderType := OrderType(args[2])
	quantity, _ := fpdecimal.NewFromString(string(args[3]))
	price, _ := fpdecimal.NewFromString(string(args[4]))
	stop, _ := fpdecimal.NewFromString(string(args[5]))
	tif := TIF(args[6])
	oco := string(args[7])

	returnData, err := cb.sc.ProcessOrder(orderID, side, orderType, quantity, price, stop, tif, oco)
	if err != nil {
		return nil, err
	}

	return &vmcommon.VMOutput{
		ReturnData: [][]byte{returnData},
		ReturnCode: vmcommon.Ok,
	}, nil
}

func (cb *clobBuiltin) cancelOrder(args [][]byte) (*vmcommon.VMOutput, error) {
	orderID := string(args[0])
	returnData, err := cb.sc.CancelOrder(orderID)
	if err != nil {
		return nil, err
	}

	return &vmcommon.VMOutput{
		ReturnData: [][]byte{returnData},
		ReturnCode: vmcommon.Ok,
	}, nil
}

func (cb *clobBuiltin) getOrder(args [][]byte) (*vmcommon.VMOutput, error) {
	orderID := string(args[0])
	returnData, err := cb.sc.GetOrder(orderID)
	if err != nil {
		return nil, err
	}

	return &vmcommon.VMOutput{
		ReturnData: [][]byte{returnData},
		ReturnCode: vmcommon.Ok,
	}, nil
}

func (cb *clobBuiltin) getDepth(args [][]byte) (*vmcommon.VMOutput, error) {
	returnData, err := cb.sc.GetDepth()
	if err != nil {
		return nil, err
	}

	return &vmcommon.VMOutput{
		ReturnData: [][]byte{returnData},
		ReturnCode: vmcommon.Ok,
	}, nil
}

// SetNewGasConfig is required by the BuiltinFunction interface.
func (cb *clobBuiltin) SetNewGasConfig(gasCost *vmcommon.GasCost) {}

// IsActive is required by the BuiltinFunction interface.
func (cb *clobBuiltin) IsActive() bool {
	return true
}

// IsInterfaceNil is required by the BuiltinFunction interface.
func (cb *clobBuiltin) IsInterfaceNil() bool {
	return cb == nil
}