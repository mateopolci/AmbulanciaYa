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

// GetAll obtiene todos los accidentes
func (s *AccidenteService) GetAll() ([]models.AccidenteDTO, error) {
    var accidentes []models.Accidente
    result := s.db.Find(&accidentes)
    if result.Error != nil {
        return nil, result.Error
    }

    accidentesDTO := make([]models.AccidenteDTO, len(accidentes))
    for i, acc := range accidentes {
        accidentesDTO[i] = acc.ToDTO()
    }
    return accidentesDTO, nil
}

// GetById obtiene un accidente por su ID
func (s *AccidenteService) GetById(id string) (models.AccidenteDTO, error) {
    var accidente models.Accidente
    result := s.db.First(&accidente, "id = ?", id)
    if result.Error != nil {
        return models.AccidenteDTO{}, result.Error
    }
    return accidente.ToDTO(), nil
}

// Create crea un nuevo accidente
func (s *AccidenteService) Create(accidenteDTO models.AccidenteDTO) (models.Accidente, error) {
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

// Update actualiza un accidente existente
func (s *AccidenteService) Update(id string, accidenteDTO models.AccidenteDTO) (models.Accidente, error) {
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

// Delete elimina un accidente por su ID
func (s *AccidenteService) Delete(id string) error {
    result := s.db.Delete(&models.Accidente{}, "id = ?", id)
    return result.Error
}