package apperrors

import "fmt"

type AppError struct {
	Code    Code
	Message string
	Cause   error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code.String(), e.Message)
}

func (e *AppError) Is(target error) bool {
	t, ok := target.(*AppError)
	return ok && e.Code == t.Code
}

func (e *AppError) Unwrap() error {
	return e.Cause
}
