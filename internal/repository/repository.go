package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo interface {
	HistoryRepo
	ChatHistoryRepo
	ChatBotRepo
	UserRepo
}

type repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) repo {
	return repo{db: db}
}
