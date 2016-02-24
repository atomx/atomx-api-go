package atomx

import (
	"strconv"
)

type User struct {
	ID int64 `json:"id"`
}

func (u User) path() string {
	return "user/" + strconv.FormatInt(u.ID, 10)
}
