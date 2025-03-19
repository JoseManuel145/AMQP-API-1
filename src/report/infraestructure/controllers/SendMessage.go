package controllers

import (
	"net/http"
	"report/src/report/application"
	"report/src/report/domain/entities"

	"github.com/gin-gonic/gin"
)

type SendMessageController struct {
	reportMessage *application.SendMessageUseCase
}

func NewSendMessageController(useCase *application.SendMessageUseCase) *SendMessageController {
	return &SendMessageController{
		reportMessage: useCase,
	}
}

func (sm *SendMessageController) Run(c *gin.Context) {
	var report entities.Report
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	err := sm.reportMessage.Execute(report)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Message sent successfully"})
}
