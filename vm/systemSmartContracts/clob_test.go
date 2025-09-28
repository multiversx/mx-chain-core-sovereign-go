package systemSmartContracts

import (
	"testing"

	"github.com/nikolaydubina/fpdecimal"
	"github.com/stretchr/testify/assert"
)

func TestCLOB_StateSaveAndLoad(t *testing.T) {
	clob1 := NewCLOB()

	// Process a limit order
	done, err := clob1.ProcessOrder(
		"order1",
		Buy,
		TypeLimit,
		fpdecimal.FromInt(10),
		fpdecimal.FromInt(100),
		fpdecimal.Zero,
		GTC,
		"",
	)
	assert.NoError(t, err)
	assert.NotNil(t, done)
	assert.True(t, done.Stored)

	// Save the state
	state, err := clob1.SaveState()
	assert.NoError(t, err)
	assert.NotEmpty(t, state)

	// Create a new CLOB and load the state
	clob2 := NewCLOB()
	err = clob2.LoadState(state)
	assert.NoError(t, err)

	// Verify that the state was loaded correctly
	order := clob2.GetOrder("order1")
	assert.NotNil(t, order)
	assert.Equal(t, "order1", order.GetID())
	assert.Equal(t, fpdecimal.FromInt(10), order.GetQuantity())
	assert.Equal(t, fpdecimal.FromInt(100), order.GetPrice())
}