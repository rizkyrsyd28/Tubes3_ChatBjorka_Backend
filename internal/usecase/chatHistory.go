package usecase

import (
	"github.com/rizkyrsyd28/internal/entity"
	"golang.org/x/net/context"
)

type ChatHistoryUseCase interface {
	GetChatById(c context.Context, idTitle int) (entity.ChatHistory, error)
	DelChatById(c context.Context, idTitle int) error
	AddChat(c context.Context, idTitle int, userChat string, botChat string) error
}

func (u usecase) GetChatById(c context.Context, idTitle int) (entity.ChatHistory, error) {
	return u.repo.GetChatById(c, idTitle)
}
func (u usecase) DelChatById(c context.Context, idTitle int) error {
	return u.repo.DelChatById(c, idTitle)
}
func (u usecase) AddChat(c context.Context, idTitle int, userChat string, botChat string) error {
	return u.repo.AddChat(c, idTitle, userChat, botChat)
}
