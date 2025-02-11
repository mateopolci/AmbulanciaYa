package models

// Modelo
type Hospital struct {
	Id        string `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Nombre    string `json:"nombre" gorm:"type:varchar(255);not null"`
	Direccion string `json:"direccion" gorm:"type:varchar(255);not null"`
}

// DTO
type HospitalDTO struct {
	Id        string `json:"id"`
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
}

// Metodo DTO de hospital
func (h *Hospital) HospitalToDTO() HospitalDTO {
	return HospitalDTO{
		Id:        h.Id,
		Nombre:    h.Nombre,
		Direccion: h.Direccion,
	}
}

//Especificacion del nombre de la tabla para GORM
func (Hospital) TableName() string {
	return "hospitales"
}
