package model

import (
	"context"
	"log"

	"fish/internal/pkg/dao"
	pgx "fish/internal/pkg/dao"

	"fish/internal/pkg/middleware/auth"
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

func FindLogin(login user.Login) (token string, u user.User) {

	DB := dao.Open()
	sql := `SELECT * FROM public.users WHERE username=? AND password=?`

	if err := DB.Raw(sql, login.Username, login.Password).First(&u).Error; err != nil {
		log.Println(err)
	} else {
		// pm := p.PersonMember()
		// log.Println(pm.Role)
		// token, err := auth.GenerToken(p.Username, pm.Role)
		token, err = auth.GenerToken(u.Username, u.Role)
		if err != nil {
			log.Println(err)
		}

	}
	return token, u
}
