package controllers

import (
	"net/http"
	"report/src/report/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ViewOneReportController struct {
	useCase *application.ViewOneReportUseCase
}

func NewViewOneReportController(uc *application.ViewOneReportUseCase) *ViewOneReportController {
	return &ViewOneReportController{useCase: uc}
}

func (c *ViewOneReportController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid report ID"})
		return
	}

	report, err := c.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, report)
}
