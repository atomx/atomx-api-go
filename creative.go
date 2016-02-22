package atomx

import (
	"strconv"
)

type Creative struct {
	ID         int64              `json:"id"`
	Advertiser *Advertiser        `json:"advertiser,omitempty"`
	Attributes CreativeAttributes `json:"attributes,omitempty"`
	AuditedAt  *Time              `json:"audited_at,omitempty"`
	Bans       CreativeBans       `json:"bans,omitempty"`
	Size       *Size              `json:"size,omitempty"`
	SizeBytes  *int64             `json:"size_bytes,omitempty"`
	Types      CreativeTypes      `json:"types,omitempty"`
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
