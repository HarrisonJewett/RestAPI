package usecase

import "github.com/REST-API-Test/types"

type IUsecase interface {
	PlaceOrder(types.OrderRequest) error
	GetCustomerOrders(types.CustomerOrderParams) ([]types.Order, error)
	OrderHistory(types.DateRange) ([]types.DateBucket, error)
}
