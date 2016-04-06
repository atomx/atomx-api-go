package atomx

type Placement struct {
	ID    int64         `json:"id"`
	State string        `json:"state,omitempty"`
	Name  string        `json:"name,omitempty"`
	Site  *SiteRelation `json:"site,omitempty"`
}
