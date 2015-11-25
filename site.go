package atomx

import (
	"strconv"
)

type Site struct {
	Id       int `json:"id"`
	DomainId int `json:"domain"`
}

func (s *Site) Path() string {
	return "site/" + strconv.FormatInt(int64(s.Id), 10)
}
