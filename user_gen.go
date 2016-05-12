package atomx

import (
	"strconv"
	"strings"
)

func (this User) path() string {
	if this.ID > 0 {
		return "user/" + strconv.FormatInt(this.ID, 10)
	} else {
		return "user"
	}
}

type userResponse struct {
	Success bool   "json:\"success\""
	Error   string "json:\"error\""
	User    *User  "json:\"user\""
}

func (this userResponse) err() error {
	if !this.Success {
		return &ApiError{Message: this.Error}
	}

	return nil
}

func (this *User) response() response {
	return &userResponse{
		User: this,
	}
}

type UsersList struct {
	List
	Users []User "json:\"users\""
}

func (this UsersList) path() string {
	return "users?" + this.str()
}

type Users []User

func (this Users) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this Users) Has(id int64) bool {
	for _, x := range this {
		if x.ID == id {
			return true
		}
	}

	return false
}

func (this *Users) Add(y User) {
	*this = append(*this, y)
}

func (this *Users) Remove(id int64) {
	for i, x := range *this {
		if x.ID == id {
			*this = append((*this)[:i], (*this)[i+1:]...)
			return
		}
	}
}

type UserRelation struct {
	User
}

func (this *UserRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}
