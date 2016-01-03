package atomx

type Domains struct {
	List
	Domains []Domain `json:"domains"`
}

func (ds Domains) path() string {
	return "domains?" + ds.str()
}
