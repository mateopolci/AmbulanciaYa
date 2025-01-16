package services

import(
	"github.com/mateopolci/AmbulanciaYa/src/models"
	"gorm.io/gorm"
)

type HospitalService struct {
	db *gorm.DB
}

// Constructor del servicio
func NewHospitalService(db *gorm.DB) *HospitalService {
	return &HospitalService{db: db}
}

// Obtener todos los hospitales
func (s *HospitalService) GetAllHospitales() ([]models.HospitalDTO, error) {
	var hospitales []models.Hospital
	result := s.db.Find(&hospitales)
	if result.Error != nil {
		return nil, result.Error
	}

	hospitalesDTO := make([]models.HospitalDTO, len(hospitales))
	for i, hosp := range hospitales {
		hospitalesDTO[i] = hosp.HospitalToDTO()
	}
	return hospitalesDTO, nil
}

// Obtener un hospital por su ID
func (s *HospitalService) GetHospitalById(id string) (models.HospitalDTO, error) {
	var hospital models.Hospital
	result := s.db.First(&hospital, "id = ?", id)
	if result.Error != nil {
		return models.HospitalDTO{}, result.Error
	}
	return hospital.HospitalToDTO(), nil
}

// Crear un nuevo hospital
func (s *HospitalService) CreateHospital(hospitalDTO models.HospitalDTO) (models.Hospital, error) {
	hospital := models.Hospital{
		Nombre: hospitalDTO.Nombre,
		Direccion: hospitalDTO.Direccion,
	}
	
	result := s.db.Create(&hospital)
	return hospital, result.Error
}

// Actualizar un hospital existente
func (s *HospitalService) UpdateHospital(id string, hospitalDTO models.HospitalDTO) (models.Hospital, error) {
	var hospital models.Hospital
	if err := s.db.First(&hospital, "id = ?", id).Error; err != nil {
		return hospital, err
	}

	hospital.Nombre = hospitalDTO.Nombre
	hospital.Direccion = hospitalDTO.Direccion

	result := s.db.Save(&hospital)
	return hospital, result.Error
}

// Eliminar un hospital existente
func (s *HospitalService) DeleteHospital(id string) error {
    result := s.db.Delete(&models.Hospital{}, "id = ?", id)
    return result.Error
}