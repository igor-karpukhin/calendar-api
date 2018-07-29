package candidate

import "github.com/igor-karpukhin/calendar-api/timeslot"

type Candidate struct {
	ID           uint64              `json:"id",bson:"_id"`
	Name         string              `json:"name",bson:"name"`
	Availability []timeslot.TimeSlot `json:"availability",bson:"availability"`
}
