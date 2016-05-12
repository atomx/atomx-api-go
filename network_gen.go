package atomx

import (
	"strconv"
	"strings"
)

func (this Network) path() string {
	if this.ID > 0 {
		return "network/" + strconv.FormatInt(this.ID, 10)
	} else {
		return "network"
	}
}

type networkResponse struct {
	Success bool     "json:\"success\""
	Error   string   "json:\"error\""
	Network *Network "json:\"network\""
}

func (this networkResponse) err() error {
	if !this.Success {
		return &ApiError{Message: this.Error}
	}

	return nil
}

func (this *Network) response() response {
	return &networkResponse{
		Network: this,
	}
}

type NetworksList struct {
	List
	Networks []Network "json:\"networks\""
}

func (this NetworksList) path() string {
	return "networks?" + this.str()
}

type Networks []Network

func (this Networks) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this Networks) Has(id int64) bool {
	for _, x := range this {
		if x.ID == id {
			return true
		}
	}

	return false
}

func (this *Networks) Add(y Network) {
	*this = append(*this, y)
}

func (this *Networks) Remove(id int64) {
	for i, x := range *this {
		if x.ID == id {
			*this = append((*this)[:i], (*this)[i+1:]...)
			return
		}
	}
}

type NetworkRelation struct {
	Network
}

func (this *NetworkRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}
