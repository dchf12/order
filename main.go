package main

import (
	"fmt"

	"github.com/dchf12/order/order"
)

// Order は注文の状態を表すインターフェイス
type Order interface {
	ProceedOrder() (*order.ConfirmedOrder, error)
	CancelOrder() (*order.CancelledOrder, error)
	ShipOrder() (*order.ShippedOrder, error)
}

var _ Order = (*order.UnconfirmedOrder)(nil)
var _ Order = (*order.ConfirmedOrder)(nil)
var _ Order = (*order.CancelledOrder)(nil)
var _ Order = (*order.ShippedOrder)(nil)

func main() {
	order := &order.UnconfirmedOrder{} // 初期状態は未確定の注文
	confirmedOrder, err := order.ProceedOrder()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = confirmedOrder.ShipOrder()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 以下の遷移はエラーハンドリングを含む
	cancelledOrder, err := confirmedOrder.CancelOrder()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Final state: %+v\n", cancelledOrder)
}
