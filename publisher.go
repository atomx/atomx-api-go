package atomx

type Publisher struct {
	ID       int64            `json:"id"`
	State    string           `json:"state,omitempty"`
	Name     string           `json:"name,omitempty"`
	RevShare float64          `json:"revshare,omitempty"`
	Network  *NetworkRelation `json:"network,omitempty"`
}
