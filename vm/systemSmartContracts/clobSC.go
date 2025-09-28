package systemSmartContracts

import (
	"encoding/json"

	"github.com/nikolaydubina/fpdecimal"
)

const (
	processOrderEndpoint = "processOrder"
	cancelOrderEndpoint  = "cancelOrder"
	getOrderEndpoint     = "getOrder"
	getDepthEndpoint     = "getDepth"
)

const orderBookStorageKey = "orderBook"

// EEI is a temporary interface definition based on user guidance.
// I expect this to be replaced by the actual EEI interface from the project.
type EEI interface {
	GetStorage(key []byte) ([]byte, error)
	SetStorage(key []byte, value []byte) error
}

// clobSC is the system smart contract for the Central Limit Order Book.
type clobSC struct {
	eei EEI
}

// NewClobSC creates a new instance of the CLOB system smart contract.
func NewClobSC(eei EEI) *clobSC {
	return &clobSC{
		eei: eei,
	}
}

// ProcessOrder is the endpoint for processing a new order.
func (sc *clobSC) ProcessOrder(
	orderID string,
	side Side,
	orderType OrderType,
	quantity, price, stop fpdecimal.Decimal,
	tif TIF,
	oco string,
) ([]byte, error) {
	clob := NewCLOB()

	// Load state from storage
	data, err := sc.eei.GetStorage([]byte(orderBookStorageKey))
	if err != nil {
		return nil, err
	}
	err = clob.LoadState(data)
	if err != nil {
		return nil, err
	}

	// Process the order
	done, err := clob.ProcessOrder(orderID, side, orderType, quantity, price, stop, tif, oco)
	if err != nil {
		return nil, err
	}

	// Save the new state
	newState, err := clob.SaveState()
	if err != nil {
		return nil, err
	}
	err = sc.eei.SetStorage([]byte(orderBookStorageKey), newState)
	if err != nil {
		return nil, err
	}

	return json.Marshal(done)
}

// CancelOrder is the endpoint for canceling an order.
func (sc *clobSC) CancelOrder(orderID string) ([]byte, error) {
	clob := NewCLOB()

	// Load state from storage
	data, err := sc.eei.GetStorage([]byte(orderBookStorageKey))
	if err != nil {
		return nil, err
	}
	err = clob.LoadState(data)
	if err != nil {
		return nil, err
	}

	// Cancel the order
	order := clob.CancelOrder(orderID)

	// Save the new state
	newState, err := clob.SaveState()
	if err != nil {
		return nil, err
	}
	err = sc.eei.SetStorage([]byte(orderBookStorageKey), newState)
	if err != nil {
		return nil, err
	}

	return json.Marshal(order)
}

// GetOrder is the endpoint for retrieving an order.
func (sc *clobSC) GetOrder(orderID string) ([]byte, error) {
	clob := NewCLOB()

	// Load state from storage
	data, err := sc.eei.GetStorage([]byte(orderBookStorageKey))
	if err != nil {
		return nil, err
	}
	err = clob.LoadState(data)
	if err != nil {
		return nil, err
	}

	// Get the order
	order := clob.GetOrder(orderID)

	return json.Marshal(order)
}

// GetDepth is the endpoint for retrieving the order book depth.
func (sc *clobSC) GetDepth() ([]byte, error) {
	clob := NewCLOB()

	// Load state from storage
	data, err := sc.eei.GetStorage([]byte(orderBookStorageKey))
	if err != nil {
		return nil, err
	}
	err = clob.LoadState(data)
	if err != nil {
		return nil, err
	}

	// Get the depth
	depth := clob.GetDepth()

	return json.Marshal(depth)
}