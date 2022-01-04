package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/middleware"
	"github.com/m-neves/goclock/service"
)

const (
	addSubjects = "addSubjects"
)

type EventHandler struct {
	sm *http.ServeMux
	es service.EventServiceInterface
}

func NewEventHandler(sm *http.ServeMux, es service.EventServiceInterface) *EventHandler {
	return &EventHandler{sm, es}
}

func (eh *EventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		eh.postHandler(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

func (eh *EventHandler) postHandler(w http.ResponseWriter, r *http.Request) {
	switch params := strings.Split(r.URL.Path, "/")[2:]; len(params) {
	case 0:
		eh.create(w, r)
	case 1:
		switch params[0] {
		case addSubjects:
			eh.createWithSubject(w, r)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (eh *EventHandler) create(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	event := &data.Event{}

	err = json.Unmarshal(d, event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := r.Context().Value(middleware.ContextUserKey)
	event.UserId = id.(int)

	err = eh.es.Create(event)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (eh *EventHandler) createWithSubject(w http.ResponseWriter, r *http.Request) {

	eventSubjects := &data.EventSubjects{}

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, eventSubjects)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = eh.es.AddSubjects(eventSubjects)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (eh *EventHandler) SetupRoutes() {
	am := &middleware.AuthMiddleware{Next: eh}

	eh.sm.Handle("/event", am)
	eh.sm.Handle("/event/", am)
}
