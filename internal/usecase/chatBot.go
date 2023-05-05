package usecase

import (
	"fmt"
	"github.com/rizkyrsyd28/internal/algorithms"
	"github.com/rizkyrsyd28/internal/entity"
	"golang.org/x/net/context"
)

type ChatBotUseCase interface {
	GenerateAnswer(c context.Context, idTitle string, idUser string, input entity.UserInput) (entity.BotOutput, error)
}

func (u usecase) GenerateAnswer(c context.Context, idTitle string, idUser string, input entity.UserInput) (entity.BotOutput, error) {
	title, err := u.repo.GetHistoryById(c, idTitle, idUser)
	if title.Title != "" {
		fmt.Println("True")
	} else {
		u.repo.AddHistory(c, idTitle, "New Chat", idUser)
		fmt.Println("False")
	}
	hasil := algorithms.HandleQueries(u.repo, c, input.Message, input.Algo)
	result := entity.BotOutput{
		Message: hasil,
		IDTitle: idTitle,
	}
	u.repo.AddChat(c, idTitle, input.Message, result.Message)
	return result, err
}
