package models

// Modelo
type Paramedico struct {
	Id             string `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	NombreCompleto string `json:"nombreCompleto" gorm:"column:nombrecompleto;type:varchar(255);not null"`
	Dni            string `json:"dni" gorm:"type:varchar(20);not null"`
	Email          string `json:"email" gorm:"type:varchar(50);not null"`
	Password       string `json:"password" gorm:"type:varchar(20);not null"`
	IsAdmin        bool   `json:"isAdmin" gorm:"column:isadmin;type:boolean;not null"`
}

// DTO
type ParamedicoDTO struct {
	Id             string `json:"id"`
	NombreCompleto string `json:"nombreCompleto"`
	Dni            string `json:"dni"`
	Email          string `json:"email"`
	Password       string `json:"password,omitempty"`
	IsAdmin        bool   `json:"isAdmin"`
}

// Método para convertir un Paramedico en un DTO
func (p *Paramedico) ParamedicoToDTO() ParamedicoDTO {
	return ParamedicoDTO{
		Id:             p.Id,
		NombreCompleto: p.NombreCompleto,
		Dni:            p.Dni,
		Email:          p.Email,
		IsAdmin:        p.IsAdmin,
	}
}

// DTO para el login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}
