package atomx

import (
	"strconv"
	"strings"
)

type Options struct {
	Depth  int
	Expand []string
	Sort   string
	Extra  []string
}

func (o *Options) str() string {
	if o == nil {
		return "depth=0"
	}

	r := "depth=" + strconv.Itoa(o.Depth)

	if len(o.Expand) > 0 {
		r = r + "&expand=" + strings.Join(o.Expand, ",")
	}

	if o.Sort != "" {
		r = r + "&sort=" + o.Sort
	}

	if len(o.Extra) > 0 {
		r = r + "&" + strings.Join(o.Extra, "&")
	}

	return r
}
