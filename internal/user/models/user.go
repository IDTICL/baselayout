package models

import (
	"context"
	pgx "idticl.app/internal/pkg/dao"
	"idticl.app/internal/pkg/structure/users"
)

func Insert(c context.Context, user users.User) error {
	pgx.Pool.Query()

	return nil
}
