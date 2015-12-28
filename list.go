package atomx

import (
	"strconv"
)

type List struct {
	// These attributes are used as input and are set in the output.
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`

	// These attributes are only set in the output.
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Count   int64  `json:"count"`
}

func (l List) offsetLimit() string {
	r := "offset=" + strconv.FormatInt(l.Offset, 10)

	if l.Limit > 0 {
		r += "&limit=" + strconv.FormatInt(l.Limit, 10)
	}

	return r
}

func (l List) err() error {
	if !l.Success {
		return &ApiError{Message: l.Error}
	}

	return nil
}
