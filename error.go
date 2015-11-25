package atomx

type ApiError struct {
	Message string
}

func (e *ApiError) Error() string {
	return e.Message
}
