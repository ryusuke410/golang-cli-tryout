package apperrors

func Wrap(err error) *AppError {
	if err == nil {
		return nil
	}
	if appErr, ok := err.(*AppError); ok {
		return appErr
	}
	if code, ok := err.(Code); ok {
		return code.NewAppError().WithCause(err)
	}
	return NewAppError(Internal).WithCause(err)
}
