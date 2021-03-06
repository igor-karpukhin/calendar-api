package interviewer

type Repository interface {
	GetAllInterviewers() ([]Interviewer, error)
	CreateInterviewer(candidate *Interviewer) error
	UpdateInterviewer(candidate *Interviewer) error
	DeleteInterviewer(id uint64) error
	GetInterviewerByID(id uint64) (*Interviewer, error)
}
