package dao

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var Pool *pgxpool.Pool

func New() error {
	if Pool != nil {
		return nil
	}

	const pgConfig string = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s pool_max_conns=20 pool_min_conns=1 pool_max_conn_lifetime=60s pool_max_conn_idle_time=120s"

	var env string
	var err error

	if env = os.Getenv("APP_ENV"); len(env) == 0 {
		env = "dev"
	}

	connection := fmt.Sprintf(pgConfig,
		os.Getenv("POSTGRESQL_HOST"),
		os.Getenv("POSTGRESQL_PORT"),
		os.Getenv("POSTGRESQL_USER"),
		os.Getenv("POSTGRESQL_PASSWORD"),
		os.Getenv("POSTGRESQL_DATABASE"),
		os.Getenv("POSTGRESQL_SSLMODE"),
	)

	if Pool, err = pgxpool.Connect(context.Background(), connection); err != nil {
		return err
	}

	return nil
}
