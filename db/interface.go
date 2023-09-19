package db

import (
	"time"

	"github.com/REST-API-Test/types"
)

type IStoreDB interface {
	PlaceOrder([]types.Product, string) error
	GetCustomerOrders(string) ([]types.Order, error)
	GetBreakdown(time.Time, time.Time, string) ([]types.DateBucket, error)
}
