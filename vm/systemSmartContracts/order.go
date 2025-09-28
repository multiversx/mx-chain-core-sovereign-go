package systemSmartContracts

import (
	"encoding/json"

	"github.com/nikolaydubina/fpdecimal"
)

// TradeOrder structure
type TradeOrder struct {
	OrderID  string
	Role     Role
	Price    fpdecimal.Decimal
	IsQuote  bool
	Quantity fpdecimal.Decimal
}

func newTradeOrder(order *Order, quantity, price fpdecimal.Decimal) *TradeOrder {
	return &TradeOrder{
		OrderID:  order.GetID(),
		Role:     order.GetRole(),
		Price:    price,
		IsQuote:  order.GetIsQuote(),
		Quantity: quantity,
	}
}

// Order stores information about order
type Order struct {
	Id          string            `json:"id"`
	OrderType   OrderType         `json:"orderType"`
	Side        Side              `json:"side"`
	IsQuote     bool              `json:"isQuote"`
	Quantity    fpdecimal.Decimal `json:"quantity"`
	OriginalQty fpdecimal.Decimal `json:"originalQty"`
	Price       fpdecimal.Decimal `json:"price"`
	Canceled    bool              `json:"canceled"`
	Role        Role              `json:"role"`
	Stop        fpdecimal.Decimal `json:"stop"`
	Tif         TIF               `json:"tif"`
	Oco         string            `json:"oco"`
}

// NewMarketOrder creates new constant object Order
func NewMarketOrder(orderID string, side Side, quantity fpdecimal.Decimal) *Order {

	if quantity.LessThanOrEqual(fpdecimal.Zero) {
		panic(ErrInvalidQuantity)
	}

	return &Order{
		Id:          orderID,
		OrderType:   TypeMarket,
		Side:        side,
		Quantity:    quantity,
		OriginalQty: quantity,
		Price:       fpdecimal.Zero,
		Canceled:    false,
	}
}

// NewMarketQuoteOrder creates new constant object Order, but quantity is in Quote mode
func NewMarketQuoteOrder(orderID string, side Side, quantity fpdecimal.Decimal) *Order {

	if quantity.LessThanOrEqual(fpdecimal.Zero) {
		panic(ErrInvalidQuantity)
	}

	return &Order{
		Id:          orderID,
		OrderType:   TypeMarket,
		Side:        side,
		Quantity:    quantity,
		OriginalQty: quantity,
		Price:       fpdecimal.Zero,
		Canceled:    false,
		IsQuote:     true,
	}
}

// NewLimitOrder creates new constant object Order
func NewLimitOrder(orderID string, side Side, quantity, price fpdecimal.Decimal, tif TIF, oco string) *Order {

	if quantity.LessThanOrEqual(fpdecimal.Zero) {
		panic(ErrInvalidQuantity)
	}

	if price.LessThanOrEqual(fpdecimal.Zero) {
		panic(ErrInvalidPrice)
	}

	if tif != "" && tif != GTC && tif != FOK && tif != IOC {
		panic(ErrInvalidTif)
	}

	return &Order{
		Id:          orderID,
		OrderType:   TypeLimit,
		Side:        side,
		Quantity:    quantity,
		OriginalQty: quantity,
		Price:       price,
		Canceled:    false,
		Oco:         oco,
		Tif:         tif,
	}
}

// NewStopLimitOrder creates new constant object Order
func NewStopLimitOrder(orderID string, side Side, quantity, price, stop fpdecimal.Decimal, oco string) *Order {

	if quantity.LessThanOrEqual(fpdecimal.Zero) {
		panic(ErrInvalidQuantity)
	}

	if price.LessThanOrEqual(fpdecimal.Zero) || stop.LessThanOrEqual(fpdecimal.Zero) {
		panic(ErrInvalidPrice)
	}

	return &Order{
		Id:          orderID,
		OrderType:   TypeStopLimit,
		Side:        side,
		Quantity:    quantity,
		OriginalQty: quantity,
		Price:       price,
		Canceled:    false,
		Stop:        stop,
		Oco:         oco,
	}
}

// GetID returns OrderID field copy
func (o *Order) GetID() string {
	return o.Id
}

// GetSide returns side of the Order
func (o *Order) GetSide() Side {
	return o.Side
}

// GetIsQuote returns isQuote field copy
func (o *Order) GetIsQuote() bool {
	return o.IsQuote
}

// GetQuantity returns Quantity field copy
func (o *Order) GetQuantity() fpdecimal.Decimal {
	return o.Quantity
}

// GetOriginalQty returns originalQty field copy
func (o *Order) GetOriginalQty() fpdecimal.Decimal {
	return o.OriginalQty
}

// SetQuantity set Quantity field
func (o *Order) SetQuantity(quantity fpdecimal.Decimal) {
	o.Quantity = quantity
}

// DecreaseQuantity set Quantity field
func (o *Order) DecreaseQuantity(quantity fpdecimal.Decimal) {
	o.Quantity = o.Quantity.Sub(quantity)
}

// GetPrice returns Price field copy
func (o *Order) GetPrice() fpdecimal.Decimal {
	return o.Price
}

// GetStopPrice returns Price field copy
func (o *Order) GetStopPrice() fpdecimal.Decimal {
	return o.Stop
}

// GetOCO returns reference ID
func (o *Order) GetOCO() string {
	return o.Oco
}

// GetTIF returns tif field
func (o *Order) GetTIF() TIF {
	return o.Tif
}

// IsCanceled returns Canceled status
func (o *Order) IsCanceled() bool {
	return o.Canceled
}

// Cancel set Canceled status
func (o *Order) Cancel() bool {
	o.Canceled = true
	return o.Canceled
}

// IsMarketOrder returns true if Order is MARKET
func (o *Order) IsMarketOrder() bool {
	return o.OrderType == TypeMarket
}

// IsLimitOrder returns true if Order is LIMIT
func (o *Order) IsLimitOrder() bool {
	return o.OrderType == TypeLimit
}

// IsStopOrder returns true if Order is STOP-LIMIT
func (o *Order) IsStopOrder() bool {
	return o.OrderType == TypeStopLimit
}

// ActivateStopOrder transforms Stop-Order into Order
func (o *Order) ActivateStopOrder() {

	if !o.IsStopOrder() {
		panic("Order isn't Stop")
	}

	o.Stop = fpdecimal.Zero

	o.OrderType = TypeLimit
}

// SetMaker sets Maker role
func (o *Order) SetMaker() {
	o.Role = MAKER
}

// SetTaker sets Taker role
func (o *Order) SetTaker() {
	o.Role = TAKER
}

// GetRole returns role of Order
func (o *Order) GetRole() Role {
	if o.Role == MAKER {
		return MAKER
	}

	return TAKER
}

// ToSimple returns TradeOrder
func (o *Order) ToSimple() *TradeOrder {
	return &TradeOrder{
		OrderID:  o.GetID(),
		Role:     o.GetRole(),
		IsQuote:  o.GetIsQuote(),
		Quantity: o.GetQuantity(),
		Price:    o.GetPrice(),
	}
}

// String implements Stringer interface
func (o *Order) String() string {
	j, _ := json.Marshal(o.ToSimple())
	return string(j)
}