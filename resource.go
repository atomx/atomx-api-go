package atomx

type Resource interface {
	path() string
	response() response
}
