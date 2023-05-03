package repository

import (
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/rizkyrsyd28/internal/entity"
	"golang.org/x/net/context"
)

type ChatHistoryRepo interface {
	GetChatById(c context.Context, idTitle int) (entity.ChatHistory, error)
	DelChatById(c context.Context, idTitle int) error
	AddChat(c context.Context, idTitle int, userChat string, botChat string) error
}

func (r repo) GetChatById(c context.Context, idTitle int) (chat entity.ChatHistory, err error) {
	chatCont := make([]entity.ChatContent, 0)
	rawChat := make([]struct {
		IdChat   int    `json:"id_chat"`
		IdTitle  int    `json:"id_title"`
		UserChat string `json:"user_chat"`
		BotChat  string `json:"bot_chat"`
	}, 0)

	const query = "SELECT id_chat, id_title, user_chat, bot_chat FROM chat_history WHERE id_title=$1 ORDER BY id_chat"
	err = pgxscan.Select(c, r.db, &rawChat, query, idTitle)
	for _, ch := range rawChat {
		chatCont = append(chatCont, entity.ChatContent{IDChat: ch.IdChat, UserChat: ch.UserChat, BotChat: ch.BotChat})
	}
	chat.IDTitle = idTitle
	chat.Chat = chatCont
	return chat, err
}

func (r repo) DelChatById(c context.Context, idTitle int) error {
	fmt.Println("Masuk Repo Chat")
	const query = "DELETE FROM chat_history WHERE id_title=$1"
	_, err := r.db.Exec(c, query, idTitle)
	if err != nil {
		return err
	}
	return nil
}

func (r repo) AddChat(c context.Context, idTitle int, userChat string, botChat string) error {
	var id int
	const query = "INSERT INTO chat_history (id_title, user_chat, bot_chat) VALUES ($1, $2, $3) RETURNING id_chat"
	err := r.db.QueryRow(c, query, idTitle, userChat, botChat).Scan(&id)
	if err != nil {
		fmt.Printf("AddRepo, Result : %s", err.Error())
		return err
	}
	return nil
}
