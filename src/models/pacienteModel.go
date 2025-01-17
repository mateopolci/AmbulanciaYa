package models

// Modelo
type Paciente struct {
	Id             string `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	NombreCompleto string `json:"nombreCompleto" gorm:"column:nombrecompleto;type:varchar(255);not null"`
	Telefono       string `json:"telefono" gorm:"type:varchar(20);not null"`
}

// DTO
type PacienteDTO struct {
	NombreCompleto string `json:"nombreCompleto"`
	Telefono       string `json:"telefono"`
}

// Método para convertir un Accidente en un DTO
func (p *Paciente) PacienteToDTO() PacienteDTO {
	return PacienteDTO{
		NombreCompleto: p.NombreCompleto,
		Telefono:       p.Telefono,
	}
}
