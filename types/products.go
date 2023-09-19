package types

type Product struct {
	ID       string  `json:"id"`
	Quantity float32 `json:"quantity"`
}

type OrderRequest struct {
	CustomerID string    `json:"customer_id" binding:"required"`
	Items      []Product `json:"items" binding:"required"`
}

type Breakdown struct {
	Items map[string]float32 `json:"items"`
}

type DateBucket struct {
	Date     string             `json:"date"`
	Products map[string]float32 `json:"products"`
}
