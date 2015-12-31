package atomx

import (
	"strconv"
)

type Creative struct {
	ID         int64              `json:"id"`
	Size       *Size              `json:"size,omitempty"`
	Types      CreativeTypes      `json:"types,omitempty"`
	Attributes CreativeAttributes `json:"attributes,omitempty"`
	Advertiser *Advertiser        `json:"advertiser,omitempty"`
	SizeBytes  int64              `json:"size_bytes,omitempty"`
}

func (c Creative) path() string {
	return "creative/" + strconv.FormatInt(c.ID, 10)
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
