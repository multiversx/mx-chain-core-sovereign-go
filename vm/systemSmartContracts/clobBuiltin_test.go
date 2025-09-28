package systemSmartContracts

import (
	"encoding/json"
	"testing"

	vmcommon "github.com/multiversx/mx-chain-vm-common-go"
	"github.com/stretchr/testify/assert"
)

func TestClobBuiltin_ProcessBuiltinFunction(t *testing.T) {
	storage := newMockAccountDataHandler()
	sc := NewClobSC(storage)
	builtin := NewClobBuiltin(sc)

	// Prepare input for processOrder
	orderID := "order1"
	side := "1" // Buy
	orderType := "LIMIT"
	quantity := "10"
	price := "100"
	stop := "0"
	tif := "GTC"
	oco := ""

	input := &vmcommon.ContractCallInput{
		Function: processOrderEndpoint,
		Arguments: [][]byte{
			[]byte(orderID),
			[]byte(side),
			[]byte(orderType),
			[]byte(quantity),
			[]byte(price),
			[]byte(stop),
			[]byte(tif),
			[]byte(oco),
		},
	}

	// Execute the call
	output, err := builtin.ProcessBuiltinFunction(nil, nil, input)
	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, vmcommon.Ok, output.ReturnCode)

	// Check the result
	var done Done
	err = json.Unmarshal(output.ReturnData[0], &done)
	assert.NoError(t, err)
	assert.True(t, done.Stored)
	assert.Equal(t, orderID, done.Order.GetID())

	// Check that the state was updated
	assert.NotEmpty(t, storage.storage[orderBookStorageKey])
}