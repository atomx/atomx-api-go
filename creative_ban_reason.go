package atomx

import (
	"strconv"
)

type CreativeBanReason struct {
	ID          int64  `json:"id"`
	Description string `json:"description,omitempty"`
}

func (cbr CreativeBanReason) path() string {
	return "creative_ban_reason/" + strconv.FormatInt(cbr.ID, 10)
}

type creativeBanReasonResponse struct {
	Success           bool               `json:"success"`
	Error             string             `json:"error"`
	CreativeBanReason *CreativeBanReason `json:"creative_ban_reason"`
}

func (cbrr creativeBanReasonResponse) err() error {
	if !cbrr.Success {
		return &ApiError{Message: cbrr.Error}
	}

	return nil
}

func (cbr *CreativeBanReason) response() response {
	return &creativeBanReasonResponse{
		CreativeBanReason: cbr,
	}
}
