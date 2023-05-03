package main

import (
	"os"
	"tubes3-chatbjorka-backend/internal/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	route.Routes(r)

	port := os.Getenv("HTTP_PLATFORM_PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
