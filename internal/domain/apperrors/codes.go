package apperrors

type Code int

const (
	Ok              Code = 0
	Unknown         Code = 1
	InvalidArgument Code = 2
	Internal        Code = 3
)

func (c Code) String() string {
	switch c {
	case Ok:
		return "OK"
	case Unknown:
		return "UNKNOWN"
	case InvalidArgument:
		return "INVALID_ARGUMENT"
	case Internal:
		return "INTERNAL"
	default:
		return "UNKNOWN"
	}
}

func (c Code) NewAppError() *AppError {
	return NewAppError(c)
}

func (c Code) Error() string {
	return c.String()
}
