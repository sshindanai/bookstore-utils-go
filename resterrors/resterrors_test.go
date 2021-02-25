package resterrors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	testMsg := "test_error_500"

	err := NewInternalServerError(testMsg, errors.New("internal error"))

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode())
	assert.EqualValues(t, testMsg, err.Message())
	assert.EqualValues(t, "internal_server_error", err.Error())

	assert.NotNil(t, err.Causes())
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "internal error", err.Causes()[0])
}

func TestNewBadrequestError(t *testing.T) {
	testMsg := "test_error_400"

	err := NewBadRequestError(testMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.StatusCode())
	assert.EqualValues(t, testMsg, err.Message())
	assert.EqualValues(t, "bad_request", err.Error())
}

func TestNewUnauthorizedError(t *testing.T) {
	testMsg := "test_error_401"

	err := NewUnauthorizedError(testMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode())
	assert.EqualValues(t, testMsg, err.Message())
	assert.EqualValues(t, "unauthorized", err.Error())
}

func TestNewNotFoundError(t *testing.T) {
	testMsg := "test_error_404"

	err := NewNotFoundError(testMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode())
	assert.EqualValues(t, testMsg, err.Message())
	assert.EqualValues(t, "not_found", err.Error())
}

func TestNewConflictError(t *testing.T) {
	testMsg := "test_error_409"

	err := NewConflictError(testMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusConflict, err.StatusCode())
	assert.EqualValues(t, testMsg, err.Message())
	assert.EqualValues(t, "conflict", err.Error())
}

func TestNewRestErrorFromBytes(t *testing.T) {
	testMsg := []byte("test_error_500")

	_, err := NewRestErrorFromBytes(testMsg)

	assert.NotNil(t, err)
	//assert.EqualValues(t, http.StatusBadRequest, restErr.StatusCode())
	//assert.EqualValues(t, string(testMsg), restErr.Message())
}

func TestNewRestError(t *testing.T) {
	message := "Test_TestNewRestError"
	statusCode := http.StatusBadRequest
	err := "bad_request_test"

	newErr := NewRestError(message, statusCode, err, nil)

	assert.NotNil(t, newErr)
	assert.EqualValues(t, statusCode, newErr.StatusCode())
	assert.EqualValues(t, message, newErr.Message())
	assert.EqualValues(t, err, newErr.Error())
}
