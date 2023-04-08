package storage

import (
	"log"
	"microauth/internal/server/models/users"
	"microauth/pkg/storage"
)

type Postgres struct {
	*storage.Postgres
}

func (p *Postgres) AutoMigrate(v ...any) error {
	return p.DB.AutoMigrate(v...)
}

func (p *Postgres) Find(key string, v any) error {
	return p.DB.First(v, "email = ?", key).Error
}

func (p *Postgres) Save(v any) error {
	return p.DB.Save(v).Error
}

func (p *Postgres) PasswordHashByEmail(email string) (string, error) {
	var user users.User
	err := p.Find(email, &user)
	return user.PasswordHash, err
}

func (p *Postgres) SetPasswordHashByEmail(email, phash string) error {
	user := users.User{
		Email:        email,
		PasswordHash: phash,
	}
	return p.DB.Save(&user).Error
}

func NewPostgres() *Postgres {
	p := storage.Postgres{
		Host:     "localhost",
		User:     "postgres",
		Password: "postgres",
		DBName:   "postgres",
		Port:     5432,
	}
	if err := p.Connect(); err != nil {
		log.Printf("could not create connection to postgres: %s\n", err.Error())
		return nil
	}
	return &Postgres{&p}
}
