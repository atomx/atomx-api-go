package atomx

type HistoryUser struct {
	ID        int64  `json:"id"`
	IP        string `json:"ip"`
	Name      string `json:"name"`
	UserAgent string `json:"user_agent"`
}

type HistoryEntry struct {
	Changes map[string][]interface{} `json:"changes"`
	User    HistoryUser              `json:"user"`
}

type History struct {
	// Input.
	Obj Resource `json:"-"`

	// Output.
	History []HistoryEntry `json:"history"`

	// Both.
	List
}

func (h History) path() string {
	return "history/" + h.Obj.path() + "?" + h.str()
}
