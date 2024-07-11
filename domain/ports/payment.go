package ports

import (
	"time"

	"github.com/fabianogoes/fiap-payment/domain/entities"
)

type PaymentUseCasePort interface {
	GetPaymentById(id string) (*entities.Payment, error)
	CreatePayment(orderID uint, method string, value float64, date time.Time) (*entities.Payment, error)
	UpdatePayment(id string, status string, method string) (*entities.Payment, error)
}

type PaymentRepositoryPort interface {
	GetPaymentById(id string) (*entities.Payment, error)
	CreatePayment(payment *entities.Payment) (*entities.Payment, error)
	UpdateStatus(id string, status string, method string) (*entities.Payment, error)
}
