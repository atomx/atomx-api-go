package atomx

import (
	"strconv"
)

type DomainAttribute struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (da DomainAttribute) path() string {
	return "domain_attribute/" + strconv.FormatInt(da.ID, 10)
}

type domainAttributeResponse struct {
	Success         bool             `json:"success"`
	Error           string           `json:"error"`
	DomainAttribute *DomainAttribute `json:"domain_attribute"`
}

func (dar domainAttributeResponse) err() error {
	if !dar.Success {
		return &ApiError{Message: dar.Error}
	}

	return nil
}

func (da *DomainAttribute) response() response {
	return &domainAttributeResponse{
		DomainAttribute: da,
	}
}
