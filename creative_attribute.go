package atomx

import (
	"strconv"
)

type CreativeAttribute struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (ca CreativeAttribute) path() string {
	return "creative_attribute/" + strconv.FormatInt(ca.Id, 10)
}

type creativeAttributeResponse struct {
	Success           bool               `json:"success"`
	Error             string             `json:"error"`
	CreativeAttribute *CreativeAttribute `json:"creative_attribute"`
}

func (car creativeAttributeResponse) err() error {
	if !car.Success {
		return &ApiError{Message: car.Error}
	}

	return nil
}

func (ca *CreativeAttribute) response() response {
	return &creativeAttributeResponse{
		CreativeAttribute: ca,
	}
}
