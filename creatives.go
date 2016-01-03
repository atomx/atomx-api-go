package atomx

type Creatives struct {
	List
	Creatives []Creative `json:"creatives"`
}

func (cs Creatives) path() string {
	return "creatives?" + cs.str()
}
