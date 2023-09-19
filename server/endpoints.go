package http

import (
	"log"

	"github.com/REST-API-Test/types"
	"github.com/gin-gonic/gin"
)

// Ping ...
func (h *StoreServer) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})

}

// OrderHistory ...
func (h *StoreServer) OrderHistory(c *gin.Context) {
	var r types.DateRange
	if c.ShouldBindQuery(&r) != nil {
		c.JSON(400, gin.H{
			"error": "invalid query params",
		})
		return
	}

	breakdown, err := h.usecase.OrderHistory(r)
	if err != nil {
		c.JSON(404, nil)
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"orders":  breakdown,
	})
}

// GetCustomerOrder ...
func (h *StoreServer) GetCustomerOrders(c *gin.Context) {
	var orderParams types.CustomerOrderParams
	orderParams.CustomerID = c.Param("customer_id")

	if c.Query("asc") == "" || c.Query("asc") == "true" {
		orderParams.Asc = true
	} else {
		orderParams.Asc = false
	}

	orders, err := h.usecase.GetCustomerOrders(orderParams)
	if err != nil {
		c.JSON(404, nil)
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"orders":  orders,
	})
}

func (h *StoreServer) PlaceOrder(c *gin.Context) {
	var order types.OrderRequest
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{
			"error": "Bad Object",
		})
		log.Println(err)
		return
	}

	err := h.usecase.PlaceOrder(order)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Could not place order",
		})
		log.Println(err)
		return
	}
	c.JSON(200, nil)
}
