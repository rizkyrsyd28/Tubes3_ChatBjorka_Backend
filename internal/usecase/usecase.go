package usecase

import "tubes3-chatbjorka-backend/internal/repository"

type UseCase interface {
	HistoryUseCase
	ChatHistoryUseCase
}

type usecase struct {
	repo repository.Repo
}

func NewUseCase(repo repository.Repo) usecase {
	return usecase{repo: repo}
}
