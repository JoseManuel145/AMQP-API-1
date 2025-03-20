package controllers

import (
	"net/http"
	"report/src/report/application"
	"report/src/report/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateReportController struct {
	reportSaver *application.CreateReportUseCase
}

func NewCreateReportController(useCase *application.CreateReportUseCase) *CreateReportController {
	return &CreateReportController{
		reportSaver: useCase,
	}
}
func (cr *CreateReportController) Run(c *gin.Context) {
	var report entities.Report
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	err := cr.reportSaver.Execute(report)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create report: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Report created successfully",
		"report":  report,
	})
}
