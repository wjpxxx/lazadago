package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetStatusToReadyToShipOrderItemsResponseEntity struct{
    OrderItemId	int	`json:"order_item_id"`
    PurchaseOrderId	int	`json:"purchase_order_id"`
    PurchaseOrderNumber	string	`json:"purchase_order_number"`
}
func (g SetStatusToReadyToShipOrderItemsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
