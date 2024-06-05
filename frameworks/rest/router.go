package rest

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	paymentHandler *PaymentHandler,
) (*Router, error) {
	router := gin.Default()

	router.GET("/", Welcome)
	router.GET("/health", Health)
	router.GET("/env", Environment)

	customers := router.Group("/payments")
	{
		customers.GET("/:id", paymentHandler.GetPayment)
		customers.POST("/", paymentHandler.CreatePayment)
	}

	return &Router{
		router,
	}, nil
}
