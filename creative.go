package atomx

type Creative struct {
	ID         int64               `json:"id"`
	Advertiser *AdvertiserRelation `json:"advertiser,omitempty"`
	Attributes CreativeAttributes  `json:"attributes,omitempty"`
	AuditedAt  *Time               `json:"audited_at,omitempty"`
	Bans       CreativeBans        `json:"bans,omitempty"`
	Size       *SizeRelation       `json:"size,omitempty"`
	SizeBytes  *int64              `json:"size_bytes,omitempty"`
	Types      CreativeTypes       `json:"types,omitempty"`
}
