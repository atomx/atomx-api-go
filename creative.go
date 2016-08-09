package atomx

type Creative struct {
	ID          int64               `json:"id"`
	Name        *string             `json:"name,omitempty"`
	Advertiser  *AdvertiserRelation `json:"advertiser,omitempty"`
	Attributes  CreativeAttributes  `json:"attributes,omitempty"`
	AuditedAt   *Time               `json:"audited_at,omitempty"`
	Bans        CreativeBans        `json:"bans,omitempty"`
	Size        *SizeRelation       `json:"size,omitempty"`
	SizeBytes   *int64              `json:"size_bytes,omitempty"`
	Types       CreativeTypes       `json:"types,omitempty"`
	FinalURL    *string             `json:"final_url,omitempty"`
	ContentType *string             `json:"content_type,omitempty"`
	Javascript  *string             `json:"javascript,omitempty"`
	CheckedAt   *Time               `json:"checked_at,omitempty"`
}
