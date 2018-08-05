package candidate

import (
	"database/sql"

	"fmt"

	"github.com/igor-karpukhin/calendar-api/timeslot"
	"go.uber.org/zap"
)

type PostgresCandidateRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func NewPostgresCandidateRepository(db *sql.DB, log *zap.Logger) *PostgresCandidateRepository {
	return &PostgresCandidateRepository{
		db:  db,
		log: log,
	}
}

func (p *PostgresCandidateRepository) GetAllCandidates() ([]Candidate, error) {
	return nil, nil
}

func (p *PostgresCandidateRepository) CreateCandidate(candidate *Candidate) error {

	rows, err := p.db.Query(`INSERT INTO ki.candidates (name) VALUES ($1) RETURNING id`, candidate.Name)
	if err != nil {
		return err
	}
	var cId int
	rows.Next()
	rows.Scan(&cId)
	fmt.Println("> scan id", cId)

	for _, ts := range candidate.Availability {
		_, err := p.db.Exec(`INSERT INTO ki.time_slots (date_from, date_to, ca_id) 
							VALUES ($1, $2, $3)`, ts.From, ts.To, cId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *PostgresCandidateRepository) UpdateCandidate(candidate *Candidate) error {
	panic("implement me")
}

func (p *PostgresCandidateRepository) DeleteCandidate(id uint64) error {
	_, err := p.db.Exec(`DELETE FROM ki.candidates WHERE ki.candidates.id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresCandidateRepository) GetCandidateByID(id uint64) (*Candidate, error) {
	res, err := p.db.Query(`SELECT id, name FROM ki.candidates WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	var result Candidate
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
