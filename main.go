package main

import (
	"github.com/gin-gonic/gin"
	"tubes3-chatbjorka-backend/internal/route"
)

func main() {
	r := gin.Default()

	route.Routes(r)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
