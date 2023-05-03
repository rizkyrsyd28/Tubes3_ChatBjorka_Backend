package route

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/rizkyrsyd28/internal/handler"
	"github.com/rizkyrsyd28/internal/repository"
	"github.com/rizkyrsyd28/internal/usecase"
	"os"
)

func Routes(r *gin.Engine) {
	// Connect DB
	conn, err := pgx.Connect(context.Background(), "postgres://stima3_admin@stima3-chat:Rizkyrasy.id28@stima3-chat.postgres.database.azure.com:5432/stima3?sslmode=require")
	if err != nil {
		fmt.Fprint(os.Stderr, "Unable to Connect to Database : %v\n", err)
		os.Exit(1)
	}

	repo := repository.NewRepo(conn)
	ucase := usecase.NewUseCase(repo)

	r.GET("/history/:uuid", handler.GetAllHistory(ucase))
	r.DELETE("/history/:uuid", handler.DeleteHistory(ucase))
	r.GET("/chat_history/:uuid/:id", handler.GetChatHistory(ucase))
	//r.GET("/:id/bot_respond", handler.GetBotRespond())
	//r.POST("/:id/user_respond", handler.PostBotPrompt())
}
