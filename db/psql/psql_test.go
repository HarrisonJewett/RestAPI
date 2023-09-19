package psql

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/REST-API-Test/types"
	"github.com/stretchr/testify/assert"
)

func setup() (*StoreDB, *sql.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()

	return &StoreDB{
		handler: db,
	}, db, mock
}

func TestPlaceOrderSuccess(t *testing.T) {
	storeDb, sqlDB, mock := setup()

	items := []types.Product{
		{
			ID:       "0",
			Quantity: 1.0,
		},
		{
			ID:       "1",
			Quantity: 100.0,
		},
	}
	defer sqlDB.Close()

	mock.ExpectBegin()

	mock.ExpectExec(`INSERT INTO orders`).
		WithArgs(sqlmock.AnyArg(), "1", sqlmock.AnyArg(), "ready").
		WillReturnResult(sqlmock.NewResult(1, 1))

	quan := float32(1)
	mock.ExpectExec("INSERT INTO products_ordered").
		WithArgs(sqlmock.AnyArg(), "0", quan).
		WillReturnResult(sqlmock.NewResult(1, 1))

	quan = float32(100)
	mock.ExpectExec("INSERT INTO products_ordered").
		WithArgs(sqlmock.AnyArg(), "1", quan).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	storeDb.PlaceOrder(items, "1")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetCustomerOrders(t *testing.T) {
	storeDb, sqlDB, mock := setup()

	defer sqlDB.Close()

	timeCheck := time.Now()

	mock.ExpectQuery(`SELECT order_id, order_time, order_status FROM orders`).
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"order_id", "order_time", "order_status"}).
			AddRow("0", timeCheck, "ready"))

	var orders []types.Order
	var err error
	if orders, err = storeDb.GetCustomerOrders("1"); err != nil {
		assert.FailNow(t, "Failed to get Customer Orders")
	}

	assert.Equal(t, orders[0].OrderTime, timeCheck, "Wrong Time")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetBucket(t *testing.T) {
	storeDb, sqlDB, mock := setup()

	defer sqlDB.Close()

	layout := "2006-01-02"
	start, err := time.Parse(layout, "2021-05-01")
	if err != nil {
		fmt.Println(err)
	}
	end, err := time.Parse(layout, "2021-06-01")
	if err != nil {
		fmt.Println(err)
	}

	firstQuan := 100
	secondQuan := 1000

	mock.ExpectQuery(`SELECT product_id, quantity, order_time FROM orders`).
		WithArgs(start, end).
		WillReturnRows(sqlmock.NewRows([]string{"product_id", "quantity", "order_time"}).
			AddRow("1", firstQuan, start.Add(100)).
			AddRow("1", secondQuan, start.Add(1000)))

	buckets, err := storeDb.GetBreakdown(start, end, "day")
	if err != nil {
		assert.FailNow(t, "Failed to get buckets. %+v", err)
	}

	assert.Equal(t, float32(firstQuan+secondQuan), buckets[0].Products["1"])

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%s", err)
	}
}
