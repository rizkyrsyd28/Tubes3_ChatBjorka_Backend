package main

import (
	"github.com/rizkyrsyd28/internal/route"
	"os"

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
