package atomx

import (
	"strconv"
	"strings"
)

func (this Advertiser) path() string {
	if this.ID > 0 {
		return "advertiser/" + strconv.FormatInt(this.ID, 10)
	} else {
		return "advertiser"
	}
}

type advertiserResponse struct {
	Success    bool        "json:\"success\""
	Error      string      "json:\"error\""
	Advertiser *Advertiser "json:\"advertiser\""
}

func (this advertiserResponse) err() error {
	if !this.Success {
		return &ApiError{Message: this.Error}
	}

	return nil
}

func (this *Advertiser) response() response {
	return &advertiserResponse{
		Advertiser: this,
	}
}

type Advertisers struct {
	List
	Advertisers []Advertiser "json:\"advertisers\""
}

func (this Advertisers) path() string {
	return "advertisers?" + this.str()
}

func (this Advertisers) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this.Advertisers {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this Advertisers) Has(id int64) bool {
	for _, x := range this.Advertisers {
		if x.ID == id {
			return true
		}
	}

	return false
}

type AdvertiserRelation struct {
	Advertiser
}

func (this *AdvertiserRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}
