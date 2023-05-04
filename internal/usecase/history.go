package usecase

import (
	"fmt"
	"github.com/rizkyrsyd28/internal/entity"
	"golang.org/x/net/context"
)

type HistoryUseCase interface {
	GetAllHistory(c context.Context, idUser string) ([]entity.HistoryOutput, error)
	AddHistory(c context.Context, idTitle string, title string, idUser string) error
	DelHistoryById(c context.Context, idTitle string) error
}

func (u usecase) GetAllHistory(c context.Context, idUser string) ([]entity.HistoryOutput, error) {
	result := make([]entity.HistoryOutput, 0)
	data, err := u.repo.GetAllHistory(c, idUser)
	if err != nil {
		return nil, err
	}
	if len(data) != 0 {
		for _, item := range data {
			result = append(result, entity.HistoryOutput{Title: item.Title, IDTitle: item.IDTitle})
		}
	} else {
		fmt.Println("Add User")
		err := u.repo.AddUser(c, idUser)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (u usecase) AddHistory(c context.Context, idTitle string, title string, idUser string) error {
	return u.repo.AddHistory(c, idTitle, title, idUser)
}

func (u usecase) DelHistoryById(c context.Context, idTitle string) error {
	err := u.repo.DelChatById(c, idTitle)
	fmt.Println("Masuk UC")
	if err != nil {
		return err
	}
	return u.repo.DelHistoryById(c, idTitle)
}
