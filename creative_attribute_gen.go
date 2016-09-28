package atomx

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (this CreativeAttribute) path() string {
	if this.ID > 0 {
		return "creative_attribute/" + strconv.FormatInt(this.ID, 10)
	} else {
		return "creative_attribute"
	}
}

type creative_attributeResponse struct {
	Success           bool               "json:\"success\""
	Error             string             "json:\"error\""
	CreativeAttribute *CreativeAttribute "json:\"creative_attribute\""
}

func (this creative_attributeResponse) err() error {
	if !this.Success {
		return &ApiError{Message: this.Error}
	}

	return nil
}

func (this *CreativeAttribute) response() response {
	return &creative_attributeResponse{
		CreativeAttribute: this,
	}
}

type CreativeAttributesList struct {
	List
	CreativeAttributes []CreativeAttribute "json:\"creative_attributes\""
}

func (this CreativeAttributesList) path() string {
	return "creative_attributes?" + this.str()
}

type CreativeAttributes []CreativeAttribute

func (this CreativeAttributes) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this CreativeAttributes) Has(id int64) bool {
	for _, x := range this {
		if x.ID == id {
			return true
		}
	}

	return false
}

func (this *CreativeAttributes) Add(y CreativeAttribute) {
	*this = append(*this, y)
}

func (this *CreativeAttributes) Remove(id int64) {
	for i, x := range *this {
		if x.ID == id {
			*this = append((*this)[:i], (*this)[i+1:]...)
			return
		}
	}
}

type CreativeAttributeRelation struct {
	CreativeAttribute
}

func (this *CreativeAttributeRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}

func (this *CreativeAttributeRelation) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		return json.Unmarshal(data, &this.CreativeAttribute)
	} else {
		return json.Unmarshal(data, &this.ID)
	}
}
