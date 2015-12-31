package atomx

import (
	"strconv"
)

type Size struct {
	ID     int64  `json:"id"`
	Name   string `json:"name,omitempty"`
	Width  int64  `json:"width,omitempty"`
	Height int64  `json:"height,omitempty"`
}

func (s Size) path() string {
	return "size/" + strconv.FormatInt(s.ID, 10)
}

type SizeResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Size    *Size  `json:"size"`
}

func (sr SizeResponse) err() error {
	if !sr.Success {
		return &ApiError{Message: sr.Error}
	}

	return nil
}

func (s *Size) response() response {
	return &SizeResponse{
		Size: s,
	}
}
