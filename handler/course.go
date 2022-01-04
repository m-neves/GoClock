package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/middleware"
	"github.com/m-neves/goclock/service"
)

type courseHandler struct {
	sm *http.ServeMux
	cs service.CourseServiceInterface
}

func NewCouseHandler(sm *http.ServeMux, cs service.CourseServiceInterface) *courseHandler {
	return &courseHandler{sm, cs}
}

func (ch *courseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ch.getHandler(w, r)
	case http.MethodPost:
		ch.postHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (ch *courseHandler) getHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.ContextUserKey)
	courses, err := ch.cs.FindAll(userId.(int))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d, err := json.Marshal(courses)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(d)
}

func (ch *courseHandler) postHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.ContextUserKey)

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	course := &data.Course{}
	err = json.Unmarshal(d, course)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ch.cs.Create(course, userId.(int))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ch *courseHandler) SetupRoutes() {
	am := &middleware.AuthMiddleware{Next: ch}

	ch.sm.Handle("/course", am)
	ch.sm.Handle("/course/", am)
}
