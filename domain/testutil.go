package domain

import (
	"fmt"
	"time"

	"github.com/fabianogoes/fiap-payment/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

var PaymentIdSuccess = uuid.NewString()
var PaymentIdFail = "4530c2d1-bcd4-439a-a740-584f7aa039cc"
var OrderIdFail = uint(2)
var JsonDateTimeLayout = "2006-01-02T15:04:05"
var DateTimeNowBr = time.Date(
	time.Now().Year(),
	time.Now().Month(),
	time.Now().Day(),
	time.Now().Hour(),
	time.Now().Minute(),
	time.Now().Second(),
	0,
	time.Local)

func BuildPayment() *entities.Payment {
	date := DateTimeNowBr
	return &entities.Payment{
		ID:        PaymentIdSuccess,
		OrderID:   1,
		Date:      date,
		Method:    entities.PaymentMethodCreditCard,
		Status:    entities.PaymentStatusPending,
		Value:     100_00,
		CreatedAt: time.Now().Truncate(time.Second),
	}
}

func BuildPaymentCreditPaid() *entities.Payment {
	payment := *BuildPayment()
	payment.Method = entities.PaymentMethodCreditCard
	payment.Status = entities.PaymentStatusPaid
	return &payment
}

func BuildPaymentCreditError() *entities.Payment {
	return &entities.Payment{
		ID:        PaymentIdFail,
		OrderID:   2,
		Date:      DateTimeNowBr,
		Method:    entities.PaymentMethodCreditCard,
		Status:    entities.PaymentStatusPending,
		Value:     100_00,
		CreatedAt: time.Now().Truncate(time.Second),
	}
}

type PaymentRepositoryMock struct {
	mock.Mock
}

func (p *PaymentRepositoryMock) GetPaymentById(id string) (*entities.Payment, error) {
	args := p.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Payment), args.Error(1)
	}
	return nil, args.Error(1)
}

func (p *PaymentRepositoryMock) CreatePayment(payment *entities.Payment) (*entities.Payment, error) {
	args := p.Called(payment)

	if payment.ID == PaymentIdFail || payment.OrderID == OrderIdFail {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Payment), args.Error(1)
}

func (p *PaymentRepositoryMock) UpdateStatus(id string, status string, method string) (*entities.Payment, error) {
	args := p.Called(id, status)
	return nil, args.Error(1)
}

type RestaurantClientMock struct {
	mock.Mock
}

type PaymentWebhookRequest struct {
	Status        string `json:"status"`
	PaymentMethod string `json:"paymentMethod"`
	ErrorReason   string `json:"errorReason,omitempty"`
}

func (p *RestaurantClientMock) Webhook(orderID uint, status string, method string) error {
	fmt.Printf("Sending webhook request for order %d with status %s\n", orderID, status)
	return nil
}
