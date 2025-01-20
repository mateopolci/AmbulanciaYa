package models

// Modelo
type Reporte struct {
	Id               string `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Descripcion      string `json:"descripcion" gorm:"type:text;not null"`
	Fecha            string `json:"fecha" gorm:"type:varchar(10);not null"`
	Hora             string `json:"hora" gorm:"type:varchar(8);not null"`
	RequiereTraslado bool   `json:"requiereTraslado" gorm:"column:requieretraslado;type:boolean;not null"`
	AccidenteId      string `json:"accidenteId" gorm:"column:accidenteid;type:uuid;not null"`
}

// DTO
type ReporteDTO struct {
	Id               string `json:"id"`
	Descripcion      string `json:"descripcion" gorm:"type:text;not null"`
	Fecha            string `json:"fecha" gorm:"type:varchar(10);not null"`
	Hora             string `json:"hora" gorm:"type:varchar(8);not null"`
	RequiereTraslado bool   `json:"requiereTraslado" gorm:"column:requieretraslado;type:boolean;not null"`
	AccidenteId      string `json:"accidenteId" gorm:"column:accidenteid;type:uuid;not null"`
}

// Método DTO de reporte
func (r *Reporte) ReporteToDTO() ReporteDTO {
	return ReporteDTO{
		Id:               r.Id,
		Descripcion:      r.Descripcion,
		Fecha:            r.Fecha,
		Hora:             r.Hora,
		RequiereTraslado: r.RequiereTraslado,
		AccidenteId:      r.AccidenteId,
	}
}
