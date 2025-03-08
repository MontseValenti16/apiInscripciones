package application

import (
    "apiInscripciones/src/Enrollments/domain/entities"
    "apiInscripciones/src/Enrollments/domain/repositories"
)

type CreateEnrollmentUseCase struct {
    repo repositories.EnrollmentRepository
}

func NewCreateEnrollmentUseCase(repo repositories.EnrollmentRepository) *CreateEnrollmentUseCase {
    return &CreateEnrollmentUseCase{repo: repo}
}

func (uc *CreateEnrollmentUseCase) Execute(enrollment entities.Enrollment) error {
    return uc.repo.Save(enrollment)
}
