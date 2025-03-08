package controllers

import (
	"net/http"
	"strconv"

	"report/src/report/application/usecases"

	"github.com/gin-gonic/gin"
)

type ViewOneReportController struct {
	useCase *usecases.ViewOneReportUseCase
}

func NewViewOneReportController(uc *usecases.ViewOneReportUseCase) *ViewOneReportController {
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
