package controllers

import (
	"apiInscripciones/src/Enrollments/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewEnrollmentController struct {
    useCase *application.ViewEnrollmentUseCase
}

func NewViewEnrollmentController(useCase *application.ViewEnrollmentUseCase) *ViewEnrollmentController {
    return &ViewEnrollmentController{useCase: useCase}
}

func (ec *EnrollmentController) ViewAllEnrollments(c *gin.Context) {
	enrollments, err := ec.viewEnrollmentUseCase.GetAllEnrollments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enrollments)
}
