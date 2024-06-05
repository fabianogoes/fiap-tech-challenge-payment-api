package dbo

import (
	"github.com/fabianogoes/fiap-payment/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Payment is a Database Object for payment
type Payment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	OrderID     uint               `bson:"orderID"`
	Date        time.Time          `bson:"date"`
	Method      string             `bson:"method"`
	ErrorReason string             `bson:"errorReason"`
	Status      string             `bson:"status"`
	Value       float64            `bson:"value"`
}

func (p *Payment) ToEntity() *entities.Payment {
	return &entities.Payment{
		ID:          p.ID.Hex(),
		OrderID:     p.OrderID,
		Date:        p.Date,
		Method:      p.toPaymentMethod(),
		ErrorReason: p.ErrorReason,
		Status:      p.toPaymentStatus(),
		Value:       p.Value,
	}
}

func ToPaymentDBO(payment *entities.Payment) Payment {
	return Payment{
		CreatedAt:   payment.CreatedAt,
		UpdatedAt:   payment.UpdatedAt,
		OrderID:     payment.OrderID,
		Date:        payment.Date,
		Method:      payment.Method.ToString(),
		ErrorReason: payment.ErrorReason,
		Status:      payment.Status.ToString(),
		Value:       payment.Value,
	}
}

func (p *Payment) toPaymentStatus() entities.PaymentStatus {
	switch p.Status {
	case "PENDING":
		return entities.PaymentStatusPending
	case "PAID":
		return entities.PaymentStatusPaid
	case "REVERSED":
		return entities.PaymentStatusReversed
	case "CANCELED":
		return entities.PaymentStatusCanceled
	case "ERROR":
		return entities.PaymentStatusError
	default:
		return entities.PaymentStatusNone
	}
}

func (p *Payment) toPaymentMethod() entities.PaymentMethod {
	switch p.Method {
	case "CREDIT_CARD":
		return entities.PaymentMethodCreditCard
	case "DEBIT_CARD":
		return entities.PaymentMethodDebitCard
	case "MONEY":
		return entities.PaymentMethodMoney
	case "PIX":
		return entities.PaymentMethodPIX
	default:
		return entities.PaymentMethodNone
	}
}
