package application

import (
    "apiInscripciones/src/Enrollments/domain/entities"
    "apiInscripciones/src/Enrollments/domain/repositories"
)

type ViewEnrollmentUseCase struct {
    repo repositories.EnrollmentRepository
}

func NewViewEnrollmentUseCase(repo repositories.EnrollmentRepository) *ViewEnrollmentUseCase {
    return &ViewEnrollmentUseCase{repo: repo}
}

// 🚀 Obtener **todas** las inscripciones
func (uc *ViewEnrollmentUseCase) GetAllEnrollments() ([]entities.Enrollment, error) {
    return uc.repo.GetAll()
}
