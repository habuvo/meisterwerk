package entities

type Status string

func (s Status) Validate() bool {
	if s == "done" || s == "pending" || s == "in progress" {
		return true
	}

	return false
}

type Event struct {
	ID      string `json:"id"`
	Title   string `json:"title`
	Start   string `json:"start_time"`
	End     string `json:"end_time"`
	Address string `json:"address"`
	Status  Status `json:"status"`
}
