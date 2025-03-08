package controllers

import (
	"apiInscripciones/src/Enrollments/application"

)

type EnrollmentController struct {
	createEnrollmentUseCase *application.CreateEnrollmentUseCase
	viewEnrollmentUseCase   *application.ViewEnrollmentUseCase
	updateEnrollmentUseCase *application.UpdateEnrollmentUseCase
	deleteEnrollmentUseCase *application.DeleteEnrollmentUseCase
}

// Constructor
func NewEnrollmentController(
	createUC *application.CreateEnrollmentUseCase,
	viewUC *application.ViewEnrollmentUseCase,
	updateUC *application.UpdateEnrollmentUseCase,
	deleteUC *application.DeleteEnrollmentUseCase,
) *EnrollmentController {
	return &EnrollmentController{
		createEnrollmentUseCase: createUC,
		viewEnrollmentUseCase:   viewUC,
		updateEnrollmentUseCase: updateUC,
		deleteEnrollmentUseCase: deleteUC,
	}
}