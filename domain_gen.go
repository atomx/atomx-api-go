package atomx

import (
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

type Domains struct {
	List
	Domains []Domain "json:\"domains\""
}

func (this Domains) path() string {
	return "domains?" + this.str()
}

func (this Domains) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this.Domains {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this Domains) Has(id int64) bool {
	for _, x := range this.Domains {
		if x.ID == id {
			return true
		}
	}

	return false
}

type DomainRelation struct {
	Domain
}

func (this *DomainRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}
