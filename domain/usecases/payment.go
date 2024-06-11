package usecases

import (
	"github.com/fabianogoes/fiap-payment/adapters/restaurant"
	"github.com/fabianogoes/fiap-payment/domain/entities"
	"github.com/fabianogoes/fiap-payment/domain/ports"
	"time"
)

type PaymentService struct {
	paymentRepository ports.PaymentRepositoryPort
	restaurantClient  *restaurant.ClientAdapter
}

func NewPaymentService(rep ports.PaymentRepositoryPort, client *restaurant.ClientAdapter) *PaymentService {
	return &PaymentService{
		paymentRepository: rep,
		restaurantClient:  client,
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

func (c *PaymentService) UpdatePayment(id string, status string) (*entities.Payment, error) {
	_, err := c.paymentRepository.UpdateStatus(id, status)
	if err != nil {
		return nil, err
	}

	payment, err := c.GetPaymentById(id)
	if err != nil {
		return nil, err
	}

	err = c.restaurantClient.Webhook(payment.OrderID, status)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
