package errors

import (
	"fmt"
	"net/http"
)

type ErrorCode string

const (
	ErrorCodeNotFound      ErrorCode = "NOT_FOUND"
	ErrorCodeInvalidInput  ErrorCode = "INVALID_INPUT"
	ErrorCodeAlreadyExists ErrorCode = "ALREADY_EXISTS"
	ErrorCodeUnauthorized  ErrorCode = "UNAUTHORIZED"
	ErrorCodeForbidden     ErrorCode = "FORBIDDEN"
	ErrorCodeInternal      ErrorCode = "INTERNAL_ERROR"
	ErrorCodeValidation    ErrorCode = "VALIDATION_ERROR"
	ErrorCodeBusinessRule  ErrorCode = "BUSINESS_RULE"
	ErrorCodeDatabase      ErrorCode = "DATABASE_ERROR"
)

type AppError struct {
	Code     ErrorCode   `json:"code"`
	Message  string      `json:"message"`
	Details  interface{} `json:"details,omitempty"`
	HTTPCode int         `json:"-"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Construtors de erro
func NewNotFoundError(resource string, id string) *AppError {
	return &AppError{
		Code:     ErrorCodeNotFound,
		Message:  fmt.Sprintf("%s com ID %s não encontrado", resource, id),
		HTTPCode: http.StatusNotFound,
	}
}

func NewValidationError(message string, details interface{}) *AppError {
	return &AppError{
		Code:     ErrorCodeValidation,
		Message:  message,
		Details:  details,
		HTTPCode: http.StatusBadRequest,
	}
}

func NewBusinessRuleError(message string) *AppError {
	return &AppError{
		Code:     ErrorCodeBusinessRule,
		Message:  message,
		HTTPCode: http.StatusUnprocessableEntity,
	}
}

func NewAlreadyExistsError(resource, field, value string) *AppError {
	return &AppError{
		Code:     ErrorCodeAlreadyExists,
		Message:  fmt.Sprintf("%s com %s '%s' já existe", resource, field, value),
		HTTPCode: http.StatusConflict,
	}
}

func NewInternalError(message string, err error) *AppError {
	details := ""
	if err != nil {
		details = err.Error()
	}
	return &AppError{
		Code:     ErrorCodeInternal,
		Message:  message,
		Details:  details,
		HTTPCode: http.StatusInternalServerError,
	}
}

func NewDatabaseError(message string, err error) *AppError {
	details := ""
	if err != nil {
		details = err.Error()
	}
	return &AppError{
		Code:     ErrorCodeDatabase,
		Message:  message,
		Details:  details,
		HTTPCode: http.StatusInternalServerError,
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:     ErrorCodeUnauthorized,
		Message:  message,
		HTTPCode: http.StatusUnauthorized,
	}
}
