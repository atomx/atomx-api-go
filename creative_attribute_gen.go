package atomx

import (
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

type CreativeAttributes struct {
	List
	CreativeAttributes []CreativeAttribute "json:\"creative_attributes\""
}

func (this CreativeAttributes) path() string {
	return "creative_attributes?" + this.str()
}

func (this CreativeAttributes) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this.CreativeAttributes {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this CreativeAttributes) Has(id int64) bool {
	for _, x := range this.CreativeAttributes {
		if x.ID == id {
			return true
		}
	}

	return false
}

type CreativeAttributeRelation struct {
	CreativeAttribute
}

func (this *CreativeAttributeRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}
