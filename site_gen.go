package atomx

import (
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

type Sites struct {
	List
	Sites []Site "json:\"sites\""
}

func (this Sites) path() string {
	return "sites?" + this.str()
}

func (this Sites) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this.Sites {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this Sites) Has(id int64) bool {
	for _, x := range this.Sites {
		if x.ID == id {
			return true
		}
	}

	return false
}

type SiteRelation struct {
	Site
}

func (this *SiteRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}
