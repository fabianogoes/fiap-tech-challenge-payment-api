package rest

import (
	"github.com/fabianogoes/fiap-payment/frameworks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Welcome(t *testing.T) {
	r := frameworks.Setup()
	r.GET("/", Welcome)
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func Test_Health(t *testing.T) {
	r := frameworks.Setup()
	r.GET("/health", Health)
	request, _ := http.NewRequest("GET", "/health", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func Test_Environment(t *testing.T) {
	r := frameworks.Setup()
	r.GET("/env", Environment)
	request, _ := http.NewRequest("GET", "/env", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
