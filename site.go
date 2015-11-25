package atomx

import (
	"strconv"
)

type Site struct {
	Id       int64 `json:"id"`
	DomainId int64 `json:"domain"`
}

func (s *Site) Path() string {
	return "site/" + strconv.FormatInt(s.Id, 10)
}
