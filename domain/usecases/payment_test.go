package usecases

import (
	"errors"
	"github.com/fabianogoes/fiap-payment/domain"
	"github.com/fabianogoes/fiap-payment/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_PaymentGetPaymentCreditPaid(t *testing.T) {
	paymentCreditPaid := domain.BuildPaymentCreditPaid()
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	paymentRepositoryMock.On("GetPaymentById", mock.Anything).Return(paymentCreditPaid, nil)

	service := NewPaymentService(paymentRepositoryMock, new(domain.ClientAdapterMock))

	payment, err := service.GetPaymentById(mock.Anything)
	assert.Nil(t, err)
	assert.NotNil(t, payment)
	assert.Equal(t, entities.PaymentMethodCreditCard, payment.Method)
	assert.Equal(t, entities.PaymentStatusPaid, payment.Status)

	paymentRepositoryMock.AssertExpectations(t)
}

func Test_CreatePayment(t *testing.T) {
	paymentCreditPending := domain.BuildPaymentCreditPending()
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	paymentRepositoryMock.On("CreatePayment", mock.Anything).Return(paymentCreditPending, nil)

	service := NewPaymentService(paymentRepositoryMock, new(domain.ClientAdapterMock))

	payment, err := service.CreatePayment(
		paymentCreditPending.OrderID,
		paymentCreditPending.Method.ToString(),
		paymentCreditPending.Value,
		paymentCreditPending.Date)

	assert.Nil(t, err)
	assert.NotNil(t, payment)
	assert.Equal(t, entities.PaymentMethodCreditCard, payment.Method)
	assert.Equal(t, entities.PaymentStatusPending, payment.Status)

	paymentRepositoryMock.AssertExpectations(t)
}

func Test_UpdateSuccess(t *testing.T) {
	paymentCreditPaid := domain.BuildPaymentCreditPaid()
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	paymentRepositoryMock.On("GetPaymentById", mock.Anything).Return(paymentCreditPaid, nil)
	paymentRepositoryMock.On("UpdateStatus", mock.Anything, mock.Anything).Return(paymentCreditPaid, nil)

	restaurantClientPort := new(domain.ClientAdapterMock)
	restaurantClientPort.On("Webhook", mock.Anything, mock.Anything).Return(nil)

	service := NewPaymentService(paymentRepositoryMock, restaurantClientPort)

	payment, err := service.UpdatePayment(mock.Anything, mock.Anything)

	assert.Nil(t, err)
	assert.NotNil(t, payment)
	assert.Equal(t, paymentCreditPaid, payment)

	paymentRepositoryMock.AssertExpectations(t)
}

func Test_UpdatePaymentNotFound(t *testing.T) {
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	errNoDocumentsResult := errors.New("mongo: no documents in result")
	paymentID := domain.PaymentIdFail
	paymentRepositoryMock.On("GetPaymentById", paymentID).Return(nil, errNoDocumentsResult)

	service := NewPaymentService(paymentRepositoryMock, new(domain.ClientAdapterMock))

	payment, err := service.UpdatePayment(paymentID, entities.PaymentStatusPaid.ToString())

	assert.NotNil(t, err)
	assert.Nil(t, payment)
	assert.Equal(t, errNoDocumentsResult, err)

	paymentRepositoryMock.AssertExpectations(t)
}
