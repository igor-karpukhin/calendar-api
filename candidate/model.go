package candidate

import (
	"fmt"

	"errors"

	"github.com/igor-karpukhin/calendar-api/timeslot"
)

type Candidate struct {
	ID           uint64              `json:"id"`
	Name         string              `json:"name"`
	Availability []timeslot.TimeSlot `json:"availability"`
}

func (c *Candidate) Validate() (error, []string) {
	var incorrectFields []string
	if len(c.Name) == 0 {
		incorrectFields = append(incorrectFields, "name. Empty")
	}

	if len(c.Availability) == 0 {
		incorrectFields = append(incorrectFields, "availability. No entries")
	} else {
		for i, slot := range c.Availability {
			if slot.From.After(slot.To) {
				incorrectFields = append(incorrectFields, fmt.Sprintf("timeslot. #%d 'from' > 'to'", i))
			}
			if slot.From.Minute() != 0 {
				incorrectFields = append(incorrectFields,
					fmt.Sprintf("timeslot. #%d 'from' time should be rounded to the hour", i))
			}
			if slot.To.Minute() != 0 {
				incorrectFields = append(incorrectFields,
					fmt.Sprintf("timeslot. #%d 'to' time hsould be rounded to the hour", i))
			}
		}
	}

	if len(incorrectFields) > 0 {
		return errors.New("invalid"), incorrectFields
	}
	return nil, nil
}
