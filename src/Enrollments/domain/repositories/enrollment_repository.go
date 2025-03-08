package repositories

import "apiInscripciones/src/Enrollments/domain/entities"

type EnrollmentRepository interface {
    Save(enrollment entities.Enrollment) error
    Update(id int, enrollment entities.Enrollment) error
    Delete(id int) error
	GetAll() ([]entities.Enrollment, error)
}
