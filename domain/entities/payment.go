package entities

import "time"

type PaymentStatus int

const (
	PaymentStatusPending PaymentStatus = iota
	PaymentStatusPaid
	PaymentStatusReversed
	PaymentStatusCanceled
	PaymentStatusError
	PaymentStatusNone
)

func (ps PaymentStatus) ToString() string {
	return [...]string{"PENDING", "PAID", "REVERSED", "CANCELED", "ERROR", "NONE"}[ps]
}

type PaymentMethod int

const (
	PaymentMethodCreditCard PaymentMethod = iota
	PaymentMethodDebitCard
	PaymentMethodMoney
	PaymentMethodPIX
	PaymentMethodNone
)

func (pm PaymentMethod) ToString() string {
	return [...]string{"CREDIT_CARD", "DEBIT_CARD", "MONEY", "PIX", "NONE"}[pm]
}

func ToPaymentMethod(method string) PaymentMethod {
	switch method {
	case "CREDIT_CARD":
		return PaymentMethodCreditCard
	case "DEBIT_CARD":
		return PaymentMethodDebitCard
	case "MONEY":
		return PaymentMethodMoney
	case "PIX":
		return PaymentMethodPIX
	default:
		return PaymentMethodNone
	}
}

type Payment struct {
	ID          string
	OrderID     uint
	Date        time.Time
	Method      PaymentMethod
	ErrorReason string
	Status      PaymentStatus
	Value       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
