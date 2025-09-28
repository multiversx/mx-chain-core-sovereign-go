package systemSmartContracts

import (
	"encoding/json"
	"testing"

	"github.com/nikolaydubina/fpdecimal"
	"github.com/stretchr/testify/assert"
)

// mockEEI implements the EEI interface for testing.
type mockEEI struct {
	storage map[string][]byte
}

func newMockEEI() *mockEEI {
	return &mockEEI{
		storage: make(map[string][]byte),
	}
}

func (m *mockEEI) GetStorage(key []byte) ([]byte, error) {
	return m.storage[string(key)], nil
}

func (m *mockEEI) SetStorage(key, value []byte) error {
	m.storage[string(key)] = value
	return nil
}

func TestClobSC_ProcessOrder(t *testing.T) {
	eei := newMockEEI()
	sc := NewClobSC(eei)

	// Process a limit order
	doneBytes, err := sc.ProcessOrder(
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

	// Check the result
	var done Done
	err = json.Unmarshal(doneBytes, &done)
	assert.NoError(t, err)
	assert.True(t, done.Stored)

	// Check that the state was updated
	assert.NotEmpty(t, eei.storage[orderBookStorageKey])
}

func TestClobSC_FullFlow(t *testing.T) {
	eei := newMockEEI()
	sc := NewClobSC(eei)

	// 1. Process a limit order
	_, err := sc.ProcessOrder(
		"order1",
		Sell,
		TypeLimit,
		fpdecimal.FromInt(10),
		fpdecimal.FromInt(100),
		fpdecimal.Zero,
		GTC,
		"",
	)
	assert.NoError(t, err)

	// 2. Process a market order that partially fills the limit order
	doneBytes, err := sc.ProcessOrder(
		"order2",
		Buy,
		TypeMarket,
		fpdecimal.FromInt(5),
		fpdecimal.FromInt(100),
		fpdecimal.Zero,
		"",
		"",
	)
	assert.NoError(t, err)
	var done Done
	err = json.Unmarshal(doneBytes, &done)
	assert.NoError(t, err)
	assert.False(t, done.Stored)
	assert.Equal(t, "5.000", done.Processed.String())

	// 3. Cancel the remainder of the limit order
	canceledOrderBytes, err := sc.CancelOrder("order1")
	assert.NoError(t, err)
	var canceledOrder Order
	err = json.Unmarshal(canceledOrderBytes, &canceledOrder)
	assert.NoError(t, err)
	assert.True(t, canceledOrder.IsCanceled())

	// 4. Get the order to check its status (it should be gone)
	orderBytes, err := sc.GetOrder("order1")
	assert.NoError(t, err)
	assert.Equal(t, []byte("null"), orderBytes)

	// 5. Get the depth to check the order book state
	depthBytes, err := sc.GetDepth()
	assert.NoError(t, err)
	var depth Depth
	err = json.Unmarshal(depthBytes, &depth)
	assert.NoError(t, err)
	assert.Empty(t, depth.Ask)
	assert.Empty(t, depth.Bid)
}