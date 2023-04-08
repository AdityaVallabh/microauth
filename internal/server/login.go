package server

import (
	"net/http"
)

func (s *Server) handleLogin() http.HandlerFunc {
	type request struct {
		Email    string
		Password string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if s.decode(w, r, req) != nil {
			s.respond(w, r, http.StatusBadRequest, "invalid body")
			return
		}
		if err := validate(req.Email, req.Password); err != nil {
			s.respond(w, r, http.StatusBadRequest, err.Error())
			return
		}
		ok := s.Auth.CheckPassword(req.Email, req.Password)
		if !ok {
			s.respond(w, r, http.StatusUnauthorized, "invalid login")
			return
		}
		token, err := s.Auth.GenerateToken(req.Email)
		if err != nil {
			s.respond(w, r, http.StatusInternalServerError, "smth went wrong")
			return
		}
		http.SetCookie(w, &http.Cookie{Name: "auth_token", Value: string(token)})
		s.respond(w, r, http.StatusOK, "logged in")
	}
}
func (s *Server) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "auth_token", Value: "", MaxAge: -1})
	s.respond(w, r, http.StatusNoContent, "")
}
