package atomx

import (
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

type Sizes struct {
	List
	Sizes []Size "json:\"sizes\""
}

func (this Sizes) path() string {
	return "sizes?" + this.str()
}

func (this Sizes) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this.Sizes {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this Sizes) Has(id int64) bool {
	for _, x := range this.Sizes {
		if x.ID == id {
			return true
		}
	}

	return false
}

type SizeRelation struct {
	Size
}

func (this *SizeRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}
