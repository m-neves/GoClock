package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/middleware"
	"github.com/m-neves/goclock/service"
)

const (
	dateLayout  = "2006-01-02"
	sumByPeriod = "sumByPeriod"
)

type EntryHandler struct {
	sm *http.ServeMux
	es service.EntryServiceInterface
}

func NewEntryHandler(sm *http.ServeMux, es service.EntryServiceInterface) *EntryHandler {
	return &EntryHandler{sm, es}
}

func (eh *EntryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		eh.getHandler(w, r)
	case http.MethodPost:
		eh.postHandler(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

func (eh *EntryHandler) getHandler(w http.ResponseWriter, r *http.Request) {
	switch params := strings.Split(r.URL.Path, "/")[2:]; len(params) {
	case 1:
		switch params[0] {
		case sumByPeriod:
			eh.sumByPeriod(w, r)
		default:
			id, err := strconv.Atoi(params[0])

			if err != nil {
				http.Error(w, "Bad parameter", http.StatusBadRequest)
				return
			}

			eh.findById(w, r, id)
		}

	default:
		http.Error(w, "", http.StatusNotFound)
		return
	}
}

func (eh *EntryHandler) findById(w http.ResponseWriter, r *http.Request, id int) {
	userId := r.Context().Value(middleware.ContextUserKey)

	entry, err := eh.es.FindById(id, userId.(int))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	d, err := json.Marshal(entry)

	if err != nil {
		http.Error(w, "Unable to unmarshal entry", http.StatusInternalServerError)
		return
	}

	w.Write(d)
	w.WriteHeader(http.StatusOK)
}
func (eh *EntryHandler) sumByPeriod(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.ContextUserKey)
	q := r.URL.Query()

	if _, ok := q["start"]; !ok {
		http.Error(w, "Missing start parameter", http.StatusBadRequest)
		return
	}

	if _, ok := q["end"]; !ok {
		http.Error(w, "Missing end parameter", http.StatusBadRequest)
		return
	}

	start, err := time.Parse(dateLayout, q["start"][0])

	if err != nil {
		http.Error(w, "Bad start parameter, should be of type"+dateLayout, http.StatusBadRequest)
		return
	}

	end, err := time.Parse(dateLayout, q["end"][0])

	if err != nil {
		http.Error(w, "Bad end parameter, should be of type"+dateLayout, http.StatusBadRequest)
		return
	}

	entries, err := eh.es.SumByPeriod(start, end, userId.(int))

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	d := entries.TotalDiff()

	w.Write([]byte(d.String()))
}

func (eh *EntryHandler) postHandler(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Invalid entry value", http.StatusBadRequest)
		return
	}

	entry := &data.Entry{}

	id := r.Context().Value(middleware.ContextUserKey)
	entry.UserId = id.(int)

	json.Unmarshal(d, entry)

	err = eh.es.Create(entry)

	if err != nil {
		http.Error(w, "Unable to create entry with message:"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (eh *EntryHandler) SetupRoutes() {
	am := &middleware.AuthMiddleware{Next: eh}

	eh.sm.Handle("/entry/", am)
}
