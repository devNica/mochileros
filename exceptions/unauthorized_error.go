package exceptions

type UnauthorizedError struct {
	Message string
}

func (unathorizedError UnauthorizedError) Error() string {
	return unathorizedError.Message
}
