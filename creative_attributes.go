package atomx

import (
	"strconv"
	"strings"
)

type CreativeAttributes []CreativeAttribute

func (cas CreativeAttributes) Has(id int64) bool {
	for _, da := range cas {
		if da.Id == id {
			return true
		}
	}

	return false
}

func (cas CreativeAttributes) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, da := range cas {
		ids = append(ids, strconv.FormatInt(da.Id, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}
