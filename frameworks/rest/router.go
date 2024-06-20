package rest

import (
	"github.com/fabianogoes/fiap-payment/frameworks/rest/payment"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	paymentHandler *payment.PaymentHandler,
) (*Router, error) {
	router := gin.Default()

	router.GET("/", Welcome)
	router.GET("/health", Health)
	router.GET("/env", Environment)

	payments := router.Group("/payments")
	{
		payments.GET("/:id", paymentHandler.GetPayment)
		payments.POST("/", paymentHandler.CreatePayment)
		payments.POST("/:id/status", paymentHandler.UpdateStatus)
	}

	return &Router{
		router,
	}, nil
}
