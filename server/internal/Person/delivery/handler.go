package delivery

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	personUsecase "server/server/internal/Person/usecase"
	"server/server/internal/domain/dto"
	mw "server/server/internal/middleware"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//Result struct
type Result struct {
	Body interface{}
}

//RespError struct
type RespError struct {
	Err string
}

//PersonHandler handles requests connectded to persons
type PersonHandler struct {
	persons personUsecase.PersonUsecaseI
	logger  *mw.ACLog
}

//NewPersonHandler creates new person handler
func NewPersonHandler(persons personUsecase.PersonUsecaseI, logger *mw.ACLog) *PersonHandler {
	return &PersonHandler{
		persons: persons,
		logger:  logger,
	}
}

//RegisterHandler registers api of person info
func (handler *PersonHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/people", handler.GetPersonList).Methods(http.MethodGet)
	router.HandleFunc("/api/people/age/{age:[0-9]+}", handler.GetPersonByAgeList).Methods(http.MethodGet)
	router.HandleFunc("/api/people/gender/{gender}", handler.GetPersonByGenderList).Methods(http.MethodGet)
	router.HandleFunc("/api/people/nation/{nation}", handler.GetPersonByNationList).Methods(http.MethodGet)
	router.HandleFunc("/api/people/limit/{limit:[0-9]+}", handler.GetPersonWithLimitList).Methods(http.MethodGet)
	router.HandleFunc("/api/people/{id:[0-9]+}", handler.DeletePerson).Methods(http.MethodDelete)
	router.HandleFunc("/api/people/{id:[0-9]+}", handler.UpdatePerson).Methods(http.MethodPatch)
	router.HandleFunc("/api/people", handler.CreatePerson).Methods(http.MethodPost)
}

func (handler *PersonHandler) GetPersonList(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	pers, err := handler.persons.GetPersons()

	if err != nil {
		handler.logger.LogError("problems with getting people", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := pers

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *PersonHandler) GetPersonByAgeList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	strage, ok := vars["age"]
	if !ok {
		handler.logger.LogError("problems with parameters", errors.New("age is missing in parameters"), w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	age64, err := strconv.ParseUint(strage, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		handler.logger.LogError("problems with parameters", errors.New("age is not number"), w.Header().Get("request-id"), r.URL.Path)
		return
	}

	age := uint(age64)

	pers, err := handler.persons.GetPersonsByAge(age)

	if err != nil {
		handler.logger.LogError("problems with getting people", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := pers

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *PersonHandler) GetPersonByGenderList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	gender, ok := vars["gender"]
	if !ok {
		handler.logger.LogError("problems with parameters", errors.New("gender is missing in parameters"), w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pers, err := handler.persons.GetPersonsByGender(gender)

	if err != nil {
		handler.logger.LogError("problems with getting people", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := pers

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *PersonHandler) GetPersonByNationList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	nation, ok := vars["nation"]
	if !ok {
		handler.logger.LogError("problems with parameters", errors.New("nation is missing in parameters"), w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pers, err := handler.persons.GetPersonsByNation(strings.ToUpper(nation))

	if err != nil {
		handler.logger.LogError("problems with getting people", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := pers

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *PersonHandler) GetPersonWithLimitList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	strlimit, ok := vars["limit"]
	if !ok {
		handler.logger.LogError("problems with parameters", errors.New("limit is missing in parameters"), w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limit64, err := strconv.ParseUint(strlimit, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		handler.logger.LogError("problems with parameters", errors.New("limit is not number"), w.Header().Get("request-id"), r.URL.Path)
		return
	}

	limit := uint(limit64)

	pers, err := handler.persons.GetPersonsWithLimit(limit)

	if err != nil {
		handler.logger.LogError("problems with getting people", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := pers

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *PersonHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		handler.logger.LogError("problems with parameters", errors.New("id is missing in parameters"), w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id64, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		handler.logger.LogError("problems with parameters", errors.New("id is not number"), w.Header().Get("request-id"), r.URL.Path)
		return
	}

	id := uint(id64)

	err = handler.persons.DeletePerson(id)
	if err != nil {
		handler.logger.LogError("problems deleting person", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *PersonHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		handler.logger.LogError("problems with parameters", errors.New("id is missing in parameters"), w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id64, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		handler.logger.LogError("problems with parameters", errors.New("id is not number"), w.Header().Get("request-id"), r.URL.Path)
		return
	}

	id := uint(id64)

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatePerson := &dto.Person{ID: id}
	err = json.Unmarshal(jsonbody, &updatePerson)
	if err != nil {
		handler.logger.LogError("prbolems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.persons.UpdatePerson(updatePerson)
	if err != nil {
		if err == dto.ErrNotFound {
			handler.logger.LogError("person not found", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		handler.logger.LogError("problems updating person", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *PersonHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Header.Get("Content-Type") != "application/json" {
		handler.logger.LogError("bad content-type", errors.New("bad content-type"), w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqPerson := dto.Person{}

	jsonbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(jsonbody, &reqPerson)

	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := handler.persons.CreatePerson(&reqPerson)
	if err != nil {
		handler.logger.LogError("problems with creating user", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	body := &dto.RespID{ID: id}

	err = json.NewEncoder(w).Encode(&Result{Body: body})
	if err != nil {
		handler.logger.LogError("problems marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
