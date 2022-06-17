package entities

import "time"

type Status string

func (s Status) Validate() bool {
	if s == "done" || s == "pending" || s == "in progress" {
		return true
	}

	return false
}

type Event struct {
	ID      string    `json:"id"`
	Title   string    `json:"title`
	Start   time.Time `json:"start_time"`
	End     time.Time `json:"end_time"`
	Address string    `json:"address"`
	Status  Status    `json:"status"`
}

type TransportEvent struct {
	ID      string  `json:"id"`
	Title   *string `json:"title, omitempty"`
	Start   *string `json:"start, omitempty"`
	End     *string `json:"end, omitempty"`
	Address *string `json:"address, omitempty"`
	Status  *Status `json:"status, omitempty"`
}
