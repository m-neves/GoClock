package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/service"
)

type LoginHandler struct {
	sm *http.ServeMux
	us service.UserServiceInterface
}

func NewLoginHandler(sm *http.ServeMux, us service.UserServiceInterface) *LoginHandler {
	return &LoginHandler{sm, us}
}

func (lh *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		lh.postHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func (lh *LoginHandler) postHandler(w http.ResponseWriter, r *http.Request) {

	switch params := strings.Split(r.URL.Path, "/")[2:]; len(params) {
	case 0:
		lh.generateToken(w, r)
	case 1:
		lh.createUser(w, r)
	}
}

func (lh *LoginHandler) generateToken(w http.ResponseWriter, r *http.Request) {
	user := &data.User{}
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Bad user format", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, user)

	if err != nil {
		http.Error(w, "Unable to unsmarshal User", http.StatusInternalServerError)
		return
	}

	u, err := lh.us.FindByCredentials(user.Email, user.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if u == nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	jwt := service.NewJWTService()
	token, err := jwt.GenerateToken(uint(u.Id))
	if err != nil {
		http.Error(w, "Unable to generate token", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

func (lh *LoginHandler) createUser(w http.ResponseWriter, r *http.Request) {
	u := &data.User{}
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Can't read user", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, u)

	if err != nil {
		http.Error(w, "Can't marshal user", http.StatusBadRequest)
		return
	}

	err = lh.us.Create(u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (lh *LoginHandler) SetupRoutes() {
	lh.sm.Handle("/login", lh)
	lh.sm.Handle("/login/", lh)
}
