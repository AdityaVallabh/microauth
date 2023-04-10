package server

import (
	"context"
	"net/http"

	"github.com/AdityaVallabh/microauth/internal/server/models/users"
)

type Ckey string

var (
	emailKey = Ckey("email")
	authKey  = Ckey("authKey")
)

func (s *Server) auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("auth_token")
		if err != nil {
			s.respond(w, r, http.StatusUnauthorized, "not logged in")
			return
		}
		email, ok := s.Auth.Validate(token.Value)
		if !ok {
			s.respond(w, r, http.StatusUnauthorized, "not logged in")
			return
		}
		_, err = users.Manager.Find(email)
		if err != nil {
			s.respond(w, r, http.StatusUnauthorized, "not logged in")
			return
		}
		ctx := context.WithValue(r.Context(), emailKey, email)
		ctx = context.WithValue(ctx, authKey, token.Value)
		r = r.WithContext(ctx)
		next(w, r)
	})
}
