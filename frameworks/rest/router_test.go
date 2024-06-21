package rest

import (
	"github.com/fabianogoes/fiap-payment/domain"
	"github.com/fabianogoes/fiap-payment/domain/usecases"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Router(t *testing.T) {
	paymentService := usecases.NewPaymentService(new(domain.PaymentRepositoryMock), new(domain.ClientAdapterMock))
	paymentHandler := NewPaymentHandler(paymentService)
	router, err := NewRouter(paymentHandler)
	assert.Nil(t, err)
	assert.NotNil(t, router)
}
