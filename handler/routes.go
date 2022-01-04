package handler

import (
	"net/http"
	"os"

	"github.com/m-neves/goclock/repository"
	"github.com/m-neves/goclock/service"
)

type routesHandler struct {
	sm *http.ServeMux
}

func NewRoutesHandler(sm *http.ServeMux) *routesHandler {
	return &routesHandler{sm}
}

func (rh *routesHandler) SetupRoutes() {
	rh.sm.HandleFunc("/bae", func(w http.ResponseWriter, r *http.Request) {
		f, _ := os.ReadFile("splash.txt")
		w.Write(f)
	})

	ur := repository.NewUserRepository()
	us := service.NewUserService(ur)
	uh := NewUserHandler(rh.sm, us)
	uh.SetupRoutes()

	sr := repository.NewSubjectRepository()
	ss := service.NewSubjectService(sr)
	sh := NewSubjectHandler(rh.sm, ss)
	sh.SetupRoutes()

	er := repository.NewEntryRepository()
	es := service.NewEntryService(er)
	eh := NewEntryHandler(rh.sm, es)
	eh.SetupRoutes()

	evr := repository.NewEventRepository()
	evs := service.NewEventService(evr)
	evh := NewEventHandler(rh.sm, evs)
	evh.SetupRoutes()

	cr := repository.NewCourseRepository()
	cs := service.NewCourseService(cr)
	ch := NewCouseHandler(rh.sm, cs)
	ch.SetupRoutes()

	spr := repository.NewStudyPlanRepository()
	sps := service.NewStudyPlanService(spr)
	sph := NewStudyPlanHandler(rh.sm, sps)
	sph.SetupRoutes()

	lh := NewLoginHandler(rh.sm, us)
	lh.SetupRoutes()
}
