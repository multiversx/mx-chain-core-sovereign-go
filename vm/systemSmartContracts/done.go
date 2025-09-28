package systemSmartContracts

import (
	"encoding/json"

	"github.com/nikolaydubina/fpdecimal"
)

// Done structure
type Done struct {
	Order     *Order
	Trades    []*TradeOrder
	Canceled  []string
	Activated []string
	Stored    bool
	Quantity  fpdecimal.Decimal
	Left      fpdecimal.Decimal
	Processed fpdecimal.Decimal
}

// DoneJSON structure
type DoneJSON struct {
	Order     *TradeOrder  `json:"order"`
	Trades    []TradeOrder `json:"trades"`
	Canceled  []string     `json:"canceled"`
	Activated []string     `json:"activated"`
	Left      string       `json:"left"`
	Processed string       `json:"processed"`
	Stored    bool         `json:"stored"`
}

func newDone(order *Order) *Done {
	return &Done{
		Order:     order,
		Trades:    make([]*TradeOrder, 0),
		Canceled:  make([]string, 0),
		Activated: make([]string, 0),
		Quantity:  order.GetOriginalQty(),
		Left:      fpdecimal.Zero,
		Processed: fpdecimal.Zero,
	}
}

// GetTradeOrder returns TradeOrder by id
func (d *Done) GetTradeOrder(id string) *TradeOrder {
	for _, t := range d.Trades {
		if t.OrderID == id {
			return t
		}
	}
	return nil
}

func (d *Done) appendOrder(order *Order, quantity, price fpdecimal.Decimal) {

	if len(d.Trades) == 0 {
		d.Trades = append(d.Trades, newTradeOrder(d.Order, fpdecimal.Zero, d.Order.GetPrice()))
	}

	d.Trades = append(d.Trades, newTradeOrder(order, quantity, price))
}

func (d *Done) tradesToSlice() []TradeOrder {
	slice := make([]TradeOrder, 0, len(d.Trades))
	for _, v := range d.Trades {
		slice = append(slice, *v)
	}
	return slice
}

func (d *Done) appendCanceled(order *Order) {
	d.Canceled = append(d.Canceled, order.GetID())
}

func (d *Done) appendActivated(order *Order) {
	d.Activated = append(d.Activated, order.GetID())
}

func (d *Done) setLeftQuantity(quantity *fpdecimal.Decimal) {
	if len(d.Trades) == 0 {
		return
	}
	d.Left = *quantity
	d.Processed = d.Quantity.Sub(d.Left)
	if len(d.Trades) != 0 {
		d.Trades[0].Quantity = d.Processed
	}
}

// MarshalJSON implements Marshaler interface
func (d *Done) MarshalJSON() ([]byte, error) {
	customStruct := struct {
		Order     *Order            `json:"order"`
		Trades    []*TradeOrder     `json:"trades"`
		Canceled  []string          `json:"canceled"`
		Activated []string          `json:"activated"`
		Left      fpdecimal.Decimal `json:"left"`
		Processed fpdecimal.Decimal `json:"processed"`
		Stored    bool              `json:"stored"`
	}{
		Order:     d.Order,
		Trades:    d.Trades,
		Canceled:  d.Canceled,
		Activated: d.Activated,
		Left:      d.Left,
		Processed: d.Processed,
		Stored:    d.Stored,
	}
	return json.Marshal(customStruct)
}

// String implements Stringer interface
func (d *Done) String() string {
	j, _ := d.MarshalJSON()
	return string(j)
}