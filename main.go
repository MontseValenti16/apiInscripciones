package main

import (
	"apiInscripciones/src/Enrollments/application"
	repoEnrollment "apiInscripciones/src/Enrollments/infraestructure/repositories"
	"apiInscripciones/src/Enrollments/infraestructure/controllers"
	"apiInscripciones/src/Enrollments/infraestructure/routes"
	"apiInscripciones/src/core"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Inicializar base de datos
	core.InitDB()
	defer core.CloseDB()

	// Crear router
	router := gin.Default()

	// Crear repositorio
	enrollmentRepo := repoEnrollment.NewEnrollmentRepositoryImpl(core.DB)

	// Crear casos de uso
	createEnrollmentUseCase := application.NewCreateEnrollmentUseCase(enrollmentRepo)
	viewEnrollmentUseCase := application.NewViewEnrollmentUseCase(enrollmentRepo)
	updateEnrollmentUseCase := application.NewUpdateEnrollmentUseCase(enrollmentRepo)
	deleteEnrollmentUseCase := application.NewDeleteEnrollmentUseCase(enrollmentRepo)

	// Crear controlador
	enrollmentController := controllers.NewEnrollmentController(
		createEnrollmentUseCase,
		viewEnrollmentUseCase,
		updateEnrollmentUseCase,
		deleteEnrollmentUseCase,
	)

	// Registrar rutas
	routes.RegisterEnrollmentRoutes(router, enrollmentController)

	// Ruta de prueba
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}