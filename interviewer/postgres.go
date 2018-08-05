package interviewer

import (
	"database/sql"
	"fmt"

	"github.com/igor-karpukhin/calendar-api/timeslot"
	"go.uber.org/zap"
)

type PostgresInterviewersRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func NewPostgresInterviewersRepository(db *sql.DB, log *zap.Logger) *PostgresInterviewersRepository {
	return &PostgresInterviewersRepository{
		db:  db,
		log: log,
	}
}

func (p *PostgresInterviewersRepository) GetAllInterviewers() ([]Interviewer, error) {
	return nil, nil
}

func (p *PostgresInterviewersRepository) CreateInterviewer(interviewer *Interviewer) error {

	rows, err := p.db.Query(`INSERT INTO ki.interviewers (name) VALUES ($1) RETURNING id`, interviewer.Name)
	if err != nil {
		return err
	}
	var cId int
	rows.Next()
	rows.Scan(&cId)
	fmt.Println("> scan id", cId)

	for _, ts := range interviewer.Availability {
		_, err := p.db.Exec(`INSERT INTO ki.time_slots (date_from, date_to, ca_id) 
							VALUES ($1, $2, $3)`, ts.From, ts.To, cId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *PostgresInterviewersRepository) UpdateInterviewer(interviewer *Interviewer) error {
	panic("implement me")
}

func (p *PostgresInterviewersRepository) DeleteInterviewer(id uint64) error {
	_, err := p.db.Exec(`DELETE FROM ki.interviewers WHERE ki.interviewers.id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresInterviewersRepository) GetInterviewerByID(id uint64) (*Interviewer, error) {
	res, err := p.db.Query(`SELECT id, name FROM ki.interviewers WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	var result Interviewer
	res.Next()
	err = res.Scan(&result.ID, &result.Name)
	if err != nil {
		return nil, err
	}

	res, err = p.db.Query(`SELECT date_from, date_to FROM ki.time_slots WHERE ca_id = $1`, id)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var ts timeslot.TimeSlot
		err := res.Scan(
			&ts.From,
			&ts.To,
		)
		if err != nil {
			return nil, err
		}
		result.Availability = append(result.Availability, ts)
	}

	return &result, nil
}
