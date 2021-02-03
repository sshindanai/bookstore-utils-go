package resterrors

import (
	"encoding/json"
	"net/http"
)

type RestErr interface {
	Message() string
	StatusCode() int
	Error() string
	Causes() []interface{}
}

// RestErr is a custom error handling struct for this project
type restErr struct {
	message    string        `json:"message"`
	statusCode int           `json:"status_statusCode"`
	err        string        `json:"error"`
	causes     []interface{} `json:"causes"`
}

func (e restErr) Message() string {
	return e.message
}

func (e restErr) StatusCode() int {
	return e.statusCode
}

func (e restErr) Error() string {
	return e.err
}

func (e restErr) Causes() []interface{} {
	return e.causes
}

func NewRestError(message string, statusCode int, err string, causes []interface{}) RestErr {
	return restErr{
		message:    message,
		statusCode: statusCode,
		err:        err,
		causes:     causes,
	}
}

// NewBadRequestError is a custom error handling for BadRequest error
func NewBadRequestError(message string) RestErr {
	return restErr{
		message:    message,
		statusCode: http.StatusBadRequest,
		err:        "bad_request",
	}
}

// NewNotFoundError is a custom error handling for NotFound error
func NewNotFoundError(message string) RestErr {
	return restErr{
		message:    message,
		statusCode: http.StatusNotFound,
		err:        "not_found",
	}
}

// NewUnauthorizedError is a custom error handling for Unauthorized error
func NewUnauthorizedError(message string) RestErr {
	return restErr{
		message:    message,
		statusCode: http.StatusUnauthorized,
		err:        "unauthorized",
	}
}

// NewInternalServerError is a custom error handling for InternalServerError error
func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		message:    message,
		statusCode: http.StatusInternalServerError,
		err:        "internal_server_error",
	}
	if err != nil {
		result.causes = append(result.causes, err.Error())
	}

	return result
}

// NewConflictError is a custom error handling for Conflict error
func NewConflictError(message string) RestErr {
	return restErr{
		message:    message,
		statusCode: http.StatusConflict,
		err:        "conflict",
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
