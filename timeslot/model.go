package timeslot

import "time"

type TimeSlot struct {
	From time.Time `json:"from",bson:"from"`
	To   time.Time `json:"to",bson:"to"`
}
