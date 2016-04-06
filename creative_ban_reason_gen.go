package atomx

import (
	"strconv"
	"strings"
)

func (this CreativeBanReason) path() string {
	if this.ID > 0 {
		return "creative_ban_reason/" + strconv.FormatInt(this.ID, 10)
	} else {
		return "creative_ban_reason"
	}
}

type creative_ban_reasonResponse struct {
	Success           bool               "json:\"success\""
	Error             string             "json:\"error\""
	CreativeBanReason *CreativeBanReason "json:\"creative_ban_reason\""
}

func (this creative_ban_reasonResponse) err() error {
	if !this.Success {
		return &ApiError{Message: this.Error}
	}

	return nil
}

func (this *CreativeBanReason) response() response {
	return &creative_ban_reasonResponse{
		CreativeBanReason: this,
	}
}

type CreativeBanReasons struct {
	List
	CreativeBanReasons []CreativeBanReason "json:\"creative_ban_reasons\""
}

func (this CreativeBanReasons) path() string {
	return "creative_ban_reasons?" + this.str()
}

func (this CreativeBanReasons) MarshalJSON() ([]byte, error) {
	var ids []string

	for _, x := range this.CreativeBanReasons {
		ids = append(ids, strconv.FormatInt(x.ID, 10))
	}

	return []byte("[" + strings.Join(ids, ",") + "]"), nil
}

func (this CreativeBanReasons) Has(id int64) bool {
	for _, x := range this.CreativeBanReasons {
		if x.ID == id {
			return true
		}
	}

	return false
}

type CreativeBanReasonRelation struct {
	CreativeBanReason
}

func (this *CreativeBanReasonRelation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(this.ID, 10)), nil
}
