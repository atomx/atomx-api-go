package atomx

type Advertiser struct {
	ID      int64            `json:"id"`
	Name    string           `json:"name,omitempty"`
	IsDSP   bool             `json:"is_dsp,omitempty"`
	Network *NetworkRelation `json:"network,omitempty"`
}
