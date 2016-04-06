package atomx

import (
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

type Placements struct {
	List
	Placements []Placement "json:\"placements\""
}

func (this Placements) path() string {
	return "placements?" + this.str()
}

func (this Placements) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this.Placements {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this Placements) Has(id int64) bool {
	for _, x := range this.Placements {
		if x.ID == id {
			return true
		}
	}

	return false
}

type PlacementRelation struct {
	Placement
}

func (this *PlacementRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}
