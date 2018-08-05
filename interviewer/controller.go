package interviewer

import (
	"net/http"

	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Controller struct {
	repository Repository
	log        *zap.Logger
}

func NewController(repository Repository, log *zap.Logger) *Controller {
	return &Controller{
		repository: repository,
	}
}

func (c *Controller) BuildRoutes(r *mux.Router) *mux.Router {
	sub := r.PathPrefix("/interviewer").Subrouter()

	sub.HandleFunc("/", c.ListInterviewers).Methods("GET")
	sub.HandleFunc("/{id:[0-9]+}", c.GetInterviewer).Methods("GET")
	sub.HandleFunc("/{id:[0-9]+}/interview", c.CreateInterviewer).Methods("GET")
	sub.HandleFunc("/", c.CreateInterviewer).Methods("POST")
	sub.HandleFunc("/{id:[0-9]+}", c.DeleteInterviewer).Methods("DELETE")
	return sub
}

func (c *Controller) ListInterviewers(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) GetInterviewer(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	c.log.Info("GetInterviewer", zap.String("id", param))
	id, err := strconv.Atoi(param)
	if err != nil {
		c.log.Info("not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res, err := c.repository.GetInterviewerByID(uint64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.log.Error("getInterviewer error", zap.Error(err))
		return
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
	return
}

func (c *Controller) CreateInterviewer(w http.ResponseWriter, r *http.Request) {
	c.log.Info("CreateInterviewer")
	interviewer := &Interviewer{}
	err := json.NewDecoder(r.Body).Decode(interviewer)
	if err != nil {
		c.log.Error("unable to decode json", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Add("Content-type", "application/json")
	err, missingFields := interviewer.Validate()
	if err != nil {
		c.log.Error("invalid interviewer object", zap.Any("obj", *interviewer))
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error(), "missing_fields": missingFields})
		return
	}

	err = c.repository.CreateInterviewer(interviewer)
	if err != nil {
		c.log.Error("unable to create interviewer", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) DeleteInterviewer(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.repository.DeleteInterviewer(uint64(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
