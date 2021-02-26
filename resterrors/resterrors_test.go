package resterrors

import (
	"encoding/json"
	"errors"
	"fmt"
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
	assert.EqualValues(t, "message: test_error_500 - status: 500 - error: internal_server_error - causes: [internal error]", err.Error())

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
	assert.EqualValues(t, "message: test_error_400 - status: 400 - error: bad_request - causes: []", err.Error())
}

func TestNewUnauthorizedError(t *testing.T) {
	testMsg := "test_error_401"

	err := NewUnauthorizedError(testMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode())
	assert.EqualValues(t, testMsg, err.Message())
	assert.EqualValues(t, "message: test_error_401 - status: 401 - error: unauthorized - causes: []", err.Error())
}

func TestNewNotFoundError(t *testing.T) {
	testMsg := "test_error_404"

	err := NewNotFoundError(testMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode())
	assert.EqualValues(t, testMsg, err.Message())
	assert.EqualValues(t, "message: test_error_404 - status: 404 - error: not_found - causes: []", err.Error())
}

func TestNewConflictError(t *testing.T) {
	testMsg := "test_error_409"

	err := NewConflictError(testMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusConflict, err.StatusCode())
	assert.EqualValues(t, testMsg, err.Message())
	assert.EqualValues(t, "message: test_error_409 - status: 409 - error: conflict - causes: []", err.Error())
}

func TestNewRestError(t *testing.T) {
	message := "Test_TestNewRestError"
	statusCode := http.StatusBadRequest
	err := "bad_request_test"

	newErr := NewRestError(message, statusCode, err, nil)

	assert.NotNil(t, newErr)
	assert.EqualValues(t, statusCode, newErr.StatusCode())
	assert.EqualValues(t, message, newErr.Message())
	assert.EqualValues(t, fmt.Sprintf("message: %s - status: %d - error: %s - causes: []",
		message,
		statusCode,
		err),
		newErr.Error())
}

func TestNewRestFromBytes(t *testing.T) {
	b := []byte("test_error_500")

	_, err := NewRestErrorFromBytes(b)

	assert.NotNil(t, err)
}

func TestNewRestFromBytesReturnRestErr(t *testing.T) {
	testErr := restErr{
		ErrMessage:    "test_error_400",
		ErrStatusCode: http.StatusBadRequest,
		ErrError:      "bad_request",
	}

	json, _ := json.Marshal(testErr)
	restErr, _ := NewRestErrorFromBytes(json)

	assert.NotNil(t, restErr)
	assert.EqualValues(t, testErr.ErrMessage, restErr.Message())
	assert.EqualValues(t, testErr.ErrStatusCode, restErr.StatusCode())
	assert.EqualValues(t, fmt.Sprintf("message: %s - status: %d - error: %s - causes: []",
		testErr.ErrMessage,
		testErr.ErrStatusCode,
		testErr.ErrError), restErr.Error())
}
