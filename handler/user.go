package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/m-neves/goclock/middleware"
	"github.com/m-neves/goclock/service"
)

type UserHandler struct {
	sm *http.ServeMux
	us service.UserServiceInterface
}

func NewUserHandler(sm *http.ServeMux, us service.UserServiceInterface) *UserHandler {
	return &UserHandler{sm, us}
}

func (uh *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		uh.getHandler(w, r)
	default:
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (uh *UserHandler) getHandler(w http.ResponseWriter, r *http.Request) {
	switch params := strings.Split(r.URL.Path, "/")[2:]; len(params) {
	case 1:
		id, err := strconv.Atoi(params[0])

		if err != nil {
			http.Error(w, "Bad parameter", http.StatusBadRequest)
			return
		}

		uh.findById(w, r, id)
	default:
		http.Error(w, "", http.StatusNotFound)
		return
	}
}

func (uh *UserHandler) findById(w http.ResponseWriter, r *http.Request, id int) {
	user, err := uh.us.FindById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d, err := json.Marshal(user)

	if err != nil {
		http.Error(w, "Unable to unmarshal user", http.StatusInternalServerError)
		return
	}

	w.Write(d)
}

func (uh *UserHandler) SetupRoutes() {
	am := &middleware.AuthMiddleware{Next: uh}

	uh.sm.Handle("/user", am)
	uh.sm.Handle("/user/", am)
}
