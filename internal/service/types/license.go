package types

import "time"

// Contract TODO issue#docs
type Contract struct {
	Since      *time.Time `json:"since,omitempty"`
	Until      *time.Time `json:"until,omitempty"`
	Workplaces uint32     `json:"workplace_limits,omitempty"`
	Limits
}

// License TODO issue#docs
type License struct {
	ID       ID `json:"id"`
	Contract `json:"contract"`
}

// Limits TODO issue#docs
type Limits struct {
	Rate     Rate   `json:"rate_limits,omitempty"`
	Requests uint32 `json:"request_limits,omitempty"`
}
