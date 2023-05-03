package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tubes3-chatbjorka-backend/internal/usecase"
)

func GetAllHistory(uc usecase.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("\nRequestType : %s\n", c.Request.Method)
		//var data []entity.History
		uuid := c.Param("uuid")
		data, err := uc.GetAllHistory(c.Copy().Request.Context(), uuid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"history": data})
	}
}

func DeleteHistory(uc usecase.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data struct {
			Index int `json:"index"`
		}

		uuid := c.Param("uuid")

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("Index %d dihapus\n", data.Index)
		err := uc.DelHistoryById(c.Copy().Request.Context(), data.Index, uuid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	}
}
