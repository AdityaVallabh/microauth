package server

import (
	"log"
	"microauth/internal/server/models/users"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	templatesDir = "internal/server/templates"
)

type Auther interface {
	SavePassword(string, string) error
	GenerateToken(string) (string, error)
	Validate(string) (string, bool)
	CheckPassword(string, string) bool
	GetExpiry(string) time.Time
}

type Storage interface {
	AutoMigrate(...any) error
	Find(string, any) error
}

type Server struct {
	Router *mux.Router
	DB     Storage
	Auth   Auther
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) Setup() {
	if err := s.DB.AutoMigrate(users.User{}); err != nil {
		log.Printf("error migrating: %s\n", err.Error())
	}
	log.Println("migrations complete")
	s.routes()
}
