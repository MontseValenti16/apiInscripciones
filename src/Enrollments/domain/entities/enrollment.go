package entities

import "time"

type Enrollment struct {
    ID           int       `json:"id" gorm:"primaryKey"`
    EstudianteID int       `json:"estudiante_id" gorm:"column:estudiante_id"`
    MateriaID    int       `json:"materia_id" gorm:"column:materia_id"`
    Estado       string    `json:"estado" gorm:"column:estado"`
    CreadoEn     time.Time `json:"creado_en" gorm:"column:creado_en"`
}
