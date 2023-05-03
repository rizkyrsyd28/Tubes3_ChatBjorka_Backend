package route

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"os"
	"tubes3-chatbjorka-backend/internal/handler"
	"tubes3-chatbjorka-backend/internal/repository"
	"tubes3-chatbjorka-backend/internal/usecase"
)

func Routes(r *gin.Engine) {
	// Connect DB
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:ras28@localhost:5432/tubes3_try")
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
