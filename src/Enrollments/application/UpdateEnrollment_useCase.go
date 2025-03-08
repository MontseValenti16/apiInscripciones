package application

import (
    "apiInscripciones/src/Enrollments/domain/entities"
    "apiInscripciones/src/Enrollments/domain/repositories"
)

type UpdateEnrollmentUseCase struct {
    repo repositories.EnrollmentRepository
}

func NewUpdateEnrollmentUseCase(repo repositories.EnrollmentRepository) *UpdateEnrollmentUseCase {
    return &UpdateEnrollmentUseCase{repo: repo}
}

func (uc *UpdateEnrollmentUseCase) Execute(id int, enrollment entities.Enrollment) error {
    return uc.repo.Update(id, enrollment)
}
