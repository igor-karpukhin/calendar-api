package candidate

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
	repository Repository
}

func NewController(repository Repository) *Controller {
	return &Controller{
		repository: repository,
	}
}

func (c *Controller) BuildRoutes(r *mux.Router) *mux.Router {
	sub := r.PathPrefix("/candidate").Subrouter()

	sub.HandleFunc("/", c.ListCandidates).Methods("GET")
	sub.HandleFunc("/{id:[0-9]+}", c.GetCandidate).Methods("GET")
	sub.HandleFunc("/{id:[0-9]+/interview", c.Interview).Methods("GET")
	sub.HandleFunc("/create", c.NewCandidate).Methods("POST")
	sub.HandleFunc("/{id:[0-9]+}", c.DeleteCandidate).Methods("DELETE")
	return sub
}

func (c *Controller) ListCandidates(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not found"))
}

func (c *Controller) GetCandidate(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) Interview(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) NewCandidate(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) DeleteCandidate(w http.ResponseWriter, r *http.Request) {

}
