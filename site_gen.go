package atomx

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (this Site) path() string {
	if this.ID > 0 {
		return "site/" + strconv.FormatInt(this.ID, 10)
	} else {
		return "site"
	}
}

type siteResponse struct {
	Success bool   "json:\"success\""
	Error   string "json:\"error\""
	Site    *Site  "json:\"site\""
}

func (this siteResponse) err() error {
	if !this.Success {
		return &ApiError{Message: this.Error}
	}

	return nil
}

func (this *Site) response() response {
	return &siteResponse{
		Site: this,
	}
}

type SitesList struct {
	List
	Sites []Site "json:\"sites\""
}

func (this SitesList) path() string {
	return "sites?" + this.str()
}

type Sites []Site

func (this Sites) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this Sites) Has(id int64) bool {
	for _, x := range this {
		if x.ID == id {
			return true
		}
	}

	return false
}

func (this *Sites) Add(y Site) {
	*this = append(*this, y)
}

func (this *Sites) Remove(id int64) {
	for i, x := range *this {
		if x.ID == id {
			*this = append((*this)[:i], (*this)[i+1:]...)
			return
		}
	}
}

type SiteRelation struct {
	Site
}

func (this *SiteRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}

func (this *SiteRelation) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		return json.Unmarshal(data, &this.Site)
	} else {
		return json.Unmarshal(data, &this.ID)
	}
}
