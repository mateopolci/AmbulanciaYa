package services

import (
	"github.com/mateopolci/AmbulanciaYa/src/models"
	"gorm.io/gorm"
)

type ParamedicoService struct {
	db *gorm.DB
}

// Constructor del servicio
func NewParamedicoService(db *gorm.DB) *ParamedicoService {
	return &ParamedicoService{db: db}
}

// GetAll obtiene todos los paramedico
func (s *ParamedicoService) GetAll() ([]models.ParamedicoDTO, error) {
	var paramedicos []models.Paramedico
	result := s.db.Find(&paramedicos)
	if result.Error != nil {
		return nil, result.Error
	}

	paramedicosDTO := make([]models.ParamedicoDTO, len(paramedicos))
	for i, param := range paramedicos {
		paramedicosDTO[i] = param.ToDTO()
	}
	return paramedicosDTO, nil
}

// GetById obtiene un paramedico por su ID
func (s *ParamedicoService) GetById(id string) (models.ParamedicoDTO, error) {
	var paramedico models.Paramedico
	result := s.db.First(&paramedico, "id = ?", id)
	if result.Error != nil {
		return models.ParamedicoDTO{}, result.Error
	}
	return paramedico.ToDTO(), nil
}

// Create crea un nuevo paramedico
func (s *ParamedicoService) Create(paramedicoDTO models.ParamedicoDTO) (models.Paramedico, error) {
	paramedico := models.Paramedico{
		NombreCompleto: paramedicoDTO.NombreCompleto,
		Dni:            paramedicoDTO.Dni,
	}

	result := s.db.Create(&paramedico)
	return paramedico, result.Error
}

// Update actualiza un paramedico existente
func (s *ParamedicoService) Update(id string, paramedicoDTO models.ParamedicoDTO) (models.Paramedico, error) {
	var paramedico models.Paramedico
	if err := s.db.First(&paramedico, "id = ?", id).Error; err != nil {
		return paramedico, err
	}

	paramedico.NombreCompleto = paramedicoDTO.NombreCompleto
	paramedico.Dni = paramedicoDTO.Dni

	result := s.db.Save(&paramedico)
	return paramedico, result.Error
}

// Delete elimina un paramedico por su ID
func (s *ParamedicoService) Delete(id string) error {
	result := s.db.Delete(&models.Paramedico{}, "id = ?", id)
	return result.Error
}
