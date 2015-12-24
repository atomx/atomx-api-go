package atomx

import (
	"strconv"
	"strings"
)

type DomainAttributes []DomainAttribute

func (das DomainAttributes) Has(id int64) bool {
	for _, da := range das {
		if da.ID == id {
			return true
		}
	}

	return false
}

func (das DomainAttributes) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, da := range das {
		ids = append(ids, strconv.FormatInt(da.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}
