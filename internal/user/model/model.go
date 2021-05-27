package model

import (
	"context"

	pgx "fish/internal/pkg/dao"
	"fish/internal/pkg/structure/user"
)

func Insert(c context.Context, user user.User) error {

	_, err := pgx.Pool.Exec(c, "insert into user('username','password','role') vale($1,$2,$3,$4)",
		user.Username, user.Password, user.Role, user.Age)

	return err
}

func FindAll(c context.Context) error {
	_, err := pgx.Pool.Exec(c, "select * from user")
	return err
}
