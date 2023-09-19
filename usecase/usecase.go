package usecase

import (
	"github.com/REST-API-Test/db"
	"github.com/REST-API-Test/types"
)

type Usecase struct {
	db db.IStoreDB
}

func NewUsecase(db db.IStoreDB) *Usecase {
	return &Usecase{
		db: db,
	}
}

func (u *Usecase) OrderHistory(r types.DateRange) ([]types.DateBucket, error) {
	return u.db.GetBreakdown(r.Start, r.End, r.Type)
}

func (u *Usecase) GetCustomerOrders(r types.CustomerOrderParams) ([]types.Order, error) {
	return u.db.GetCustomerOrders(r.CustomerID)
}

func (u *Usecase) PlaceOrder(r types.OrderRequest) error {
	return u.db.PlaceOrder(r.Items, r.CustomerID)
}
