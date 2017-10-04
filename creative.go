package atomx

type Creative struct {
	ID         int64               `json:"id"`
	State      *string             `json:"state,omitempty"`
	Name       *string             `json:"name,omitempty"`
	Advertiser *AdvertiserRelation `json:"advertiser,omitempty"`

	Attributes  CreativeAttributes `json:"attributes,omitempty"`
	Bans        CreativeBans       `json:"bans,omitempty"`
	CategoryID  *int64             `json:"category,omitempty"`
	ContentType *string            `json:"content_type,omitempty"`
	FinalURL    *string            `json:"final_url,omitempty"`
	Javascript  *string            `json:"javascript,omitempty"`
	HTML        *string            `json:"html,omitempty"`
	Size        *SizeRelation      `json:"size,omitempty"`
	SizeBytes   *int64             `json:"size_bytes,omitempty"`
	Types       CreativeTypes      `json:"types,omitempty"`
	Checks      CreativeChecks     `json:"checks,omitempty"`
	Https       *bool              `json:"https,omitempty"`

	AuditedAt *Time `json:"audited_at,omitempty"`
	CheckedAt *Time `json:"checked_at,omitempty"`
	UpdatedAt *Time `json:"updated_at,omitempty"`

	Notify     *bool `json:"notify,omitempty"`
	ModifiedBy int64 `json:"modified_by,omitempty"`
}
