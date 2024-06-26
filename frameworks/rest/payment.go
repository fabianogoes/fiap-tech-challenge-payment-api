package rest

import (
	"github.com/fabianogoes/fiap-payment/domain"
	"github.com/fabianogoes/fiap-payment/domain/ports"
	"github.com/fabianogoes/fiap-payment/frameworks/rest/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PaymentHandler struct {
	UseCase ports.PaymentUseCasePort
}

func NewPaymentHandler(useCase ports.PaymentUseCasePort) *PaymentHandler {
	return &PaymentHandler{
		UseCase: useCase,
	}
}

func (h *PaymentHandler) GetPayment(c *gin.Context) {
	var err error
	id := c.Param("id")

	payment, err := h.UseCase.GetPaymentById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToPaymentResponse(payment))
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var request dto.CreatePaymentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dateTime, err := time.Parse(domain.JsonDateTimeLayout, request.Date)
	payment, err := h.UseCase.CreatePayment(request.OrderID, request.Method, request.Value, dateTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToPaymentResponse(payment))
}

func (h *PaymentHandler) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	status := c.Query("status")

	payment, err := h.UseCase.UpdatePayment(id, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToPaymentResponse(payment))
}
