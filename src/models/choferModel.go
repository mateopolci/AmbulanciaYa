package models

// Modelo
type Chofer struct {
	Id             string `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	NombreCompleto string `json:"nombreCompleto" gorm:"column:nombrecompleto;type:varchar(255);not null"`
	Dni            string `json:"dni" gorm:"type:varchar(20);not null"`
}

// DTO
type ChoferDTO struct {
	Id             string `json:"id"`
	NombreCompleto string `json:"nombreCompleto"`
	Dni            string `json:"dni"`
}

// Método para convertir un Chofer en un DTO
func (c *Chofer) ChoferToDTO() ChoferDTO {
	return ChoferDTO{
		Id:             c.Id,
		NombreCompleto: c.NombreCompleto,
		Dni:            c.Dni,
	}
}

//Especificacion del nombre de la tabla para GORM
func (Chofer) TableName() string {
	return "choferes"
}
