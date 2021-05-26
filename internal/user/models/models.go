package models

import (
	"context"

	pgx "idticl.app/internal/pkg/dao"
	"idticl.app/internal/pkg/structure/users"
)

func Insert(c context.Context, user users.User) error {

	_, err := pgx.Pool.Exec(c, "insert into users('username','password','role') vale($1,$2,$3,$4)",
		user.Username, user.Password, user.Role, user.Age)

	return err
}

func FindAll(c context.Context) error {
	_, err := pgx.Pool.Exec(c, "select * from users")
	return err
}
