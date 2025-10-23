package controllers

import (
	"kids-note-api/models"
	"kids-note-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendUserEmail(c *gin.Context) {
	var input models.Mail
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": 4, "Data": err.Error()})
	}
	// 登録があるメールアドレスかどうかをチェック
	existEmailCount := services.CheckExistUserEmail(input.Email)
	if existEmailCount != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "Mail Error"})
		return
	} else {
		token, err := services.CreateMailToken(input.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "Token Error"})
			return
		}
		mailSendStatus := services.SendUserEmail(input.Email, token, input.NewUser)
		if !mailSendStatus {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": 5, "Data": "Mail Failed"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"Status": 0, "Data": "Mail Sent"})
}
