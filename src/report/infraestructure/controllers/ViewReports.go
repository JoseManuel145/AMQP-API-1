package controllers

import (
	"fmt"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error diferente de nil": err.Error()})
		return
	}
	if reports == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No reports found"})
		return
	}
	// Depuración
	for _, report := range reports {
		fmt.Printf("Report: %#v\n", report) // Imprime los datos para depurar
	}
	ctx.JSON(http.StatusOK, reports)
}
