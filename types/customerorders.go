package types

type CustomerOrderParams struct {
	CustomerID string `json:"customer_id"`
	Asc        bool
}
