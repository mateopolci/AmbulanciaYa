package services

import (
	"github.com/mateopolci/AmbulanciaYa/src/models"
	"gorm.io/gorm"
)

type ChoferService struct {
	db *gorm.DB
}

// Constructor del servicio
func NewChoferService(db *gorm.DB) *ChoferService {
	return &ChoferService{db: db}
}

// GetAll obtiene todos los chofer
func (s *ChoferService) GetAll() ([]models.ChoferDTO, error) {
	var choferes []models.Chofer
	result := s.db.Find(&choferes)
	if result.Error != nil {
		return nil, result.Error
	}

	choferesDTO := make([]models.ChoferDTO, len(choferes))
	for i, chof := range choferes {
		choferesDTO[i] = chof.ChoferToDTO()
	}
	return choferesDTO, nil
}

// GetById obtiene un choferes por su ID
func (s *ChoferService) GetById(id string) (models.ChoferDTO, error) {
	var chofer models.Chofer
	result := s.db.First(&chofer, "id = ?", id)
	if result.Error != nil {
		return models.ChoferDTO{}, result.Error
	}
	return chofer.ChoferToDTO(), nil
}

// Create crea un nuevo chofer
func (s *ChoferService) Create(choferDTO models.ChoferDTO) (models.Chofer, error) {
	chofer := models.Chofer{
		NombreCompleto: choferDTO.NombreCompleto,
		Dni:            choferDTO.Dni,
	}

	result := s.db.Create(&chofer)
	return chofer, result.Error
}

// Update actualiza un chofer existente
func (s *ChoferService) Update(id string, choferDTO models.ChoferDTO) (models.Chofer, error) {
	var chofer models.Chofer
	if err := s.db.First(&chofer, "id = ?", id).Error; err != nil {
		return chofer, err
	}

	chofer.NombreCompleto = choferDTO.NombreCompleto
	chofer.Dni = choferDTO.Dni

	result := s.db.Save(&chofer)
	return chofer, result.Error
}

// Delete elimina un chofer por su ID
func (s *ChoferService) Delete(id string) error {
	result := s.db.Delete(&models.Chofer{}, "id = ?", id)
	return result.Error
}
