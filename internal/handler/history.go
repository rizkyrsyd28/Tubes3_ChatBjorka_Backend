package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rizkyrsyd28/internal/usecase"
	"net/http"
)

func GetAllHistory(uc usecase.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("\nRequestType : %s\n", c.Request.Method)
		//var data []entity.History
		idUser := c.Param("id_user")
		data, err := uc.GetAllHistory(c.Copy().Request.Context(), idUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"historyTitle": data})
	}
}

func DeleteHistory(uc usecase.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {

		idTitle := c.Param("id_title")

		err := uc.DelHistoryById(c.Copy().Request.Context(), idTitle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	}
}

func RenameTitle(uc usecase.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		idTitle := c.Param("id_title")
		var data struct {
			New string `json:"nama"`
		}

		err := c.ShouldBindJSON(&data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		err = uc.SetTitleById(c, idTitle, data.New)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": data})
	}
}
