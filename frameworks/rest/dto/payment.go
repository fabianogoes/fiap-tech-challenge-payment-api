package dto

import (
	"github.com/fabianogoes/fiap-payment/domain/entities"
)

type GetPaymentResponse struct {
	ID          string  `json:"id"`
	OrderID     uint    `json:"orderId"`
	Date        string  `json:"date"`
	Method      string  `json:"method"`
	ErrorReason string  `json:"errorReason"`
	Status      string  `json:"status"`
	Value       float64 `json:"value"`
}

func ToPaymentResponse(entity *entities.Payment) GetPaymentResponse {
	return GetPaymentResponse{
		ID:          entity.ID,
		OrderID:     entity.OrderID,
		Date:        entity.Date.Format("2006-01-02"),
		Method:      entity.Method.ToString(),
		ErrorReason: entity.ErrorReason,
		Status:      entity.Status.ToString(),
		Value:       entity.Value,
	}
}

type CreatePaymentRequest struct {
	OrderID uint    `json:"orderId"`
	Method  string  `json:"method"`
	Value   float64 `json:"value"`
}
