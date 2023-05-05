package repository

import (
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/rizkyrsyd28/internal/entity"
	"golang.org/x/net/context"
)

type HistoryRepo interface {
	GetAllHistory(c context.Context, idUser string) ([]entity.History, error)
	GetHistoryById(c context.Context, idTitle string, idUser string) (entity.History, error)
	AddHistory(c context.Context, idTitle string, title string, idUser string) error
	DelHistoryById(c context.Context, idTitle string) error
	SetTitleById(c context.Context, idTitle string, newTitle string) error
}

func (r repo) GetAllHistory(c context.Context, idUser string) ([]entity.History, error) {
	result := make([]entity.History, 0)
	const query = "SELECT * FROM title_history WHERE id_user = $1"
	err := pgxscan.Select(c, r.db, &result, query, idUser)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r repo) GetHistoryById(c context.Context, idTitle string, idUser string) (result entity.History, err error) {
	const query = "SELECT * FROM title_history WHERE id_title=$1 AND id_user=$2"
	err = r.db.QueryRow(c, query, idTitle, idUser).Scan(&result.IDTitle, &result.Title, &result.IDUser)
	return result, err
}

func (r repo) AddHistory(c context.Context, idTitle string, title string, idUser string) error {
	var id string
	const query = "INSERT INTO title_history (id_title, title, id_user) VALUES ($1, $2, $3) RETURNING id_title"
	err := r.db.QueryRow(c, query, idTitle, title, idUser).Scan(&id)
	if err != nil {
		fmt.Printf("AddRepo, Result : %s", err.Error())
		return err
	}
	return nil
}

func (r repo) DelHistoryById(c context.Context, idTitle string) error {
	fmt.Println("Masuk Repo History")
	const query = "DELETE FROM title_history WHERE id_title=$1"
	_, err := r.db.Exec(c, query, idTitle)
	if err != nil {
		return err
	}
	return nil
}

func (r repo) SetTitleById(c context.Context, idTitle string, newTitle string) error {
	const query = "UPDATE title_history SET title=$1 WHERE id_title=$2"
	_, err := r.db.Exec(c, query, newTitle, idTitle)
	if err != nil {
		return err
	}
	return nil
}
