package infraestructure

import (
	"report/src/report/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupReportRoutes(
	router *gin.Engine,
	sendMessageController *controllers.SendMessageController,
	createReportController *controllers.CreateReportController,
	viewReportsController *controllers.ViewReportsController,
	viewOneReportController *controllers.ViewOneReportController,
) {
	reportGroup := router.Group("/reports")
	{
		reportGroup.POST("", createReportController.Run)
		reportGroup.POST("/msg", sendMessageController.Run)
		reportGroup.GET("", viewReportsController.Run)
		reportGroup.GET("/{id}", viewOneReportController.Run)
	}
}
