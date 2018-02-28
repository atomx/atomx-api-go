package atomx

type Network struct {
	ID              int64                  `json:"id"`
	Name            string                 `json:"name,omitempty"`
	SellerProfile   *SellerProfileRelation `json:"seller_profile,omitempty"`
	OnlyBuyFromSelf *bool                  `json:"only_buy_from_self,omitempty"`
}
