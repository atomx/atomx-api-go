package atomx

type DomainAttributes []DomainAttribute

func (das DomainAttributes) Has(id int64) bool {
	for _, da := range das {
		if da.Id == id {
			return true
		}
	}

	return false
}
