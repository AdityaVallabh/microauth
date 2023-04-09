package server

import (
	"log"
	"microauth/internal/server/models/users"
	"net/http"
)

func (s *Server) handleSignup() http.HandlerFunc {
	type request struct {
		Email    string
		Password string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := s.decode(w, r, req); err != nil {
			log.Println("error decoding signup body", err.Error())
			s.respond(w, r, http.StatusBadRequest, "invalid body")
			return
		}
		if err := validate(req.Email, req.Password); err != nil {
			log.Printf("signup email: %s, validation error: %s\n", req.Email, err.Error())
			s.respond(w, r, http.StatusBadRequest, err.Error())
			return
		}
		if _, err := users.Manager.Find(req.Email); err == nil {
			log.Printf("signup email %s already exists", req.Email)
			s.respond(w, r, http.StatusForbidden, "email exists")
			return
		}
		if err := s.Auth.SavePassword(req.Email, req.Password); err != nil {
			log.Printf("error setting password: %s", err.Error())
			s.respond(w, r, http.StatusInternalServerError, "unable to signup")
			return
		}
		s.respond(w, r, http.StatusOK, "signed up")
	}
}
