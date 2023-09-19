package http

import (
	"github.com/REST-API-Test/usecase"
	"github.com/gin-gonic/gin"
)

type StoreServer struct {
	http    *gin.Engine
	usecase usecase.IUsecase
}

func NewHttpServer(u usecase.IUsecase) *StoreServer {
	ginInstance := setupGin()

	server := &StoreServer{
		http:    ginInstance,
		usecase: u,
	}

	server.registerEndpoints()

	return server
}

func (s *StoreServer) registerEndpoints() {
	v1 := s.http.Group("/v1")
	{
		v1.GET("/ping", s.Ping)
		v1.GET("/orderhistory", s.OrderHistory)
		v1.GET("/orders/:customer_id", s.GetCustomerOrders)
		v1.POST("/placeorder", s.PlaceOrder)
	}
}

func setupGin() *gin.Engine {
	return gin.Default()
}

func (s *StoreServer) Start() {
	s.http.Run(":8080")
}
