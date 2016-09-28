package atomx

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (this Domain) path() string {
	if this.ID > 0 {
		return "domain/" + strconv.FormatInt(this.ID, 10)
	} else {
		return "domain"
	}
}

type domainResponse struct {
	Success bool    "json:\"success\""
	Error   string  "json:\"error\""
	Domain  *Domain "json:\"domain\""
}

func (this domainResponse) err() error {
	if !this.Success {
		return &ApiError{Message: this.Error}
	}

	return nil
}

func (this *Domain) response() response {
	return &domainResponse{
		Domain: this,
	}
}

type DomainsList struct {
	List
	Domains []Domain "json:\"domains\""
}

func (this DomainsList) path() string {
	return "domains?" + this.str()
}

type Domains []Domain

func (this Domains) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this Domains) Has(id int64) bool {
	for _, x := range this {
		if x.ID == id {
			return true
		}
	}

	return false
}

func (this *Domains) Add(y Domain) {
	*this = append(*this, y)
}

func (this *Domains) Remove(id int64) {
	for i, x := range *this {
		if x.ID == id {
			*this = append((*this)[:i], (*this)[i+1:]...)
			return
		}
	}
}

type DomainRelation struct {
	Domain
}

func (this *DomainRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}

func (this *DomainRelation) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		return json.Unmarshal(data, &this.Domain)
	} else {
		return json.Unmarshal(data, &this.ID)
	}
}
