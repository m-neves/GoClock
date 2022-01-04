package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/m-neves/goclock/middleware"
	"github.com/m-neves/goclock/service"
)

type StudyPlanHandlerInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	SetupRoutes()
}

type studyPlanHandler struct {
	sm  *http.ServeMux
	sps service.StudyPlanServiceInterface
}

func NewStudyPlanHandler(sm *http.ServeMux, sps service.StudyPlanServiceInterface) StudyPlanHandlerInterface {
	return &studyPlanHandler{sm, sps}
}

func (sph *studyPlanHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		sph.getHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (sph *studyPlanHandler) getHandler(w http.ResponseWriter, r *http.Request) {
	switch params := strings.Split(r.URL.Path, "/")[2:]; len(params) {
	case 1:
		studyPlanId, err := strconv.ParseUint(params[0], 10, 32)

		if err != nil {
			http.Error(w, "Bad parameter", http.StatusBadRequest)
			return
		}

		sph.findAll(w, r, int(studyPlanId))
	}
}

func (sph *studyPlanHandler) findAll(w http.ResponseWriter, r *http.Request, studyPlanId int) {
	ctxUserId := r.Context().Value(middleware.ContextUserKey)
	userId := uint(ctxUserId.(int))

	id := uint(studyPlanId)

	studyPlanCourses, err := sph.sps.FindById(id, userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d, err := json.Marshal(studyPlanCourses)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(d)
}

func (sph *studyPlanHandler) SetupRoutes() {
	am := &middleware.AuthMiddleware{Next: sph}

	sph.sm.Handle("/study_plan", am)
	sph.sm.Handle("/study_plan/", am)
}
