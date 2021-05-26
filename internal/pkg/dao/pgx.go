package dao

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var Pool *pgxpool.Pool

func New() error {
	if Pool != nil {
		return nil
	}

	port := os.Getenv("PG_PORT")
	host := os.Getenv("PG_HOST")
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DATABASE")
	sslmode := os.Getenv("PG_SSLMODE")

	connection := "postgres://" + user + ":" + pass + "@" + host + ":" + port + "/" + dbname + "?" + sslmode

	var env string
	var err error

	if env = os.Getenv("APP_ENV"); len(env) == 0 {
		env = "dev"
	}

	if Pool, err = pgxpool.Connect(context.Background(), connection); err != nil {
		return err
	}

	return nil
}
