package atomx

import (
	"strconv"
	"strings"
)

func (this DomainAttribute) path() string {
	if this.ID > 0 {
		return "domain_attribute/" + strconv.FormatInt(this.ID, 10)
	} else {
		return "domain_attribute"
	}
}

type domain_attributeResponse struct {
	Success         bool             "json:\"success\""
	Error           string           "json:\"error\""
	DomainAttribute *DomainAttribute "json:\"domain_attribute\""
}

func (this domain_attributeResponse) err() error {
	if !this.Success {
		return &ApiError{Message: this.Error}
	}

	return nil
}

func (this *DomainAttribute) response() response {
	return &domain_attributeResponse{
		DomainAttribute: this,
	}
}

type DomainAttributes struct {
	List
	DomainAttributes []DomainAttribute "json:\"domain_attributes\""
}

func (this DomainAttributes) path() string {
	return "domain_attributes?" + this.str()
}

func (this DomainAttributes) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this.DomainAttributes {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this DomainAttributes) Has(id int64) bool {
	for _, x := range this.DomainAttributes {
		if x.ID == id {
			return true
		}
	}

	return false
}

type DomainAttributeRelation struct {
	DomainAttribute
}

func (this *DomainAttributeRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}
