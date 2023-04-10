package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/AdityaVallabh/microauth/auth"
	"github.com/AdityaVallabh/microauth/internal/server"
	"github.com/AdityaVallabh/microauth/internal/server/models/users"
	"github.com/AdityaVallabh/microauth/internal/storage"
	"github.com/AdityaVallabh/microauth/pkg/crypto/hasher"
	"github.com/AdityaVallabh/microauth/pkg/token"

	"github.com/gorilla/mux"
)

func main() {
	// db := storage.NewInMem()
	db := storage.NewPostgres()
	users.Init(db)
	db.AutoMigrate(&token.PersistedToken{})
	tm := &token.PersistedTokenManager{
		Rand:     *rand.New(rand.NewSource(time.Now().Unix())),
		Store:    db,
		Duration: time.Second * 30,
	}
	tm.RunCleaner(context.Background().Done())
	s := &server.Server{
		Router: mux.NewRouter(),
		DB:     db,
		Auth: auth.Auther{
			// TokenManager: &token.SimpleTokenManager{
			// 	Cipher:   cipher.HexCipher{},
			// 	Duration: 2 * time.Minute,
			// },
			TokenManager: tm,
			Hasher:       hasher.Sha2{},
			Storage:      users.Manager,
		},
	}
	s.Setup()
	http.HandleFunc("/", s.ServeHTTP)
	log.Println("going to listen")
	http.ListenAndServe("localhost:8000", nil)
}
