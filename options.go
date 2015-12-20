package atomx

import (
	"strconv"
	"strings"
)

type Options struct {
	Depth  int
	Expand []string
}

func (o *Options) str() string {
	if o == nil {
		return "depth=0"
	}

	return "depth=" + strconv.Itoa(o.Depth) + "&expand=" + strings.Join(o.Expand, ",")
}
