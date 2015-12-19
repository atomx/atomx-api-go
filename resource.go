package atomx

type resource interface {
	path() string
	response() response
}
