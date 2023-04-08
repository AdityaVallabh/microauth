package main

import (
	"log"
	"microauth/auth"
	"microauth/internal/server"
	"microauth/internal/storage"
	"microauth/pkg/crypto/cipher"
	"microauth/pkg/crypto/hasher"
	"microauth/pkg/token"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// db := storage.NewInMem()
	db := storage.NewPostgres()
	s := &server.Server{
		Router: mux.NewRouter(),
		DB:     db,
		Auth: auth.Auther{
			TokenManager: &token.SimpleTokenManager{
				Cipher:   cipher.HexCipher{},
				Duration: 2 * time.Minute,
			},
			Hasher:  hasher.Sha2{},
			Storage: db,
		},
	}
	s.Setup()
	http.HandleFunc("/", s.ServeHTTP)
	log.Println("going to listen")
	http.ListenAndServe("localhost:8000", nil)
}
