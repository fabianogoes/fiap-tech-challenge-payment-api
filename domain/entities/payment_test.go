package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PaymentStatus(t *testing.T) {
	assert.Equal(t, "PENDING", PaymentStatusPending.ToString())
	assert.Equal(t, "PAID", PaymentStatusPaid.ToString())
	assert.Equal(t, "REVERSED", PaymentStatusReversed.ToString())
	assert.Equal(t, "CANCELED", PaymentStatusCanceled.ToString())
	assert.Equal(t, "ERROR", PaymentStatusError.ToString())
	assert.Equal(t, "NONE", PaymentStatusNone.ToString())
}

func Test_PaymentMethod(t *testing.T) {
	assert.Equal(t, "CREDIT_CARD", PaymentMethodCreditCard.ToString())
	assert.Equal(t, "DEBIT_CARD", PaymentMethodDebitCard.ToString())
	assert.Equal(t, "MONEY", PaymentMethodMoney.ToString())
	assert.Equal(t, "PIX", PaymentMethodPIX.ToString())
	assert.Equal(t, "NONE", PaymentMethodNone.ToString())

	assert.Equal(t, PaymentMethodCreditCard, ToPaymentMethod("CREDIT_CARD"))
	assert.Equal(t, PaymentMethodDebitCard, ToPaymentMethod("DEBIT_CARD"))
	assert.Equal(t, PaymentMethodMoney, ToPaymentMethod("MONEY"))
	assert.Equal(t, PaymentMethodPIX, ToPaymentMethod("PIX"))
	assert.Equal(t, PaymentMethodNone, ToPaymentMethod("NONE"))
}
