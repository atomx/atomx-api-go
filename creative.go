package atomx

import (
	"strconv"
)

type Creative struct {
	Id         int64              `json:"id"`
	Size       *Size              `json:"size"`
	Attributes CreativeAttributes `json:"attributes"`
}

func (c Creative) path() string {
	return "creative/" + strconv.FormatInt(c.Id, 10)
}

type creativeResponse struct {
	Success  bool      `json:"success"`
	Error    string    `json:"error"`
	Creative *Creative `json:"creative"`
}

func (cr creativeResponse) err() error {
	if !cr.Success {
		return &ApiError{Message: cr.Error}
	}

	return nil
}

func (c *Creative) response() response {
	return &creativeResponse{
		Creative: c,
	}
}
