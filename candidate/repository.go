package candidate

type Repository interface {
	GetAllCandidates() ([]Candidate, error)
	CreateCandidate(candidate *Candidate) error
	UpdateCandidate(candidate *Candidate) error
	DeleteCandidate(id uint64) error
	GetCandidateByID(id uint64) (*Candidate, error)
}
