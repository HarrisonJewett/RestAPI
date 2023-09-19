package types

import "time"

type Order struct {
	CustomerID  string    `json:"customer_id"`
	OrderID     string    `json:"order_id"`
	OrderStatus string    `json:"order_status"`
	OrderTime   time.Time `json:"order_time"`
}

//order id
//customer id
//order status
