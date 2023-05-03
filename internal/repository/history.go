package repository

import (
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/rizkyrsyd28/internal/entity"
	"golang.org/x/net/context"
)

type HistoryRepo interface {
	GetAllHistory(c context.Context, uuid string) ([]entity.History, error)
	GetHistoryById(c context.Context, id int) (entity.History, error)
	AddHistory(c context.Context, title string) error
	DelHistoryById(c context.Context, id int, uuid string) error
	//RenameHistoryById(c context.Context, id int) error
}

func (r repo) GetAllHistory(c context.Context, uuid string) ([]entity.History, error) {
	result := make([]entity.History, 0)
	const query = "SELECT * FROM title_history WHERE uuid = $1"
	err := pgxscan.Select(c, r.db, &result, query, uuid)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r repo) GetHistoryById(c context.Context, id int) (result entity.History, err error) {
	const query = "SELECT * FROM title_history WHERE id=$1"
	err = pgxscan.Select(c, r.db, &result, query, id)
	return result, err
}

func (r repo) AddHistory(c context.Context, title string) error {
	var id int
	const query = "INSERT INTO title_history (title, uuid) VALUES ($1, $2) RETURNING id_title"
	err := r.db.QueryRow(c, query, title).Scan(&id)
	if err != nil {
		fmt.Printf("AddRepo, Result : %s", err.Error())
		return err
	}
	return nil
}

func (r repo) DelHistoryById(c context.Context, id int, uuid string) error {
	fmt.Println("Masuk Repo History")
	const query = "DELETE FROM title_history WHERE id_title=$1 AND uuid=$2"
	_, err := r.db.Exec(c, query, id, uuid)
	if err != nil {
		return err
	}
	return nil
}

//func (r repo) RenameHistoryById(c context.Context, id int, newTitle string) error {
//	const query = ""
//	return nil
//}
