package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkyrsyd28/internal/usecase"
	"net/http"
)

func GetChatHistory(uc usecase.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		idTitle := c.Param("id_title")
		data, err := uc.GetChatById(c.Copy().Request.Context(), idTitle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"chatData": data})
	}
}
