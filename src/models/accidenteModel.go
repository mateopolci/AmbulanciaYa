package models

// Modelo
type Accidente struct {
	Id           string  `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Direccion    string  `json:"direccion" gorm:"type:varchar(255);not null"`
	Descripcion  string  `json:"descripcion" gorm:"type:text;not null"`
	Fecha        string  `json:"fecha" gorm:"type:varchar(10);not null"`
	Hora         string  `json:"hora" gorm:"type:varchar(8);not null"`
	AmbulanciaId string  `json:"ambulanciaId" gorm:"column:ambulanciaid;type:uuid;not null"`
	HospitalId   *string `json:"hospitalId,omitempty" gorm:"column:hospitalid;type:uuid;null"`
	PacienteId   *string `json:"pacienteId" gorm:"column:pacienteid;type:uuid;null"`
}

// DTO
type AccidenteDTO struct {
	Id           string  `json:"id"`
	Direccion    string  `json:"direccion"`
	Descripcion  string  `json:"descripcion"`
	Fecha        string  `json:"fecha"`
	Hora         string  `json:"hora"`
	AmbulanciaId string  `json:"ambulanciaId"`
	HospitalId   *string `json:"hospitalId,omitempty"`
	PacienteId   *string `json:"pacienteId,omitempty"`
}

// Método DTO de accidente
func (a *Accidente) AccidenteToDTO() AccidenteDTO {
	return AccidenteDTO{
		Id:           a.Id,
		Direccion:    a.Direccion,
		Descripcion:  a.Descripcion,
		Fecha:        a.Fecha,
		Hora:         a.Hora,
		AmbulanciaId: a.AmbulanciaId,
		HospitalId:   a.HospitalId,
		PacienteId:   a.PacienteId,
	}
}

// DTO para service de descripcion de Accidentes
type AccidenteDescDTO struct {
	Id           string  `json:"id"`
	Direccion    string  `json:"direccion"`
	Descripcion  string  `json:"descripcion"`
	Fecha        string  `json:"fecha"`
	Hora         string  `json:"hora"`
	Ambulancia   string  `json:"ambulancia"`
	Hospital     *string `json:"hospital,omitempty"`
	Paciente     string  `json:"paciente"`
	TieneReporte bool    `json:"reporte"`
}
