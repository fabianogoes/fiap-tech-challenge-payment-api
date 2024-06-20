package usecases

import (
	"github.com/fabianogoes/fiap-payment/domain/entities"
	"github.com/fabianogoes/fiap-payment/domain/ports"
	"time"
)

type PaymentService struct {
	paymentRepository ports.PaymentRepositoryPort
	restaurantClient  ports.RestaurantClientPort
}

func NewPaymentService(rep ports.PaymentRepositoryPort, client ports.RestaurantClientPort) *PaymentService {
	return &PaymentService{
		paymentRepository: rep,
		restaurantClient:  client,
	}
}

func (c *PaymentService) GetPaymentById(id string) (*entities.Payment, error) {
	return c.paymentRepository.GetPaymentById(id)
}

func (c *PaymentService) CreatePayment(orderID uint, method string, value float64, date time.Time) (*entities.Payment, error) {
	payment := &entities.Payment{
		OrderID: orderID,
		Date:    date,
		Value:   value,
		Method:  entities.ToPaymentMethod(method),
	}
	return c.paymentRepository.CreatePayment(payment)
}

func (c *PaymentService) UpdatePayment(id string, status string) (*entities.Payment, error) {
	payment, err := c.GetPaymentById(id)
	if err != nil {
		return nil, err
	}

	_, err = c.paymentRepository.UpdateStatus(payment.ID, status)
	if err != nil {
		return nil, err
	}

	err = c.restaurantClient.Webhook(payment.OrderID, status)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
