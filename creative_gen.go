package atomx

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (this Creative) path() string {
	if this.ID > 0 {
		return "creative/" + strconv.FormatInt(this.ID, 10)
	} else {
		return "creative"
	}
}

type creativeResponse struct {
	Success  bool      "json:\"success\""
	Error    string    "json:\"error\""
	Creative *Creative "json:\"creative\""
}

func (this creativeResponse) err() error {
	if !this.Success {
		return &ApiError{Message: this.Error}
	}

	return nil
}

func (this *Creative) response() response {
	return &creativeResponse{
		Creative: this,
	}
}

type CreativesList struct {
	List
	Creatives []Creative "json:\"creatives\""
}

func (this CreativesList) path() string {
	return "creatives?" + this.str()
}

type Creatives []Creative

func (this Creatives) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this Creatives) Has(id int64) bool {
	for _, x := range this {
		if x.ID == id {
			return true
		}
	}

	return false
}

func (this *Creatives) Add(y Creative) {
	*this = append(*this, y)
}

func (this *Creatives) Remove(id int64) {
	for i, x := range *this {
		if x.ID == id {
			*this = append((*this)[:i], (*this)[i+1:]...)
			return
		}
	}
}

type CreativeRelation struct {
	Creative
}

func (this *CreativeRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}

func (this *CreativeRelation) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		return json.Unmarshal(data, &this.Creative)
	} else {
		return json.Unmarshal(data, &this.ID)
	}
}
