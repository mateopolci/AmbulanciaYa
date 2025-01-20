package services

import (
	"github.com/mateopolci/AmbulanciaYa/src/models"
	"gorm.io/gorm"
)

type AmbulanciaService struct {
	db *gorm.DB
}

// Constructor del servicio
func NewAmbulanciaService(db *gorm.DB) *AmbulanciaService {
	return &AmbulanciaService{db: db}
}

// Obtener todas las ambulancias
func (s *AmbulanciaService) GetAllAmbulancias() ([]models.AmbulanciaDTO, error) {
	var ambulancias []models.Ambulancia
	result := s.db.Find(&ambulancias)
	if result.Error != nil {
		return nil, result.Error
	}

	ambulanciasDTO := make([]models.AmbulanciaDTO, len(ambulancias))
	for i, amb := range ambulancias {
		ambulanciasDTO[i] = amb.AmbulanciaToDTO()
	}
	return ambulanciasDTO, nil
}

// Obtener todas las ambulancias con descripcion
func (s *AmbulanciaService) GetAllAmbulanciasDesc() ([]models.AmbulanciaDescDTO, error) {
    var results []struct {
        models.Ambulancia
        ChoferNombre     string `gorm:"column:chofer_nombre"`
        ParamedicoNombre string `gorm:"column:paramedico_nombre"`
    }

    err := s.db.Table("ambulancias").
        Select("ambulancias.*, choferes.nombrecompleto as chofer_nombre, paramedicos.nombrecompleto as paramedico_nombre").
        Joins("LEFT JOIN choferes ON ambulancias.choferid = choferes.id").
        Joins("LEFT JOIN paramedicos ON ambulancias.paramedicoid = paramedicos.id").
        Find(&results).Error

    if err != nil {
        return nil, err
    }

    ambulanciasDTO := make([]models.AmbulanciaDescDTO, len(results))
    for i, res := range results {
        ambulanciasDTO[i] = models.AmbulanciaDescDTO{
            Id:    res.Id,
            Patente:    res.Patente,
            Inventario: res.Inventario,
            Vtv:        res.Vtv,
            Seguro:     res.Seguro,
            Chofer:     res.ChoferNombre,
            Paramedico: res.ParamedicoNombre,
            Base:       res.Base,
        }
    }
    return ambulanciasDTO, nil
}

// Obtener una ambulancia por su ID
func (s *AmbulanciaService) GetAmbulanciaById(id string) (models.AmbulanciaDTO, error) {
	var ambulancia models.Ambulancia
	result := s.db.First(&ambulancia, "id = ?", id)
	if result.Error != nil {
		return models.AmbulanciaDTO{}, result.Error
	}
	return ambulancia.AmbulanciaToDTO(), nil
}

// Crear una nueva ambulancia
func (s *AmbulanciaService) CreateAmbulancia(ambulanciaDTO models.AmbulanciaDTO) (models.Ambulancia, error) {
	ambulancia := models.Ambulancia{
		Patente:      ambulanciaDTO.Patente,
		Inventario:   ambulanciaDTO.Inventario,
		Vtv:          ambulanciaDTO.Vtv,
		Seguro:       ambulanciaDTO.Seguro,
		ChoferId:     ambulanciaDTO.ChoferId,
		ParamedicoId: ambulanciaDTO.ParamedicoId,
		Base:          ambulanciaDTO.Base,
	}

	result := s.db.Create(&ambulancia)
	return ambulancia, result.Error
}

// Actualizar una ambulancia existente
func (s *AmbulanciaService) UpdateAmbulancia(id string, ambulanciaDTO models.AmbulanciaDTO) (models.Ambulancia, error) {
	var ambulancia models.Ambulancia
	if err := s.db.First(&ambulancia, "id = ?", id).Error; err != nil {
		return models.Ambulancia{}, err
	}

	ambulancia.Patente = ambulanciaDTO.Patente
	ambulancia.Inventario = ambulanciaDTO.Inventario
	ambulancia.Vtv = ambulanciaDTO.Vtv
	ambulancia.Seguro = ambulanciaDTO.Seguro
	ambulancia.ChoferId = ambulanciaDTO.ChoferId
	ambulancia.ParamedicoId = ambulanciaDTO.ParamedicoId
	ambulancia.Base = ambulanciaDTO.Base

	result := s.db.Save(&ambulancia)
	return ambulancia, result.Error
}

// Eliminar una ambulancia por su ID
func (s *AmbulanciaService) DeleteAmbulancia(id string) error {
	result := s.db.Delete(&models.Ambulancia{}, "id = ?", id)
	return result.Error
}