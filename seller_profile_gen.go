package atomx

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (this SellerProfile) path() string {
	if this.ID > 0 {
		return "seller_profile/" + strconv.FormatInt(this.ID, 10)
	} else {
		return "seller_profile"
	}
}

type seller_profileResponse struct {
	Success       bool           "json:\"success\""
	Error         string         "json:\"error\""
	SellerProfile *SellerProfile "json:\"seller_profile\""
}

func (this seller_profileResponse) err() error {
	if !this.Success {
		return &ApiError{Message: this.Error}
	}

	return nil
}

func (this *SellerProfile) response() response {
	return &seller_profileResponse{
		SellerProfile: this,
	}
}

type SellerProfilesList struct {
	List
	SellerProfiles []SellerProfile "json:\"seller_profiles\""
}

func (this SellerProfilesList) path() string {
	return "seller_profiles?" + this.str()
}

type SellerProfiles []SellerProfile

func (this SellerProfiles) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this SellerProfiles) Has(id int64) bool {
	for _, x := range this {
		if x.ID == id {
			return true
		}
	}

	return false
}

func (this *SellerProfiles) Add(y SellerProfile) {
	*this = append(*this, y)
}

func (this *SellerProfiles) Remove(id int64) {
	for i, x := range *this {
		if x.ID == id {
			*this = append((*this)[:i], (*this)[i+1:]...)
			return
		}
	}
}

type SellerProfileRelation struct {
	SellerProfile
}

func (this *SellerProfileRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}

func (this *SellerProfileRelation) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		return json.Unmarshal(data, &this.SellerProfile)
	} else {
		return json.Unmarshal(data, &this.ID)
	}
}
