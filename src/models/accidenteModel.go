package models

type Accidente struct {
    Id string `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Direccion string `json:"direccion" gorm:"type:varchar(255);not null"`
    Descripcion string `json:"descripcion" gorm:"type:text;not null"`
    Fecha string `json:"fecha" gorm:"type:varchar(10);not null"`
    Hora string `json:"hora" gorm:"type:varchar(8);not null"`
    AmbulanciaId string `json:"ambulanciaId" gorm:"column:ambulanciaid;type:uuid;not null"`
    HospitalId string `json:"hospitalId" gorm:"column:hospitalid;type:uuid"`
    PacienteId string `json:"pacienteId" gorm:"column:pacienteid;type:uuid;not null"`
}

type AccidenteDTO struct {
    Direccion string `json:"direccion"`
    Descripcion string `json:"descripcion"`
    Fecha string `json:"fecha"`
    Hora string `json:"hora"`
    AmbulanciaId string `json:"ambulanciaId"`
    HospitalId string `json:"hospitalId"`
    PacienteId string `json:"pacienteId"`
}