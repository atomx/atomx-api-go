package atomx

import (
	"strconv"
)

type Site struct {
	ID       int64 `json:"id"`
	DomainID int64 `json:"domain,omitempty"`
}

func (s *Site) path() string {
	return "site/" + strconv.FormatInt(s.ID, 10)
}

type siteResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Site    *Site  `json:"site"`
}

func (sr *siteResponse) err() error {
	if !sr.Success {
		return &ApiError{Message: sr.Error}
	}

	return nil
}

func (s *Site) response() response {
	return &siteResponse{
		Site: s,
	}
}
