package objects

import "time"

// EventStatus defines the status of the event
type EventStatus string

// Some default event status
const (
	Original    EventStatus = "original"
	Cancelled   EventStatus = "cancelled"
	Rescheduled EventStatus = "rescheduled"
)

// TimeSlot for Event
type TimeSlot struct {
	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
}

// Event object for the API
type Event struct {
	ID string `gorm:"primary_key" json:"id,omitempty"`

	// General details
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Website     string `json:"website,omitempty"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Event Slot Duration
	Slot *TimeSlot `gorm:"embedded" json:"slot,omitempty"`

	// Change Status
	Status EventStatus `json:"status,omitempty"`

	// Meta Information
	CreatedOn     time.Time `json:"create_on,omitempty"`
	UpdatedOn     time.Time `json:"updated_up,omitempty"`
	CancelledOn   time.Time `json:"cancelled_on,omitempty"`
	RescheduledOn time.Time `json:"rescheduled_on,omitempty"`
}
