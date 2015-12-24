package atomx

import (
	"strconv"
)

type Advertiser struct {
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	Network *Network `json:"network"`
}

func (a Advertiser) path() string {
	return "advertiser/" + strconv.FormatInt(a.ID, 10)
}

type advertiserResponse struct {
	Success    bool        `json:"success"`
	Error      string      `json:"error"`
	Advertiser *Advertiser `json:"advertiser"`
}

func (ar advertiserResponse) err() error {
	if !ar.Success {
		return &ApiError{Message: ar.Error}
	}

	return nil
}

func (a *Advertiser) response() response {
	return &advertiserResponse{
		Advertiser: a,
	}
}
