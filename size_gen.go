package atomx

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (this Size) path() string {
	if this.ID > 0 {
		return "size/" + strconv.FormatInt(this.ID, 10)
	} else {
		return "size"
	}
}

type sizeResponse struct {
	Success bool   "json:\"success\""
	Error   string "json:\"error\""
	Size    *Size  "json:\"size\""
}

func (this sizeResponse) err() error {
	if !this.Success {
		return &ApiError{Message: this.Error}
	}

	return nil
}

func (this *Size) response() response {
	return &sizeResponse{
		Size: this,
	}
}

type SizesList struct {
	List
	Sizes []Size "json:\"sizes\""
}

func (this SizesList) path() string {
	return "sizes?" + this.str()
}

type Sizes []Size

func (this Sizes) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this Sizes) Has(id int64) bool {
	for _, x := range this {
		if x.ID == id {
			return true
		}
	}

	return false
}

func (this *Sizes) Add(y Size) {
	*this = append(*this, y)
}

func (this *Sizes) Remove(id int64) {
	for i, x := range *this {
		if x.ID == id {
			*this = append((*this)[:i], (*this)[i+1:]...)
			return
		}
	}
}

type SizeRelation struct {
	Size
}

func (this *SizeRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}

func (this *SizeRelation) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		return json.Unmarshal(data, &this.Size)
	} else {
		return json.Unmarshal(data, &this.ID)
	}
}
