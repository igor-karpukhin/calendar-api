package candidate

import (
	"net/http"

	"strconv"

	"encoding/json"

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
		log:        log,
	}
}

func (c *Controller) BuildRoutes(r *mux.Router) *mux.Router {
	sub := r.PathPrefix("/candidate").Subrouter()

	sub.HandleFunc("/", c.ListCandidates).Methods("GET")
	sub.HandleFunc("/{id:[0-9]+}", c.GetCandidate).Methods("GET")
	sub.HandleFunc("/{id:[0-9]+}/interview", c.Interview).Methods("GET")
	sub.HandleFunc("/", c.CreateCandidate).Methods("POST")
	sub.HandleFunc("/{id:[0-9]+}", c.DeleteCandidate).Methods("DELETE")
	return sub
}

func (c *Controller) ListCandidates(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not found"))
}

func (c *Controller) GetCandidate(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	c.log.Info("GetCandidate", zap.String("id", param))
	id, err := strconv.Atoi(param)
	if err != nil {
		c.log.Info("not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res, err := c.repository.GetCandidateByID(uint64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.log.Error("getCandidate error", zap.Error(err))
		return
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
	return
}

func (c *Controller) Interview(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) CreateCandidate(w http.ResponseWriter, r *http.Request) {
	c.log.Info("CreateCandidate")
	candidate := &Candidate{}
	err := json.NewDecoder(r.Body).Decode(candidate)
	if err != nil {
		c.log.Error("unable to decode json", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Add("Content-type", "application/json")
	err, missingFields := candidate.Validate()
	if err != nil {
		c.log.Error("invalid candidate object", zap.Any("obj", *candidate))
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error(), "missing_fields": missingFields})
		return
	}

	err = c.repository.CreateCandidate(candidate)
	if err != nil {
		c.log.Error("unable to create candidate", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) DeleteCandidate(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.repository.DeleteCandidate(uint64(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
