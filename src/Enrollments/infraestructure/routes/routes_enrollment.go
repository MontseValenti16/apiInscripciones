package routes

import (
	"apiInscripciones/src/Enrollments/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterEnrollmentRoutes(router *gin.Engine, enrollmentController *controllers.EnrollmentController) {
    enrollmentGroup := router.Group("/inscripciones")
    {
        enrollmentGroup.POST("/", enrollmentController.CreateEnrollment)
        enrollmentGroup.GET("/", enrollmentController.ViewAllEnrollments)  // 🚀 Ver todas
        enrollmentGroup.PUT("/:id", enrollmentController.UpdateEnrollment)
        enrollmentGroup.DELETE("/:id", enrollmentController.DeleteEnrollment)
    }
}
