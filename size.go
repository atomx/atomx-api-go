package atomx

import (
	"strconv"
)

type Size struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

func (s Size) path() string {
	return "size/" + strconv.FormatInt(s.Id, 10)
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
