package usecases

import (
	"fmt"
	"time"

	"github.com/fabianogoes/fiap-payment/domain/entities"
	"github.com/fabianogoes/fiap-payment/domain/ports"
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

func (c *PaymentService) UpdatePayment(id string, status string, method string) (*entities.Payment, error) {
	fmt.Printf("update payemnt id %s status %s method %s\n", id, status, method)
	payment, err := c.GetPaymentById(id)
	if err != nil {
		return nil, err
	}

	_, err = c.paymentRepository.UpdateStatus(payment.ID, status, method)
	if err != nil {
		return nil, fmt.Errorf("error updating payment status: %v", err)
	}

	err = c.restaurantClient.Webhook(payment.OrderID, status, method)
	if err != nil {
		return nil, fmt.Errorf("error calling restaurant webhook: %v", err)
	}

	fmt.Printf("payment %s status %s method %s updated successfully", id, status, method)
	return payment, nil
}
