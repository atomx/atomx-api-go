package atomx

type Resources interface {
	path() string
	err() error
}
