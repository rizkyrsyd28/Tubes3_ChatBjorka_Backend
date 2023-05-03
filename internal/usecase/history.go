package usecase

import (
	"fmt"
	"github.com/rizkyrsyd28/internal/entity"
	"golang.org/x/net/context"
)

type HistoryUseCase interface {
	GetAllHistory(c context.Context, uuid string) ([]entity.History, error)
	GetHistoryById(c context.Context, id int) (entity.History, error)
	AddHistory(c context.Context, title string) error
	DelHistoryById(c context.Context, id int, uuid string) error
}

func (u usecase) GetAllHistory(c context.Context, uuid string) ([]entity.History, error) {
	return u.repo.GetAllHistory(c, uuid)
}
func (u usecase) GetHistoryById(c context.Context, id int) (entity.History, error) {
	return u.repo.GetHistoryById(c, id)
}
func (u usecase) AddHistory(c context.Context, title string) error {
	fmt.Println("AddUC")
	return u.repo.AddHistory(c, title)
}
func (u usecase) DelHistoryById(c context.Context, id int, uuid string) error {
	err := u.repo.DelChatById(c, id)
	fmt.Println("Masuk UC")
	if err != nil {
		return err
	}
	return u.repo.DelHistoryById(c, id, uuid)
}
