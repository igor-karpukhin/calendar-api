package interviewer

import "gopkg.in/mgo.v2"

type MongoInterviewerRepository struct {
	session *mgo.Session
	dbName  string
}

func NewMongoRepository(session *mgo.Session, dbName string) *MongoInterviewerRepository {
	return &MongoInterviewerRepository{
		session: session,
		dbName:  dbName,
	}
}

func (*MongoInterviewerRepository) GetAllInterviewers() ([]Interviewer, error) {
	return nil, nil
}

func (*MongoInterviewerRepository) CreateInterviewer(interviewer *Interviewer) error {
	return nil
}

func (*MongoInterviewerRepository) UpdateInterviewer(interviewer *Interviewer) error {
	return nil
}

func (*MongoInterviewerRepository) DeleteInterviewer(interviewer *Interviewer) error {
	return nil
}

func (*MongoInterviewerRepository) GetInterviewerByID(id uint64) (*Interviewer, error) {
	return nil, nil
}
