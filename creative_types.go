package atomx

type CreativeTypes []string

func (cts CreativeTypes) Has(t string) bool {
	for _, x := range cts {
		if x == t {
			return true
		}
	}

	return false
}
