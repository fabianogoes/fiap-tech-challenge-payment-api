package usecases

import (
	"github.com/fabianogoes/fiap-payment/domain/entities"
	"github.com/fabianogoes/fiap-payment/domain/ports"
	"time"
)

type PaymentService struct {
	paymentRepository ports.PaymentRepositoryPort
}

func NewPaymentService(rep ports.PaymentRepositoryPort) *PaymentService {
	return &PaymentService{
		paymentRepository: rep,
	}
}

func (c *PaymentService) GetPaymentById(id string) (*entities.Payment, error) {
	return c.paymentRepository.GetPaymentById(id)
}

func (c *PaymentService) CreatePayment(orderID uint, method string, value float64) (*entities.Payment, error) {
	payment := &entities.Payment{
		OrderID: orderID,
		Date:    time.Now(),
		Value:   value,
		Method:  entities.ToPaymentMethod(method),
	}
	return c.paymentRepository.CreatePayment(payment)
}
