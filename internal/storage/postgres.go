package storage

import (
	"log"

	"github.com/AdityaVallabh/microauth/pkg/storage"
)

type Postgres struct {
	*storage.Postgres
}

func (p *Postgres) AutoMigrate(v ...any) error {
	return p.DB.AutoMigrate(v...)
}

func (p *Postgres) Find(v any, keyName, keyValue string) error {
	return p.DB.First(v, keyName+" = ?", keyValue).Error
}

func (p *Postgres) FindAll(v any) error {
	return p.DB.Find(v).Error
}

func (p *Postgres) Save(v any, _ string) error {
	return p.DB.Save(v).Error
}

func (p *Postgres) Delete(v any, _ string) error {
	return p.DB.Delete(v).Error
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
