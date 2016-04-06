package atomx

type Site struct {
	ID        int64              `json:"id"`
	State     string             `json:"state,omitempty"`
	Name      string             `json:"name,omitempty"`
	Domain    *DomainRelation    `json:"domain,omitempty"`
	Publisher *PublisherRelation `json:"publisher,omitempty"`
}
