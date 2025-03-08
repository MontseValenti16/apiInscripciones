package application

import "apiInscripciones/src/Enrollments/domain/repositories"

type DeleteEnrollmentUseCase struct {
    repo repositories.EnrollmentRepository
}

func NewDeleteEnrollmentUseCase(repo repositories.EnrollmentRepository) *DeleteEnrollmentUseCase {
    return &DeleteEnrollmentUseCase{repo: repo}
}

func (uc *DeleteEnrollmentUseCase) Execute(id int) error {
    return uc.repo.Delete(id)
}
