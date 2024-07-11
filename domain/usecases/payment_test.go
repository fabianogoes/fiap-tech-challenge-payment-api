package usecases

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/fabianogoes/fiap-payment/domain"
	"github.com/fabianogoes/fiap-payment/domain/entities"
	"github.com/stretchr/testify/mock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPayment(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Payment Suite")
}

var _ = Describe("Payment", func() {
	orderID := uint(1)
	paymentDate := time.Now()
	methodCreditCard := entities.PaymentMethodCreditCard.ToString()
	paymentValue := 100.50

	Context("initially", func() {
		paymentPending := entities.NewPayment(orderID, paymentDate, methodCreditCard, paymentValue)
		repositoryMock := new(domain.PaymentRepositoryMock)
		repositoryMock.On("CreatePayment", mock.Anything).Return(&paymentPending, nil)
		useCase := NewPaymentService(repositoryMock, new(domain.RestaurantClientMock))

		payment, err := useCase.CreatePayment(orderID, methodCreditCard, paymentValue, paymentDate)

		It("has no error on CreatePayment", func() {
			Expect(err).Should(BeNil())
		})

		It("has id not be nil", func() {
			Expect(payment.ID).Should(Not(BeNil()))
		})

		It(fmt.Sprintf("has order id %d", orderID), func() {
			Expect(payment.OrderID).Should(Equal(orderID))
		})

		It(fmt.Sprintf("has paymentDate %v", paymentDate), func() {
			Expect(payment.Date).Should(Equal(paymentDate))
		})

		It(fmt.Sprintf("has methodCreditCard %s", methodCreditCard), func() {
			Expect(payment.Method.ToString()).Should(Equal(methodCreditCard))
		})

		It(fmt.Sprintf("has paymentValue %v", paymentValue), func() {
			Expect(payment.Value).Should(Equal(paymentValue))
		})

		It(fmt.Sprintf("has status %v", entities.PaymentStatusPending), func() {
			Expect(payment.Status).Should(Equal(entities.PaymentStatusPending))
		})

		It("has createdAt not be nil", func() {
			Expect(payment.CreatedAt).Should(Not(BeNil()))
		})
	})

	Context("get payment with paid status", func() {
		paymentCreditPaid := entities.NewPayment(orderID, paymentDate, methodCreditCard, paymentValue)
		paymentCreditPaid.ID = domain.PaymentIdSuccess
		paymentCreditPaid.Status = entities.PaymentStatusPaid

		repositoryMock := new(domain.PaymentRepositoryMock)
		repositoryMock.On("GetPaymentById", paymentCreditPaid.ID).Return(&paymentCreditPaid, nil)
		useCase := NewPaymentService(repositoryMock, new(domain.RestaurantClientMock))

		payment, err := useCase.GetPaymentById(mock.Anything)

		It("has no error on GetPaymentById", func() {
			Expect(err).Should(BeNil())
		})

		It("has not nil payment", func() {
			Expect(payment).ShouldNot(BeNil())
		})

		It(fmt.Sprintf("has payment methodCreditCard %s", paymentCreditPaid.Method.ToString()), func() {
			Expect(payment.Method.ToString()).Should(Equal(methodCreditCard))
		})

		It(fmt.Sprintf("has status %s", paymentCreditPaid.Status.ToString()), func() {
			Expect(payment.Status).Should(Equal(paymentCreditPaid.Status))
		})

	})

	Context("update payment to paid", func() {
		paymentCreditPaid := entities.NewPayment(orderID, paymentDate, methodCreditCard, paymentValue)
		paymentCreditPaid.Status = entities.PaymentStatusPaid

		repositoryMock := new(domain.PaymentRepositoryMock)
		repositoryMock.On("GetPaymentById", mock.Anything).Return(&paymentCreditPaid, nil)
		repositoryMock.On("UpdateStatus", mock.Anything, mock.Anything).Return(&paymentCreditPaid, nil)
		restaurantClientMock := new(domain.RestaurantClientMock)
		useCase := NewPaymentService(repositoryMock, restaurantClientMock)

		restaurantClientMock.On("Webhook", mock.Anything, mock.Anything).Return(nil)
		payment, err := useCase.UpdatePayment(mock.Anything, mock.Anything, mock.Anything)

		It("has no error on UpdatePayment", func() {
			Expect(err).Should(BeNil())
		})

		It("has not nil payment", func() {
			Expect(payment).ShouldNot(BeNil())
		})

		It(fmt.Sprintf("has status %s", paymentCreditPaid.Status.ToString()), func() {
			Expect(payment.Status).Should(Equal(paymentCreditPaid.Status))
		})
	})

	Context("update payment not found error", func() {
		statusPaid := entities.PaymentStatusPaid.ToString()
		methodCreditCard := entities.PaymentMethodCreditCard.ToString()

		repositoryMock := new(domain.PaymentRepositoryMock)
		repositoryMock.
			On("GetPaymentById", mock.Anything).
			Return(nil, errors.New("mongo: no documents in result"))

		useCase := NewPaymentService(repositoryMock, new(domain.RestaurantClientMock))
		payment, err := useCase.UpdatePayment(domain.PaymentIdFail, statusPaid, methodCreditCard)

		It("has error on UpdatePayment not found", func() {
			Expect(err).ShouldNot(BeNil())
		})

		It("has not nil payment", func() {
			Expect(payment).Should(BeNil())
		})

	})

})
