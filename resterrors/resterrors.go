package resterrors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// RestErr is a custom error handling struct for this project
type RestErr interface {
	Message() string
	StatusCode() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	ErrMessage    string        `json:"message"`
	ErrStatusCode int           `json:"status_code"`
	ErrError      string        `json:"error"`
	ErrCauses     []interface{} `json:"causes"`
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v", e.ErrMessage, e.ErrStatusCode, e.ErrError, e.ErrCauses)
}

func (e restErr) Message() string {
	return e.ErrMessage
}

func (e restErr) StatusCode() int {
	return e.ErrStatusCode
}

func (e restErr) Causes() []interface{} {
	return e.ErrCauses
}

// NewRestError is a custom error handling for any error
func NewRestError(message string, statusCode int, err string, causes []interface{}) RestErr {
	return restErr{
		ErrMessage:    message,
		ErrStatusCode: statusCode,
		ErrError:      err,
		ErrCauses:     causes,
	}
}

// NewBadRequestError is a custom error handling for BadRequest error
func NewBadRequestError(message string) RestErr {
	return restErr{
		ErrMessage:    message,
		ErrStatusCode: http.StatusBadRequest,
		ErrError:      "bad_request",
	}
}

// NewNotFoundError is a custom error handling for NotFound error
func NewNotFoundError(message string) RestErr {
	return restErr{
		ErrMessage:    message,
		ErrStatusCode: http.StatusNotFound,
		ErrError:      "not_found",
	}
}

// NewUnauthorizedError is a custom error handling for Unauthorized error
func NewUnauthorizedError(message string) RestErr {
	return restErr{
		ErrMessage:    message,
		ErrStatusCode: http.StatusUnauthorized,
		ErrError:      "unauthorized",
	}
}

// NewInternalServerError is a custom error handling for InternalServerError error
func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		ErrMessage:    message,
		ErrStatusCode: http.StatusInternalServerError,
		ErrError:      "internal_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}

	return result
}

// NewConflictError is a custom error handling for Conflict error
func NewConflictError(message string) RestErr {
	return restErr{
		ErrMessage:    message,
		ErrStatusCode: http.StatusConflict,
		ErrError:      "conflict",
	}
}

// NewRestErrorFromBytes is a custom error handling for byte error
func NewRestErrorFromBytes(b []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(b, &apiErr); err != nil {
		return nil, err
	}
	return apiErr, nil
}
