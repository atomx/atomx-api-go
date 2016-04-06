package atomx

import (
	"strconv"
	"strings"
)

func (this Publisher) path() string {
	if this.ID > 0 {
		return "publisher/" + strconv.FormatInt(this.ID, 10)
	} else {
		return "publisher"
	}
}

type publisherResponse struct {
	Success   bool       "json:\"success\""
	Error     string     "json:\"error\""
	Publisher *Publisher "json:\"publisher\""
}

func (this publisherResponse) err() error {
	if !this.Success {
		return &ApiError{Message: this.Error}
	}

	return nil
}

func (this *Publisher) response() response {
	return &publisherResponse{
		Publisher: this,
	}
}

type Publishers struct {
	List
	Publishers []Publisher "json:\"publishers\""
}

func (this Publishers) path() string {
	return "publishers?" + this.str()
}

func (this Publishers) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this.Publishers {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this Publishers) Has(id int64) bool {
	for _, x := range this.Publishers {
		if x.ID == id {
			return true
		}
	}

	return false
}

type PublisherRelation struct {
	Publisher
}

func (this *PublisherRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}
