package controllers

import (
	"encoding/json"
	"net/http"
	"report/src/report/application/usecases"
	"report/src/report/domain/entities"
	"report/src/report/infraestructure/adapters"

	"github.com/gin-gonic/gin"
)

type CreateReportController struct {
	reportSaver *usecases.CreateReportUseCase
	rabbit      *adapters.RabbitAdapter
}

func NewCreateReportController(useCase *usecases.CreateReportUseCase, rabbit *adapters.RabbitAdapter) *CreateReportController {
	return &CreateReportController{
		reportSaver: useCase,
		rabbit:      rabbit,
	}
}

func (cr *CreateReportController) Run(c *gin.Context) {
	var report entities.Report
	// Validar JSON y enlazar a la estructura
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	// Guardar el reporte
	err := cr.reportSaver.Execute(report.ID, report.Title, report.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create report: " + err.Error()})
		return
	}

	// Formatea el mensaje a enviar.
	messageData := struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}{
		ID:      report.ID,
		Title:   report.Title,
		Content: report.Content,
	}

	messageJSON, err := json.Marshal(messageData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize JSON: " + err.Error()})
		return
	}

	// Envía el mensaje utilizando el adaptador de RabbitMQ.
	err = cr.rabbit.SendMessage(string(messageJSON))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message: " + err.Error()})
		return
	}

	// Retornar el reporte creado
	c.JSON(http.StatusCreated, gin.H{"message": "Report saved and message sent successfully"})
}
