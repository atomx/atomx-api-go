package atomx

import (
	"strconv"
)

type Network struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (n Network) path() string {
	return "network/" + strconv.FormatInt(n.ID, 10)
}

type networkResponse struct {
	Success bool     `json:"success"`
	Error   string   `json:"error"`
	Network *Network `json:"network"`
}

func (nr networkResponse) err() error {
	if !nr.Success {
		return &ApiError{Message: nr.Error}
	}

	return nil
}

func (n *Network) response() response {
	return &networkResponse{
		Network: n,
	}
}
