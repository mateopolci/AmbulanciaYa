package models

// Modelo
type Paramedico struct {
	Id             string `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	NombreCompleto string `json:"nombreCompleto" gorm:"column:nombrecompleto;type:varchar(255);not null"`
	Dni            string `json:"dni" gorm:"type:varchar(20);not null"`
}

// DTO
type ParamedicoDTO struct {
	NombreCompleto string `json:"nombreCompleto"`
	Dni            string `json:"dni"`
}

// Método para convertir un Paramedico en un DTO
func (p *Paramedico) ToDTO() ParamedicoDTO {
	return ParamedicoDTO{
		NombreCompleto: p.NombreCompleto,
		Dni:            p.Dni,
	}
}
