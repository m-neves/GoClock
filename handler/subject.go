package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/middleware"
	"github.com/m-neves/goclock/service"
)

type SubjectHandler struct {
	sm *http.ServeMux
	ss service.SubjectServiceInterface
}

func NewSubjectHandler(sm *http.ServeMux, ss service.SubjectServiceInterface) *SubjectHandler {
	return &SubjectHandler{sm, ss}
}

func (sh *SubjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		sh.getHandler(w, r)
	case http.MethodPost:
		sh.postHandler(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

func (sh *SubjectHandler) getHandler(w http.ResponseWriter, r *http.Request) {
	switch params := strings.Split(r.URL.Path, "/")[2:]; len(params) {
	case 0:
		sh.findAll(w, r)
	case 1:
		id, err := strconv.Atoi(params[0])

		if err != nil {
			http.Error(w, "Bad parameter", http.StatusBadRequest)
			return
		}

		sh.findById(w, r, id)
	default:
		http.Error(w, "", http.StatusNotFound)
		return
	}
}

func (sh *SubjectHandler) findById(w http.ResponseWriter, r *http.Request, id int) {
	userId := r.Context().Value(middleware.ContextUserKey)

	subject, err := sh.ss.FindById(id, userId.(int))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d, err := json.Marshal(subject)

	if err != nil {
		http.Error(w, "Unable to unmarshal subject", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(d)
}

func (sh *SubjectHandler) findAll(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.ContextUserKey)

	subjects, err := sh.ss.FindAll(userId.(int))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d, err := json.Marshal(subjects)

	if err != nil {
		http.Error(w, "Unable to unmarshal subjects", http.StatusInternalServerError)
		return
	}

	w.Write(d)
}

func (sh *SubjectHandler) postHandler(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Invalid subject body", http.StatusBadRequest)
		return
	}

	subject := &data.Subject{}

	json.Unmarshal(d, subject)

	id := r.Context().Value(middleware.ContextUserKey)

	err = sh.ss.Create(subject.Name, id.(int))

	if err != nil {
		http.Error(w, "Could not create subject with message: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (sh *SubjectHandler) SetupRoutes() {
	am := &middleware.AuthMiddleware{Next: sh}

	sh.sm.Handle("/subject", am)
	sh.sm.Handle("/subject/", am)
}
