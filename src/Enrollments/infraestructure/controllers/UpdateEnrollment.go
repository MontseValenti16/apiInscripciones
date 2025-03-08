package controllers

import (
	"apiInscripciones/src/Enrollments/application"
	"apiInscripciones/src/Enrollments/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateEnrollmentController struct {
    useCase *application.UpdateEnrollmentUseCase
}

func NewUpdateEnrollmentController(useCase *application.UpdateEnrollmentUseCase) *UpdateEnrollmentController {
    return &UpdateEnrollmentController{useCase: useCase}
}

func (ec *EnrollmentController) UpdateEnrollment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) // Convertir id a int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var enrollment entities.Enrollment
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ec.updateEnrollmentUseCase.Execute(id, enrollment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inscripción actualizada correctamente"})
}