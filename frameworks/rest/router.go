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

	payments := router.Group("/payments")
	{
		payments.GET("/:id", paymentHandler.GetPayment)
		payments.GET("/order/:id", paymentHandler.GetPaymentByOrderId)
		payments.POST("/", paymentHandler.CreatePayment)
		payments.PUT("/:id/status", paymentHandler.UpdateStatus)
		payments.PUT("/order/:id/status", paymentHandler.UpdateStatusByOrderId)
	}

	return &Router{
		router,
	}, nil
}
