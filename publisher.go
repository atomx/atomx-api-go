package atomx

type Publisher struct {
	ID            int64                  `json:"id"`
	State         string                 `json:"state,omitempty"`
	Name          string                 `json:"name,omitempty"`
	RevShare      float64                `json:"revshare,omitempty"`
	Network       *NetworkRelation       `json:"network,omitempty"`
	SellerProfile *SellerProfileRelation `json:"seller_profile,omitempty"`

	// These fields are for atomx internal use.
	Banned *bool        `json:"banned,omitempty"`
	Ssp    *SspRelation `json:"ssp,omitempty"`
}
