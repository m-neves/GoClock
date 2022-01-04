package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/m-neves/goclock/service"
)

const (
	authorizationSchema = "Bearer"
)

type AuthMiddleware struct {
	Next http.Handler
}

type ContextKey string

const ContextUserKey ContextKey = "User-Id"

func (am *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	if token == "" {
		http.Error(w, "Missing identification token", http.StatusForbidden)
		return
	}

	token = strings.TrimSpace(token[len(authorizationSchema):])

	jwt := service.NewJWTService()
	id, err := jwt.GetIDFromToken(token)

	if err != nil {
		http.Error(w, "Invalid token with message"+err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(r.Context(), ContextUserKey, int(id))

	am.Next.ServeHTTP(w, r.WithContext(ctx))
}
