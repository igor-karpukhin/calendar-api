package candidate

import "gopkg.in/mgo.v2"

type MongoCandidateRepository struct {
	session *mgo.Session
	dbName  string
}

func NewMongoRepository(session *mgo.Session, dbName string) *MongoCandidateRepository {
	return &MongoCandidateRepository{
		session: session,
		dbName:  dbName,
	}
}

func (*MongoCandidateRepository) GetAllCandidates() ([]Candidate, error) {
	return nil, nil
}

func (*MongoCandidateRepository) CreateCandidate(candidate *Candidate) error {
	return nil
}

func (*MongoCandidateRepository) UpdateCandidate(candidate *Candidate) error {
	return nil
}

func (*MongoCandidateRepository) DeleteCandidate(candidate *Candidate) error {
	return nil
}

func (*MongoCandidateRepository) GetCandidateByID(id uint64) (*Candidate, error) {
	return nil, nil
}
