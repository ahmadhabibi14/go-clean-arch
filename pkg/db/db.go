package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ConfigDB struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func Connect(cfg ConfigDB) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	db, err := sqlx.Connect("postgres", dsn)
	return db, err
}
