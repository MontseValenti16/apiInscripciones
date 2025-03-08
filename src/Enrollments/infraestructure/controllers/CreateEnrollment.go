package controllers

import (
	"apiInscripciones/src/Enrollments/application"
	"apiInscripciones/src/Enrollments/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateEnrollmentController struct {
    useCase *application.CreateEnrollmentUseCase
}

func NewCreateEnrollmentController(useCase *application.CreateEnrollmentUseCase) *CreateEnrollmentController {
    return &CreateEnrollmentController{useCase: useCase}
}

func (ec *EnrollmentController) CreateEnrollment(c *gin.Context) {
	var enrollment entities.Enrollment
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ec.createEnrollmentUseCase.Execute(enrollment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Inscripción creada exitosamente"})
}

