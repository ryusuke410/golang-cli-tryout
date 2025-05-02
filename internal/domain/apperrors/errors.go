package apperrors

import "fmt"

func NewAppError(code Code) *AppError {
	return &AppError{
		Code: code,
	}
}

func (e *AppError) WithMessage(message string) *AppError {
	e.Message = message
	return e
}

func (e *AppError) WithMessagef(format string, args ...interface{}) *AppError {
	e.Message = fmt.Sprintf(format, args...)
	return e
}

func (e *AppError) WithCause(cause error) *AppError {
	e.Cause = cause
	return e
}
