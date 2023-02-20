package exceptions

type BadReqquestError struct {
	Message string
}

func (badRequestError BadReqquestError) Error() string {
	return badRequestError.Message
}
