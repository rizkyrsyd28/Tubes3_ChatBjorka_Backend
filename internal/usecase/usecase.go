package usecase

import "github.com/rizkyrsyd28/internal/repository"

type UseCase interface {
	HistoryUseCase
	ChatHistoryUseCase
	ChatBotUseCase
}

type usecase struct {
	repo repository.Repo
}

func NewUseCase(repo repository.Repo) usecase {
	return usecase{repo: repo}
}
