package postgres

import (
	"database/sql"
	"fmt"

	"github.com/caarlos0/env/v6"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type config struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
	Password string `env:"DB_PASSWORD"`
	User     string `env:"DB_USER"`
	Name     string `env:"DB_NAME"`
}

type Repository struct {
	db *sql.DB
}

func Connect() (r Repository, err error) {
	cfg := config{}
	if err = env.Parse(&cfg); err != nil {
		return
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	return Repository{db}, nil
}

func (r Repository) Close() {
	if err := r.db.Close(); err != nil {
		zap.L().Fatal("can't close DB connect", zap.Error(err))
	}
}
