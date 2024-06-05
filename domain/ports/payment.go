package ports

import "github.com/fabianogoes/fiap-payment/domain/entities"

type PaymentUseCasePort interface {
	GetPaymentById(id string) (*entities.Payment, error)
	CreatePayment(orderID uint, method string, value float64) (*entities.Payment, error)
}

type PaymentRepositoryPort interface {
	GetPaymentById(id string) (*entities.Payment, error)
	CreatePayment(payment *entities.Payment) (*entities.Payment, error)
}
