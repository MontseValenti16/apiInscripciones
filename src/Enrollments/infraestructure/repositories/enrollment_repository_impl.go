package repositories

import (
	"apiInscripciones/src/Enrollments/domain/entities"
	"database/sql"
	"apiInscripciones/src/Enrollments/domain/repositories"
)


type EnrollmentRepositoryImpl struct {
	db *sql.DB
}

// Aseguramos que cumple con la interfaz en tiempo de compilación
var _ repositories.EnrollmentRepository = (*EnrollmentRepositoryImpl)(nil)

func NewEnrollmentRepositoryImpl(db *sql.DB) repositories.EnrollmentRepository {
	return &EnrollmentRepositoryImpl{db: db}
}

// func NewEnrollmentRepositoryImpl(db *sql.DB) *EnrollmentRepositoryImpl {
// 	return &EnrollmentRepositoryImpl{db: db}
// }

// Guardar una nueva inscripción
// func (repo *EnrollmentRepositoryImpl) Save(enrollment entities.Enrollment) (int64, error) {
// 	query := "INSERT INTO inscripciones (estudiante_id, materia_id, estado, creado_en) VALUES (?, ?, ?, NOW())"
// 	result, err := repo.db.Exec(query, enrollment.EstudianteID, enrollment.MateriaID, enrollment.Estado)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return result.LastInsertId()
// }
func (repo *EnrollmentRepositoryImpl) Save(enrollment entities.Enrollment) (int64, error) {
	query := "INSERT INTO inscripciones (estudiante_id, materia_id, estado, creado_en) VALUES (?, ?, ?, NOW())"
	result, err := repo.db.Exec(query, enrollment.EstudianteID, enrollment.MateriaID, enrollment.Estado)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}



// Obtener inscripción por ID
func (repo *EnrollmentRepositoryImpl) GetByID(id int) (*entities.Enrollment, error) {
	query := "SELECT id, estudiante_id, materia_id, estado, creado_en FROM inscripciones WHERE id = ?"
	row := repo.db.QueryRow(query, id)

	var enrollment entities.Enrollment
	if err := row.Scan(&enrollment.ID, &enrollment.EstudianteID, &enrollment.MateriaID, &enrollment.Estado, &enrollment.CreadoEn); err != nil {
		return nil, err
	}
	return &enrollment, nil
}

// Actualizar inscripción
func (repo *EnrollmentRepositoryImpl) Update(id int, enrollment entities.Enrollment) error {
	query := "UPDATE inscripciones SET estudiante_id = ?, materia_id = ?, estado = ? WHERE id = ?"
	_, err := repo.db.Exec(query, enrollment.EstudianteID, enrollment.MateriaID, enrollment.Estado, id)
	return err
}

// Eliminar inscripción
func (repo *EnrollmentRepositoryImpl) Delete(id int) error {
	query := "DELETE FROM inscripciones WHERE id = ?"
	_, err := repo.db.Exec(query, id)
	return err
}

// Obtener todas las inscripciones
func (repo *EnrollmentRepositoryImpl) GetAll() ([]entities.Enrollment, error) {
	rows, err := repo.db.Query("SELECT id, estudiante_id, materia_id, estado, creado_en FROM inscripciones")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []entities.Enrollment
	for rows.Next() {
		var enrollment entities.Enrollment
		if err := rows.Scan(&enrollment.ID, &enrollment.EstudianteID, &enrollment.MateriaID, &enrollment.Estado, &enrollment.CreadoEn); err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}

	return enrollments, nil
}