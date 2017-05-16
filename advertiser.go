package atomx

type Advertiser struct {
	ID      int64            `json:"id"`
	Name    string           `json:"name,omitempty"`
	Network *NetworkRelation `json:"network,omitempty"`

	// These fields are for atomx internal use.
	Trusted *bool `json:"trusted,omitempty"`
	IsDSP   *bool `json:"is_dsp,omitempty"`
}
