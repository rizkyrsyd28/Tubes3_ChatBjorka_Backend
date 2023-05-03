package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkyrsyd28/internal/usecase"
	"net/http"
	"strconv"
)

func GetChatHistory(uc usecase.ChatHistoryUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		data, err := uc.GetChatById(c.Copy().Request.Context(), int(id))

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"ChatHistory": data})
	}
}

func GetBotRespond() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := "Jawaban Bot UserId : Unknown"
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		if id == 1 {
			data = "Jawaban Bot UserId : 1"
		}
		if id == 2 {
			data = "Jawaban Bot UserId : 2"
		}

		c.JSON(http.StatusOK, data)
	}
}

type Prompt struct {
	Data1 string `json:"data1"`
	Data2 int    `json:"data2"`
}

func PostBotPrompt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var postData Prompt

		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		err := c.BindJSON(&postData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if id == 1 {
			postData.Data2 += 10
			postData.Data1 += " New 1 "
		}
		if id == 2 {
			postData.Data2 += 200
			postData.Data1 += " New 2 "
		}

		c.JSON(http.StatusOK, postData)
	}
}
