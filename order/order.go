package order

import (
	"errors"
	"fmt"
)

type Order struct {
	Status string
}

// UnconfirmedOrder は未確定の注文状態を表す
type UnconfirmedOrder struct {
	Order
}

// ProceedOrder は注文を確定状態に遷移させる
func (o *UnconfirmedOrder) ProceedOrder() (*ConfirmedOrder, error) {
	fmt.Println("Order confirmed")
	return &ConfirmedOrder{}, nil
}

// CancelOrder と ShipOrder は未確定の注文では操作できないため、エラーを返す
func (o *UnconfirmedOrder) CancelOrder() (*CancelledOrder, error) {
	return nil, errors.New("cannot cancel an unconfirmed order")
}

func (o *UnconfirmedOrder) ShipOrder() (*ShippedOrder, error) {
	return nil, errors.New("cannot ship an unconfirmed order")
}

// ConfirmedOrder は確定した注文状態を表す
type ConfirmedOrder struct{}

func (c *ConfirmedOrder) ProceedOrder() (*ConfirmedOrder, error) {
	return nil, errors.New("order is already confirmed")
}

func (c *ConfirmedOrder) CancelOrder() (*CancelledOrder, error) {
	fmt.Println("Order cancelled")
	return &CancelledOrder{}, nil
}

func (c *ConfirmedOrder) ShipOrder() (*ShippedOrder, error) {
	fmt.Println("Order shipped")
	return &ShippedOrder{}, nil
}

// CancelledOrder はキャンセルされた注文状態を表す
type CancelledOrder struct{}

func (c *CancelledOrder) ProceedOrder() (*ConfirmedOrder, error) {
	return nil, errors.New("cannot proceed a cancelled order")
}

func (c *CancelledOrder) CancelOrder() (*CancelledOrder, error) {
	return nil, errors.New("order is already cancelled")
}

func (c *CancelledOrder) ShipOrder() (*ShippedOrder, error) {
	return nil, errors.New("cannot ship a cancelled order")
}

// ShippedOrder は発送された注文状態を表す
type ShippedOrder struct{}

func (s *ShippedOrder) ProceedOrder() (*ConfirmedOrder, error) {
	return nil, errors.New("cannot proceed a shipped order")
}

func (s *ShippedOrder) CancelOrder() (*CancelledOrder, error) {
	return nil, errors.New("cannot cancel a shipped order")
}

func (s *ShippedOrder) ShipOrder() (*ShippedOrder, error) {
	return nil, errors.New("order is already shipped")
}
