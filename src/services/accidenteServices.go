package services

import (
    "github.com/mateopolci/AmbulanciaYa/src/models"
    "gorm.io/gorm"
)

type AccidenteService struct {
    db *gorm.DB
}

// Constructor del servicio
func NewAccidenteService(db *gorm.DB) *AccidenteService {
    return &AccidenteService{db: db}
}

// Obtener todos los accidentes
func (s *AccidenteService) GetAllAccidentes() ([]models.AccidenteDTO, error) {
    var accidentes []models.Accidente
    result := s.db.Find(&accidentes)
    if result.Error != nil {
        return nil, result.Error
    }

    accidentesDTO := make([]models.AccidenteDTO, len(accidentes))
    for i, acc := range accidentes {
        accidentesDTO[i] = acc.AccidenteToDTO()
    }
    return accidentesDTO, nil
}

// Obtener todos los accidentes con descripcion
func (s *AccidenteService) GetAllAccidentesDesc() ([]models.AccidenteDescDTO, error) {
    var results []struct {
        models.Accidente
        AmbulanciaPatente string  `gorm:"column:ambulancia_patente"`
        HospitalNombre    *string `gorm:"column:hospital_nombre"`
        PacienteNombre    string  `gorm:"column:paciente_nombre"`
    }

    err := s.db.Table("accidentes").
        Select("accidentes.*, ambulancias.patente as ambulancia_patente, hospitales.nombre as hospital_nombre, pacientes.nombrecompleto as paciente_nombre").
        Joins("LEFT JOIN ambulancias ON accidentes.ambulanciaid = ambulancias.id").
        Joins("LEFT JOIN hospitales ON accidentes.hospitalid = hospitales.id").
        Joins("LEFT JOIN pacientes ON accidentes.pacienteid = pacientes.id").
        Find(&results).Error

    if err != nil {
        return nil, err
    }

    accidentesDTO := make([]models.AccidenteDescDTO, len(results))
    for i, res := range results {
        accidentesDTO[i] = models.AccidenteDescDTO{
            Id:          res.Id,
            Direccion:   res.Direccion,
            Descripcion: res.Descripcion,
            Fecha:       res.Fecha,
            Hora:        res.Hora,
            Ambulancia:  res.AmbulanciaPatente,
            Hospital:    res.HospitalNombre,
            Paciente:    res.PacienteNombre,
        }
    }
    return accidentesDTO, nil
}

// Obtener un accidente por su ID
func (s *AccidenteService) GetAccidenteById(id string) (models.AccidenteDTO, error) {
    var accidente models.Accidente
    result := s.db.First(&accidente, "id = ?", id)
    if result.Error != nil {
        return models.AccidenteDTO{}, result.Error
    }
    return accidente.AccidenteToDTO(), nil
}

// Crear un nuevo accidente
func (s *AccidenteService) CreateAccidente(accidenteDTO models.AccidenteDTO) (models.Accidente, error) {
    accidente := models.Accidente{
        Direccion: accidenteDTO.Direccion,
        Descripcion: accidenteDTO.Descripcion,
        Fecha: accidenteDTO.Fecha,
        Hora: accidenteDTO.Hora,
        AmbulanciaId: accidenteDTO.AmbulanciaId,
        HospitalId: accidenteDTO.HospitalId,
        PacienteId: accidenteDTO.PacienteId,
    }
    
    result := s.db.Create(&accidente)
    return accidente, result.Error
}

// Actualizar un accidente existente
func (s *AccidenteService) UpdateAccidente(id string, accidenteDTO models.AccidenteDTO) (models.Accidente, error) {
    var accidente models.Accidente
    if err := s.db.First(&accidente, "id = ?", id).Error; err != nil {
        return accidente, err
    }

    accidente.Direccion = accidenteDTO.Direccion
    accidente.Descripcion = accidenteDTO.Descripcion
    accidente.Fecha = accidenteDTO.Fecha
    accidente.Hora = accidenteDTO.Hora
    accidente.AmbulanciaId = accidenteDTO.AmbulanciaId
    accidente.HospitalId = accidenteDTO.HospitalId
    accidente.PacienteId = accidenteDTO.PacienteId

    result := s.db.Save(&accidente)
    return accidente, result.Error
}

// Eliminar un accidente por su ID
func (s *AccidenteService) DeleteAccidente(id string) error {
    result := s.db.Delete(&models.Accidente{}, "id = ?", id)
    return result.Error
}