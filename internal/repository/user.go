package repository

import (
	"fmt"
	"golang.org/x/net/context"
)

type UserRepo interface {
	AddUser(c context.Context, idUser string) error
}

func (r repo) AddUser(c context.Context, idUser string) error {
	//var id string
	fmt.Println(idUser)
	const query = "INSERT INTO \"user\" (id_user) VALUES ($1) ON CONFLICT (id_user) DO NOTHING;"
	//query := "INSERT INTO \"user\" (id_user) SELECT $1 WHERE NOT EXISTS (SELECT id_user FROM \"user\" WHERE id_user = $2) RETURNING id_user"
	_, err := r.db.Exec(c, query, idUser)
	if err != nil {
		return err
	}
	return nil
}
