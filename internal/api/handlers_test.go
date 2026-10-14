package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "/healthz", nil)
	requestHandler := httptest.NewRecorder()
	HealthHandler(requestHandler, request)

	assert.Equal(t, http.StatusOK, requestHandler.Code)
	assert.Equal(t, "\"OK\"\n", requestHandler.Body.String())
}
