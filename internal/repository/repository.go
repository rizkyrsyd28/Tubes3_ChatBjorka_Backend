package repository

import "github.com/jackc/pgx/v5"

type Repo interface {
	HistoryRepo
	ChatHistoryRepo
	ChatBotRepo
	UserRepo
}

type repo struct {
	db *pgx.Conn
}

func NewRepo(db *pgx.Conn) repo {
	return repo{db: db}
}
