package atomx

type Size struct {
	ID     int64  `json:"id"`
	Name   string `json:"name,omitempty"`
	Width  int64  `json:"width,omitempty"`
	Height int64  `json:"height,omitempty"`
}
