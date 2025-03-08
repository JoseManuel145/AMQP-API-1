package controllers

import (
	"net/http"
	"report/src/report/application/usecases"

	"github.com/gin-gonic/gin"
)

type ViewReportsController struct {
	UseCase *usecases.ViewReportsUseCase
}

func NewViewReportsController(useCase *usecases.ViewReportsUseCase) *ViewReportsController {
	return &ViewReportsController{UseCase: useCase}
}

func (c *ViewReportsController) Run(ctx *gin.Context) {
	reports, err := c.UseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if reports == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No reports found"})
		return
	}
	ctx.JSON(http.StatusOK, reports)
}
