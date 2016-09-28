package atomx

import (
	"encoding/json"
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

type AdvertisersList struct {
	List
	Advertisers []Advertiser "json:\"advertisers\""
}

func (this AdvertisersList) path() string {
	return "advertisers?" + this.str()
}

type Advertisers []Advertiser

func (this Advertisers) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this Advertisers) Has(id int64) bool {
	for _, x := range this {
		if x.ID == id {
			return true
		}
	}

	return false
}

func (this *Advertisers) Add(y Advertiser) {
	*this = append(*this, y)
}

func (this *Advertisers) Remove(id int64) {
	for i, x := range *this {
		if x.ID == id {
			*this = append((*this)[:i], (*this)[i+1:]...)
			return
		}
	}
}

type AdvertiserRelation struct {
	Advertiser
}

func (this *AdvertiserRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}

func (this *AdvertiserRelation) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		return json.Unmarshal(data, &this.Advertiser)
	} else {
		return json.Unmarshal(data, &this.ID)
	}
}
