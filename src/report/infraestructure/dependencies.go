package infraestructure

import (
	"log"
	"os"
	"report/src/report/application"
	"report/src/report/infraestructure/adapters"
	"report/src/report/infraestructure/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitReports(db *MySQL, router *gin.Engine) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	println("REPORTS")

	// Instanciar el adaptador de RabbitMQ
	amqpURL := os.Getenv("AMQP_URL")
	queueName := os.Getenv("QUEUE_NAME")

	rabbitAdapter, err := adapters.NewRabbitAdapter(amqpURL, queueName)
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}

	// Instanciar casos de uso pasando también el adaptador de RabbitMQ
	createReport := application.NewCreateReportUseCase(db)
	viewReports := application.NewViewReports(db)
	viewOneReport := application.NewViewOneReportUseCase(db)
	sendMessage := application.NewSendMessageUseCase(rabbitAdapter)

	// Instanciar controladores
	sendMessageController := controllers.NewSendMessageController(sendMessage)
	createReportController := controllers.NewCreateReportController(createReport)
	viewReportsController := controllers.NewViewReportsController(viewReports)
	viewOneReportController := controllers.NewViewOneReportController(viewOneReport)

	// Configurar rutas
	SetupReportRoutes(router, sendMessageController, createReportController, viewReportsController, viewOneReportController)
}
