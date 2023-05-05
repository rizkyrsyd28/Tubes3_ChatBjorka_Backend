package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkyrsyd28/internal/entity"
	"github.com/rizkyrsyd28/internal/usecase"
	"net/http"
)

func PostUserRespond(uc usecase.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input entity.UserInput

		idUser := c.Param("id_user")
		idTitle := c.Param("id_title")

		err := c.ShouldBindJSON(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, err := uc.GenerateAnswer(c.Copy().Request.Context(), idTitle, idUser, input)

		c.JSON(http.StatusOK, data)
	}
}
