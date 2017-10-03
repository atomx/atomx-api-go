package atomx

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (this Ssp) path() string {
	if this.ID > 0 {
		return "ssp/" + strconv.FormatInt(this.ID, 10)
	} else {
		return "ssp"
	}
}

type sspResponse struct {
	Success bool   "json:\"success\""
	Error   string "json:\"error\""
	Ssp     *Ssp   "json:\"ssp\""
}

func (this sspResponse) err() error {
	if !this.Success {
		return &ApiError{Message: this.Error}
	}

	return nil
}

func (this *Ssp) response() response {
	return &sspResponse{
		Ssp: this,
	}
}

type SspsList struct {
	List
	Ssps []Ssp "json:\"ssps\""
}

func (this SspsList) path() string {
	return "ssps?" + this.str()
}

type Ssps []Ssp

func (this Ssps) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this Ssps) Has(id int64) bool {
	for _, x := range this {
		if x.ID == id {
			return true
		}
	}

	return false
}

func (this *Ssps) Add(y Ssp) {
	*this = append(*this, y)
}

func (this *Ssps) Remove(id int64) {
	for i, x := range *this {
		if x.ID == id {
			*this = append((*this)[:i], (*this)[i+1:]...)
			return
		}
	}
}

type SspRelation struct {
	Ssp
}

func (this *SspRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}

func (this *SspRelation) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		return json.Unmarshal(data, &this.Ssp)
	} else {
		return json.Unmarshal(data, &this.ID)
	}
}
