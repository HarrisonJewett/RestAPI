package psql

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/REST-API-Test/types"
)

type StoreDB struct {
	handler *sql.DB
}

func NewPsql() *StoreDB {
	return &StoreDB{
		handler: createDBConn(),
	}

}

func createDBConn() *sql.DB {
	db, err := sql.Open("postgres", "")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

func (s *StoreDB) GetCustomerOrders(customerID string) ([]types.Order, error) {
	rows, err := s.handler.Query("SELECT order_id, order_time, order_status FROM orders WHERE customer_id = $1 ORDER BY order_time DESC", customerID)
	if err != nil {
		return nil, err
	}
	orders := make([]types.Order, 0)
	for rows.Next() {
		o := types.Order{}
		err := rows.Scan(&o.OrderID, &o.OrderTime, &o.OrderStatus)
		if err != nil {
			rows.Close()
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

func (s *StoreDB) GetBreakdown(start time.Time, end time.Time, status string) ([]types.DateBucket, error) {
	rows, err := s.handler.Query("SELECT product_id, quantity, order_time FROM orders o left join products_ordered p on o.order_id = p.order_id where order_time BETWEEN $1 AND $2 ORDER BY order_time asc", start, end)
	if err != nil {
		return nil, err
	}
	var duration int
	switch status {
	case "day":
		duration = 24
		break
	case "week":
		duration = 168
		break
	case "month":
		duration = 720
		break
	default:
		duration = 24
		break
	}
	bucketCollection := make([]types.DateBucket, 0)
	var previousOrderTime string
	var bucket types.DateBucket
	var curBucketDate string
	bucket.Products = make(map[string]float32)
	for rows.Next() {
		var id string
		var quantity float32
		var ordertime time.Time
		err = rows.Scan(&id, &quantity, &ordertime)
		if err != nil {
			rows.Close()
			return nil, err
		}

		bucketTime := ordertime.UTC().Truncate(time.Hour * time.Duration(duration))
		curBucketDate = bucketTime.UTC().Format("2006-01-02")
		bucket.Date = previousOrderTime

		if previousOrderTime != curBucketDate {
			if len(previousOrderTime) > 0 {

				bucketCollection = append(bucketCollection, bucket)
				bucket = types.DateBucket{}
				bucket.Products = make(map[string]float32)
			}
			previousOrderTime = curBucketDate
		}
		bucket.Products[id] += quantity
	}
	bucketCollection = append(bucketCollection, bucket)
	return bucketCollection, nil
}

func (s *StoreDB) PlaceOrder(items []types.Product, customerID string) error {
	tx, err := s.handler.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	id := uuid.New()

	_, err = tx.Exec("INSERT INTO orders (order_id, customer_id, order_time, order_status) VALUES ($1, $2, $3, $4)", id, customerID, time.Now(), "ready")
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	for _, item := range items {
		_, err = tx.Exec("INSERT INTO products_ordered (order_id, product_id, quantity) VALUES ($1, $2, $3)", id, item.ID, item.Quantity)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
