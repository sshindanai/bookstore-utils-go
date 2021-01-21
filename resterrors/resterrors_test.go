package resterrors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	testMsg := "test_error_500"

	err := NewInternalServerError(testMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, http.StatusInternalServerError)
	assert.EqualValues(t, err.Message, testMsg)
	assert.EqualValues(t, err.Error, "internal_server_error")
}

func TestNewBadrequestError(t *testing.T) {
	testMsg := "test_error_400"

	err := NewBadRequestError(testMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, http.StatusBadRequest)
	assert.EqualValues(t, err.Message, testMsg)
	assert.EqualValues(t, err.Error, "bad_request")
}

func TestNewUnauthorizedError(t *testing.T) {
	testMsg := "test_error_401"

	err := NewUnauthorizedError(testMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, http.StatusUnauthorized)
	assert.EqualValues(t, err.Message, testMsg)
	assert.EqualValues(t, err.Error, "unauthorized")
}

func TestNewNotFoundError(t *testing.T) {
	testMsg := "test_error_404"

	err := NewNotFoundError(testMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, http.StatusNotFound)
	assert.EqualValues(t, err.Message, testMsg)
	assert.EqualValues(t, err.Error, "not_found")
}

func TestNewConflictError(t *testing.T) {
	testMsg := "test_error_409"

	err := NewConflictError(testMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, http.StatusConflict)
	assert.EqualValues(t, err.Message, testMsg)
	assert.EqualValues(t, err.Error, "conflict")
}

func TestNewRestErrorFromBytes(t *testing.T) {
	testMsg := []byte("test_error_500")

	err := NewRestErrorFromBytes(testMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, http.StatusBadRequest)
	assert.EqualValues(t, err.Message, string(testMsg))
	assert.EqualValues(t, err.Error, "bad_request")
}
