package route

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rizkyrsyd28/internal/handler"
	"github.com/rizkyrsyd28/internal/repository"
	"github.com/rizkyrsyd28/internal/usecase"
	"os"
)

func Routes(r *gin.Engine) {
	// Connect DB
	conn, err := pgxpool.New(context.Background(), "DB_URL")
	if err != nil {
		fmt.Fprint(os.Stderr, "Unable to Connect to Database : %v\n", err)
		os.Exit(1)
	}

	repo := repository.NewRepo(conn)
	ucase := usecase.NewUseCase(repo)

	// No 1
	r.GET("/history/:id_user", handler.GetAllHistory(ucase))

	// No 2
	r.GET("/chat_history/:id_title", handler.GetChatHistory(ucase))

	// No 3
	r.POST("/user_respond/:id_user/:id_title", handler.PostUserRespond(ucase))

	// No 5
	r.DELETE("/delete_history/:id_title", handler.DeleteHistory(ucase))

	// No 4
	r.POST("/rename_history/:id_title", handler.RenameTitle(ucase))

}
