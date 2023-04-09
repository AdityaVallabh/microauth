package server

import (
	"net/http"
	"text/template"
	"time"
)

type HomePage struct {
	Title    string
	User     string
	ExpireIn time.Duration
}

func (s *Server) handleHome(templatesDir string) http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles(templatesDir + "/home.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.Context().Value(emailKey).(string)
		sessionExp := s.Auth.GetExpiry(r.Context().Value(authKey).(string))
		w.WriteHeader(http.StatusOK)
		if err := tmpl.Execute(w, HomePage{
			Title:    "myServer",
			User:     email,
			ExpireIn: time.Until(sessionExp),
		}); err != nil {
			s.respond(w, r, http.StatusInternalServerError, err)
		}
	}
}
