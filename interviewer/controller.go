package interviewer

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
	sub := r.PathPrefix("/interviewer").Subrouter()

	sub.HandleFunc("/", c.ListInterviewers).Methods("GET")
	sub.HandleFunc("/{id:[0-9]+}", c.GetInterviewer).Methods("GET")
	sub.HandleFunc("/{id:[0-9]+/interview", c.Interview).Methods("GET")
	sub.HandleFunc("/create", c.NewInterviewer).Methods("POST")
	sub.HandleFunc("/{id:[0-9]+}", c.DeleteInterviewer).Methods("DELETE")
	return sub
}

func (c *Controller) ListInterviewers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not found"))
}

func (c *Controller) GetInterviewer(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) Interview(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) NewInterviewer(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) DeleteInterviewer(w http.ResponseWriter, r *http.Request) {

}
