package atomx

type Ssp struct {
	ID            int64                  `json:"id"`
	SellerProfile *SellerProfileRelation `json:"seller_profile,omitempty"`
}
