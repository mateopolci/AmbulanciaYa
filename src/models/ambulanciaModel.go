package models

// Modelo
type Ambulancia struct {
	Id           string `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Patente      string `json:"patente" gorm:"type:varchar(20);not null"`
	Inventario   bool   `json:"inventario" gorm:"type:boolean;not null"`
	Vtv          bool   `json:"vtv" gorm:"type:boolean;not null"`
	Seguro       bool   `json:"seguro" gorm:"type:boolean;not null"`
	ChoferId     string `json:"choferId" gorm:"column:choferid;type:uuid;not null"`
	ParamedicoId string `json:"paramedicoId" gorm:"column:paramedicoid;type:uuid;not null"`
}

// DTO
type AmbulanciaDTO struct {
	Patente      string `json:"patente"`
	Inventario   bool   `json:"inventario"`
	Vtv          bool   `json:"vtv"`
	Seguro       bool   `json:"seguro"`
	ChoferId     string `json:"choferId"`
	ParamedicoId string `json:"paramedicoId"`
}

// Método DTO de ambulancia
func (a *Ambulancia) AmbulanciaToDTO() AmbulanciaDTO {
	return AmbulanciaDTO{
		Patente:      a.Patente,
		Inventario:   a.Inventario,
		Vtv:          a.Vtv,
		Seguro:       a.Seguro,
		ChoferId:     a.ChoferId,
		ParamedicoId: a.ParamedicoId,
	}
}

//Especificacion del nombre de la bdd
func (Ambulancia) TableName() string {
    return "ambulancias"
}