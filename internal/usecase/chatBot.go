package usecase

import (
	"fmt"
	"github.com/rizkyrsyd28/internal/entity"
	"golang.org/x/net/context"
)

type ChatBotUseCase interface {
	GenerateAnswer(c context.Context, idTitle string, idUser string, input entity.UserInput) (entity.BotOutput, error)
}

func (u usecase) GenerateAnswer(c context.Context, idTitle string, idUser string, input entity.UserInput) (entity.BotOutput, error) {
	var result entity.BotOutput = entity.BotOutput{Message: "def", IDTitle: "def"}
	title, err := u.repo.GetHistoryById(c, idTitle, idUser)
	fmt.Println(title.IDTitle + " " + title.Title + " " + title.IDUser)
	if title.Title != "" {
		fmt.Println("True")
	} else {
		u.repo.AddHistory(c, idTitle, "New Chat", idUser)
		fmt.Println("False")
	}
	datas, err := u.repo.GetAllData(c)
	for _, data := range datas {
		fmt.Printf("id_qna : %d, q : %s, a : %s\n", data.IDQna, data.Question, data.Answer)
	}
	result.IDTitle = idTitle
	err = u.repo.AddData(c, "Ini pertanyaan", "Ini Jawaban")
	return result, err
}