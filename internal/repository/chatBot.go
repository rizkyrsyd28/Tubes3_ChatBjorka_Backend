package repository

import (
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/rizkyrsyd28/internal/entity"
	"golang.org/x/net/context"
)

type ChatBotRepo interface {
	GetAllData(c context.Context) ([]entity.Qna, error)
	AddData(c context.Context, quest string, ans string) error
	DeleteDataById(c context.Context, idQna int) error
}

func (r repo) GetAllData(c context.Context) ([]entity.Qna, error) {
	result := make([]entity.Qna, 0)
	const query = "SELECT * FROM qna"
	err := pgxscan.Select(c, r.db, &result, query)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r repo) AddData(c context.Context, quest string, ans string) error {
	//var id int
	const query = "INSERT INTO qna (question, answer) VALUES ($1, $2)"
	_, err := r.db.Exec(c, query, quest, ans)
	if err != nil {
		return err
	}
	return nil
}

func (r repo) DeleteDataById(c context.Context, idQna int) error {
	const query = "DELETE FROM qna WHERE id_qna=$1"
	_, err := r.db.Exec(c, query, idQna)
	if err != nil {
		return err
	}
	return nil
}
