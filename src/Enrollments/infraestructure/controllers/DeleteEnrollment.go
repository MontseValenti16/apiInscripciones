package controllers

import (
	"apiInscripciones/src/Enrollments/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteEnrollmentController struct {
	useCase *application.DeleteEnrollmentUseCase
}

func NewDeleteEnrollmentController(useCase *application.DeleteEnrollmentUseCase) *DeleteEnrollmentController {
	return &DeleteEnrollmentController{useCase: useCase}
}

func (ec *EnrollmentController) DeleteEnrollment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) // 🔹 Convertir ID a entero
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = ec.deleteEnrollmentUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inscripción eliminada correctamente"})
}
