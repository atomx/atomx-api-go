package atomx

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (this Placement) path() string {
	if this.ID > 0 {
		return "placement/" + strconv.FormatInt(this.ID, 10)
	} else {
		return "placement"
	}
}

type placementResponse struct {
	Success   bool       "json:\"success\""
	Error     string     "json:\"error\""
	Placement *Placement "json:\"placement\""
}

func (this placementResponse) err() error {
	if !this.Success {
		return &ApiError{Message: this.Error}
	}

	return nil
}

func (this *Placement) response() response {
	return &placementResponse{
		Placement: this,
	}
}

type PlacementsList struct {
	List
	Placements []Placement "json:\"placements\""
}

func (this PlacementsList) path() string {
	return "placements?" + this.str()
}

type Placements []Placement

func (this Placements) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this Placements) Has(id int64) bool {
	for _, x := range this {
		if x.ID == id {
			return true
		}
	}

	return false
}

func (this *Placements) Add(y Placement) {
	*this = append(*this, y)
}

func (this *Placements) Remove(id int64) {
	for i, x := range *this {
		if x.ID == id {
			*this = append((*this)[:i], (*this)[i+1:]...)
			return
		}
	}
}

type PlacementRelation struct {
	Placement
}

func (this *PlacementRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}

func (this *PlacementRelation) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		return json.Unmarshal(data, &this.Placement)
	} else {
		return json.Unmarshal(data, &this.ID)
	}
}
