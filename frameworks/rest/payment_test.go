package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabianogoes/fiap-payment/domain"
	"github.com/fabianogoes/fiap-payment/domain/entities"
	"github.com/fabianogoes/fiap-payment/domain/usecases"
	"github.com/fabianogoes/fiap-payment/frameworks"
	"github.com/fabianogoes/fiap-payment/frameworks/rest/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var paymentPaid = domain.BuildPaymentCreditPaid()
var paymentFail = domain.BuildPaymentCreditError()

func Test_GetPaymentByIdPaid(t *testing.T) {
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	paymentRepositoryMock.On("GetPaymentById", paymentPaid.ID).Return(paymentPaid, nil)

	restaurantAdapterMock := new(domain.RestaurantClientMock)
	paymentUseCase := usecases.NewPaymentService(paymentRepositoryMock, restaurantAdapterMock)
	paymentHandler := NewPaymentHandler(paymentUseCase)

	r := frameworks.Setup()
	r.GET("/payments/:id", paymentHandler.GetPayment)
	request, _ := http.NewRequest("GET", fmt.Sprintf("/payments/%s", paymentPaid.ID), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK status response is expected")
	assert.NotNil(t, response.Body, "Body is expected")

	var jsonResponse map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, paymentPaid.ID, jsonResponse["id"])
	assert.Equal(t, float64(paymentPaid.OrderID), jsonResponse["orderId"])
	assert.Equal(t, paymentPaid.Date.Format("2006-01-02"), jsonResponse["date"])
	assert.Equal(t, paymentPaid.Status.ToString(), jsonResponse["status"])
	assert.Equal(t, paymentPaid.Method.ToString(), jsonResponse["method"])

	_, exists := jsonResponse["errorReason"]
	assert.False(t, exists)
}

func Test_GetPaymentByIdFail(t *testing.T) {
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	paymentRepositoryMock.On("GetPaymentById", domain.PaymentIdFail).Return(nil, errors.New("payment not found"))
	paymentUseCase := usecases.NewPaymentService(paymentRepositoryMock, new(domain.RestaurantClientMock))
	paymentHandler := NewPaymentHandler(paymentUseCase)

	r := frameworks.Setup()
	r.GET("/payments/:id", paymentHandler.GetPayment)
	request, _ := http.NewRequest("GET", fmt.Sprintf("/payments/%s", domain.PaymentIdFail), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 404, response.Code, "OK status response is expected")
	assert.NotNil(t, response.Body, "Body is expected")

	var jsonResponse map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	assert.NoError(t, err, "Error should be nil")
}

func Test_GetPaymentByOrderIdPaid(t *testing.T) {
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	paymentRepositoryMock.On("GetPaymentByOrderId", paymentPaid.OrderID).Return(paymentPaid, nil)

	restaurantAdapterMock := new(domain.RestaurantClientMock)
	paymentUseCase := usecases.NewPaymentService(paymentRepositoryMock, restaurantAdapterMock)
	paymentHandler := NewPaymentHandler(paymentUseCase)

	r := frameworks.Setup()
	r.GET("/payments/order/:id", paymentHandler.GetPaymentByOrderId)
	request, _ := http.NewRequest("GET", fmt.Sprintf("/payments/order/%d", paymentPaid.OrderID), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK status response is expected")
	assert.NotNil(t, response.Body, "Body is expected")

	var jsonResponse map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, paymentPaid.ID, jsonResponse["id"])
	assert.Equal(t, float64(paymentPaid.OrderID), jsonResponse["orderId"])
	assert.Equal(t, paymentPaid.Date.Format("2006-01-02"), jsonResponse["date"])
	assert.Equal(t, paymentPaid.Status.ToString(), jsonResponse["status"])
	assert.Equal(t, paymentPaid.Method.ToString(), jsonResponse["method"])

	_, exists := jsonResponse["errorReason"]
	assert.False(t, exists)
}

func Test_GetPaymentByOrderIdBadRequest(t *testing.T) {
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	restaurantAdapterMock := new(domain.RestaurantClientMock)
	paymentUseCase := usecases.NewPaymentService(paymentRepositoryMock, restaurantAdapterMock)
	paymentHandler := NewPaymentHandler(paymentUseCase)

	r := frameworks.Setup()
	r.GET("/payments/order/:id", paymentHandler.GetPaymentByOrderId)
	request, _ := http.NewRequest("GET", fmt.Sprintf("/payments/order/%s", "x"), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 400, response.Code, "Bad Request status response is expected")
}

func Test_GetPaymentByOrderIdFail(t *testing.T) {
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	paymentRepositoryMock.On("GetPaymentByOrderId", mock.Anything).Return(nil, errors.New("payment not found"))
	paymentUseCase := usecases.NewPaymentService(paymentRepositoryMock, new(domain.RestaurantClientMock))
	paymentHandler := NewPaymentHandler(paymentUseCase)

	r := frameworks.Setup()
	r.GET("/payments/order/:id", paymentHandler.GetPaymentByOrderId)
	request, _ := http.NewRequest("GET", fmt.Sprintf("/payments/order/%d", paymentPaid.OrderID), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 404, response.Code, "OK status response is expected")
	assert.NotNil(t, response.Body, "Body is expected")

	var jsonResponse map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	assert.NoError(t, err, "Error should be nil")
}

func Test_CreatePaymentSuccess(t *testing.T) {
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	paymentRepositoryMock.On("CreatePayment", mock.Anything).Return(paymentPaid, nil)

	restaurantAdapterMock := new(domain.RestaurantClientMock)
	paymentUseCase := usecases.NewPaymentService(paymentRepositoryMock, restaurantAdapterMock)
	paymentHandler := NewPaymentHandler(paymentUseCase)

	createPaymentRequest := dto.CreatePaymentRequest{
		OrderID: paymentPaid.OrderID,
		Method:  paymentPaid.Method.ToString(),
		Value:   paymentPaid.Value,
		Date:    paymentPaid.Date.Format(domain.JsonDateTimeLayout),
	}
	jsonRequest, _ := json.Marshal(createPaymentRequest)
	readerRequest := bytes.NewReader(jsonRequest)

	r := frameworks.Setup()
	r.POST("/payments/", paymentHandler.CreatePayment)
	request, _ := http.NewRequest("POST", "/payments/", readerRequest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 201, response.Code, "CREATE response status is expected")
	assert.NotNil(t, response.Body, "Body is expected")

	var jsonResponse map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, paymentPaid.ID, jsonResponse["id"])
	assert.Equal(t, float64(paymentPaid.OrderID), jsonResponse["orderId"])
	assert.Equal(t, paymentPaid.Date.Format("2006-01-02"), jsonResponse["date"])
	assert.Equal(t, paymentPaid.Status.ToString(), jsonResponse["status"])
	assert.Equal(t, paymentPaid.Method.ToString(), jsonResponse["method"])

	_, exists := jsonResponse["errorReason"]
	assert.False(t, exists)
}

func Test_CreatePaymentBadRequest(t *testing.T) {
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	restaurantAdapterMock := new(domain.RestaurantClientMock)
	paymentUseCase := usecases.NewPaymentService(paymentRepositoryMock, restaurantAdapterMock)
	paymentHandler := NewPaymentHandler(paymentUseCase)
	r := frameworks.Setup()
	r.POST("/payments/", paymentHandler.CreatePayment)
	request, _ := http.NewRequest("POST", "/payments/", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 400, response.Code, "Bad Request response status is expected")
}

func Test_CreatePaymentBadRequestDate(t *testing.T) {
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	restaurantAdapterMock := new(domain.RestaurantClientMock)
	paymentUseCase := usecases.NewPaymentService(paymentRepositoryMock, restaurantAdapterMock)
	paymentHandler := NewPaymentHandler(paymentUseCase)

	createPaymentRequest := dto.CreatePaymentRequest{
		OrderID: paymentPaid.OrderID,
		Method:  paymentPaid.Method.ToString(),
		Value:   paymentPaid.Value,
		Date:    "xxx",
	}
	jsonRequest, _ := json.Marshal(createPaymentRequest)
	readerRequest := bytes.NewReader(jsonRequest)

	r := frameworks.Setup()
	r.POST("/payments/", paymentHandler.CreatePayment)
	request, _ := http.NewRequest("POST", "/payments/", readerRequest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 400, response.Code, "Bad Request response status is expected")
}

func Test_CreatePaymentFail(t *testing.T) {
	errNoDocumentsResult := errors.New("error on create payment")
	createPaymentRequest := dto.CreatePaymentRequest{
		OrderID: paymentFail.OrderID,
		Method:  paymentFail.Method.ToString(),
		Value:   paymentFail.Value,
		Date:    domain.DateTimeNowBr.Format(domain.JsonDateTimeLayout),
	}

	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	paymentRepositoryMock.On("CreatePayment", mock.Anything).Return(nil, errNoDocumentsResult)

	restaurantAdapterMock := new(domain.RestaurantClientMock)
	paymentUseCase := usecases.NewPaymentService(paymentRepositoryMock, restaurantAdapterMock)
	paymentHandler := NewPaymentHandler(paymentUseCase)

	jsonRequest, _ := json.Marshal(createPaymentRequest)
	readerRequest := bytes.NewReader(jsonRequest)

	r := frameworks.Setup()
	r.POST("/payments/", paymentHandler.CreatePayment)
	request, _ := http.NewRequest("POST", "/payments/", readerRequest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 500, response.Code, "CREATE response status is expected")
	assert.NotNil(t, response.Body, "Body is expected")

	var jsonResponse map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	assert.Nil(t, err, "Error should be nil")
	_, exists := jsonResponse["error"]
	assert.True(t, exists, "should exists error field into json")
}

func Test_UpdatePaymentSuccess(t *testing.T) {
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	paymentRepositoryMock.On("GetPaymentById", paymentPaid.ID).Return(paymentPaid, nil)
	paymentRepositoryMock.On("UpdateStatus", mock.Anything, mock.Anything, mock.Anything).Return(paymentPaid, nil)

	restaurantAdapterMock := new(domain.RestaurantClientMock)
	paymentUseCase := usecases.NewPaymentService(paymentRepositoryMock, restaurantAdapterMock)
	paymentHandler := NewPaymentHandler(paymentUseCase)
	newStatus := entities.PaymentStatusPaid.ToString()

	r := frameworks.Setup()
	r.POST("/payments/:id/status", paymentHandler.UpdateStatus)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/payments/%s/status?status=%s", paymentPaid.ID, newStatus), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "CREATE response status is expected")
	assert.NotNil(t, response.Body, "Body is expected")

	var jsonResponse map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, paymentPaid.ID, jsonResponse["id"])
	assert.Equal(t, float64(paymentPaid.OrderID), jsonResponse["orderId"])
	assert.Equal(t, paymentPaid.Date.Format("2006-01-02"), jsonResponse["date"])
	assert.Equal(t, paymentPaid.Status.ToString(), jsonResponse["status"])
	assert.Equal(t, paymentPaid.Method.ToString(), jsonResponse["method"])

	_, exists := jsonResponse["errorReason"]
	assert.False(t, exists)
}

func Test_UpdatePaymentFailNotFound(t *testing.T) {
	errPaymentNotFound := errors.New("error on create payment")
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	paymentRepositoryMock.On("GetPaymentById", mock.Anything).Return(nil, errPaymentNotFound)

	paymentUseCase := usecases.NewPaymentService(paymentRepositoryMock, new(domain.RestaurantClientMock))
	paymentHandler := NewPaymentHandler(paymentUseCase)
	newStatus := entities.PaymentStatusPaid.ToString()

	r := frameworks.Setup()
	r.POST("/payments/:id/status", paymentHandler.UpdateStatus)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/payments/%s/status?status=%s", paymentFail.ID, newStatus), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 500, response.Code, "CREATE response status is expected")
	assert.NotNil(t, response.Body, "Body is expected")

	var jsonResponse map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	assert.Nil(t, err, "Error should be nil")
	_, exists := jsonResponse["error"]
	assert.True(t, exists, "should exists error field into json")
}

func Test_UpdatePaymentByOrderId(t *testing.T) {
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	paymentRepositoryMock.On("GetPaymentById", paymentPaid.ID).Return(paymentPaid, nil)
	paymentRepositoryMock.On("GetPaymentByOrderId", paymentPaid.OrderID).Return(paymentPaid, nil)
	paymentRepositoryMock.On("UpdateStatus", mock.Anything, mock.Anything, mock.Anything).Return(paymentPaid, nil)

	restaurantAdapterMock := new(domain.RestaurantClientMock)
	paymentUseCase := usecases.NewPaymentService(paymentRepositoryMock, restaurantAdapterMock)
	paymentHandler := NewPaymentHandler(paymentUseCase)
	newStatus := entities.PaymentStatusPaid.ToString()

	r := frameworks.Setup()
	r.POST("/payments/order/:id/status", paymentHandler.UpdateStatusByOrderId)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/payments/order/%d/status?status=%s", paymentPaid.OrderID, newStatus), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "CREATE response status is expected")
	assert.NotNil(t, response.Body, "Body is expected")

	var jsonResponse map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, paymentPaid.ID, jsonResponse["id"])
	assert.Equal(t, float64(paymentPaid.OrderID), jsonResponse["orderId"])
	assert.Equal(t, paymentPaid.Date.Format("2006-01-02"), jsonResponse["date"])
	assert.Equal(t, paymentPaid.Status.ToString(), jsonResponse["status"])
	assert.Equal(t, paymentPaid.Method.ToString(), jsonResponse["method"])

	_, exists := jsonResponse["errorReason"]
	assert.False(t, exists)
}

func Test_UpdatePaymentByOrderIdBadRequestId(t *testing.T) {
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	restaurantAdapterMock := new(domain.RestaurantClientMock)
	paymentUseCase := usecases.NewPaymentService(paymentRepositoryMock, restaurantAdapterMock)
	paymentHandler := NewPaymentHandler(paymentUseCase)
	newStatus := entities.PaymentStatusPaid.ToString()

	r := frameworks.Setup()
	r.POST("/payments/order/:id/status", paymentHandler.UpdateStatusByOrderId)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/payments/order/%s/status?status=%s", "x", newStatus), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 400, response.Code, "Bad Request response status is expected")
}

func Test_UpdatePaymentByOrderIdNotFound(t *testing.T) {
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	paymentRepositoryMock.On("GetPaymentByOrderId", mock.Anything).Return(nil, errors.New("order not found"))
	restaurantAdapterMock := new(domain.RestaurantClientMock)
	paymentUseCase := usecases.NewPaymentService(paymentRepositoryMock, restaurantAdapterMock)
	paymentHandler := NewPaymentHandler(paymentUseCase)
	newStatus := entities.PaymentStatusPaid.ToString()

	r := frameworks.Setup()
	r.POST("/payments/order/:id/status", paymentHandler.UpdateStatusByOrderId)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/payments/order/%d/status?status=%s", paymentPaid.OrderID, newStatus), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 404, response.Code, "Not Found response status is expected")
}

func Test_UpdatePaymentByOrderIdFailUpdate(t *testing.T) {
	paymentRepositoryMock := new(domain.PaymentRepositoryMock)
	paymentRepositoryMock.On("GetPaymentById", paymentPaid.ID).Return(paymentPaid, nil)
	paymentRepositoryMock.On("GetPaymentByOrderId", paymentPaid.OrderID).Return(paymentPaid, nil)
	paymentRepositoryMock.On("UpdateStatus", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error on update payment"))

	restaurantAdapterMock := new(domain.RestaurantClientMock)
	paymentUseCase := usecases.NewPaymentService(paymentRepositoryMock, restaurantAdapterMock)
	paymentHandler := NewPaymentHandler(paymentUseCase)
	newStatus := entities.PaymentStatusPaid.ToString()

	r := frameworks.Setup()
	r.POST("/payments/order/:id/status", paymentHandler.UpdateStatusByOrderId)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/payments/order/%d/status?status=%s", paymentPaid.OrderID, newStatus), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, 500, response.Code, "InternalServerError response status is expected")
}
