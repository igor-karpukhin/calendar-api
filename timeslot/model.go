package timeslot

import "time"

type TimeSlot struct {
	ID   int       `json:"id"`
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}
