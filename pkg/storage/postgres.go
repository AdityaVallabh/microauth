package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	DB       *gorm.DB
}

func (p *Postgres) Connect() (err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		p.Host, p.User, p.Password, p.DBName, p.Port, "disable", "Asia/Kolkata")
	p.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}
