package services

import (
	"github.com/mateopolci/AmbulanciaYa/src/models"
	"gorm.io/gorm"
)

type PacienteService struct {
	db *gorm.DB
}

// Constructor del servicio
func NewPacienteService(db *gorm.DB) *PacienteService {
	return &PacienteService{db: db}
}

// GetAll obtiene todos los paciente
func (s *PacienteService) GetAll() ([]models.PacienteDTO, error) {
	var pacientes []models.Paciente
	result := s.db.Find(&pacientes)
	if result.Error != nil {
		return nil, result.Error
	}

	pacientesDTO := make([]models.PacienteDTO, len(pacientes))
	for i, pac := range pacientes {
		pacientesDTO[i] = pac.PacienteToDTO()
	}
	return pacientesDTO, nil
}

// Consulta un paciente por su telefono
func (s *PacienteService) GetByTelefono(telefono string) (models.PacienteDTO, error) {
	var paciente models.Paciente
	result := s.db.First(&paciente, "telefono = ?", telefono)
	if result.Error != nil {
		return models.PacienteDTO{}, result.Error
	}
	return paciente.PacienteToDTO(), nil
}

// GetById obtiene un paciente por su ID
func (s *PacienteService) GetById(id string) (models.PacienteDTO, error) {
	var paciente models.Paciente
	result := s.db.First(&paciente, "id = ?", id)
	if result.Error != nil {
		return models.PacienteDTO{}, result.Error
	}
	return paciente.PacienteToDTO(), nil
}

// Create crea un nuevo paciente
func (s *PacienteService) Create(pacienteDTO models.PacienteDTO) (models.Paciente, error) {
	paciente := models.Paciente{
		NombreCompleto: pacienteDTO.NombreCompleto,
		Telefono:       pacienteDTO.Telefono,
	}

	result := s.db.Create(&paciente)
	return paciente, result.Error
}

// Update actualiza un paciente existente
func (s *PacienteService) Update(id string, pacienteDTO models.PacienteDTO) (models.Paciente, error) {
	var paciente models.Paciente
	if err := s.db.First(&paciente, "id = ?", id).Error; err != nil {
		return paciente, err
	}

	paciente.NombreCompleto = pacienteDTO.NombreCompleto
	paciente.Telefono = pacienteDTO.Telefono

	result := s.db.Save(&paciente)
	return paciente, result.Error
}

// Delete elimina un paciente por su ID
func (s *PacienteService) Delete(id string) error {
	result := s.db.Delete(&models.Paciente{}, "id = ?", id)
	return result.Error
}
