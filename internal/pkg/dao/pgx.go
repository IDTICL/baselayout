package pgx

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

	if env = os.Getenv("ENV"); len(env) == 0 {
		env = "development"
	}

	connection := fmt.Sprintf(pgConfig,
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DATABASE"),
		os.Getenv("PG_SSLMODE"),
	)

	if Pool, err = pgxpool.Connect(context.Background(), connection); err != nil {
		return err
	}

	return nil
}
