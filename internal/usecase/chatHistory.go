package usecase

import (
	"github.com/rizkyrsyd28/internal/entity"
	"golang.org/x/net/context"
)

type ChatHistoryUseCase interface {
	GetChatById(c context.Context, idTitle string) ([]entity.ChatHistoryOutput, error)
	DelChatById(c context.Context, idTitle string) error
	AddChat(c context.Context, idTitle string, userChat string, botChat string) error
}

func (u usecase) GetChatById(c context.Context, idTitle string) ([]entity.ChatHistoryOutput, error) {
	result := make([]entity.ChatHistoryOutput, 0)
	data, err := u.repo.GetChatById(c, idTitle)
	if err != nil {
		return nil, err
	}
	if len(data.Chat) != 0 {
		for _, item := range data.Chat {
			result = append(result, entity.ChatHistoryOutput{Bot: item.BotChat, User: item.UserChat})
		}
	}
	return result, nil
}
func (u usecase) DelChatById(c context.Context, idTitle string) error {
	return u.repo.DelChatById(c, idTitle)
}
func (u usecase) AddChat(c context.Context, idTitle string, userChat string, botChat string) error {
	return u.repo.AddChat(c, idTitle, userChat, botChat)
}
