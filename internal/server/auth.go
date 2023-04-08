package server

import (
	"context"
	"microauth/internal/server/models/users"
	"net/http"
)

type Ckey string

var (
	emailKey      = Ckey("email")
	sessionExpiry = Ckey("sessionExpiry")
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
		var user users.User
		err = s.DB.Find(email, &user)
		if err != nil {
			s.respond(w, r, http.StatusUnauthorized, "not logged in")
			return
		}
		r = r.WithContext(
			context.WithValue(
				context.WithValue(r.Context(), emailKey, email),
				sessionExpiry,
				s.Auth.GetExpiry(token.Value),
			),
		)
		next(w, r)
	})
}
