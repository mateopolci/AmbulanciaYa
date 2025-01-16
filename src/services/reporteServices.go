package services

import(
	"github.com/mateopolci/AmbulanciaYa/src/models"
	"gorm.io/gorm"
)

type ReporteService struct {
	db *gorm.DB
}

// Constructor del servicio
func NewReporteService(db *gorm.DB) *ReporteService {
	return &ReporteService{db: db}
}

// Obtener todos los reportes
func (s *ReporteService) GetAllReportes() ([]models.ReporteDTO, error) {
	var reportes []models.Reporte
	result := s.db.Find(&reportes)
	if result.Error != nil {
		return nil, result.Error
	}

	reportesDTO := make([]models.ReporteDTO, len(reportes))
	for i, rep := range reportes {
		reportesDTO[i] = rep.ReporteToDTO()
	}
	return reportesDTO, nil
}

// Obtener un reporte por su ID
func (s *ReporteService) GetReporteById(id string) (models.ReporteDTO, error) {
	var reporte models.Reporte
	result := s.db.First(&reporte, "id = ?", id)
	if result.Error != nil {
		return models.ReporteDTO{}, result.Error
	}
	return reporte.ReporteToDTO(), nil
}

// Crear un nuevo reporte
func (s *ReporteService) CreateReporte(reporteDTO models.ReporteDTO) (models.Reporte, error) {
	reporte := models.Reporte{
		Descripcion: reporteDTO.Descripcion,
		Fecha: reporteDTO.Fecha,
		Hora: reporteDTO.Hora,
		RequiereTraslado: reporteDTO.RequiereTraslado,
		AccidenteId: reporteDTO.AccidenteId,
	}
	
	result := s.db.Create(&reporte)
	return reporte, result.Error
}

// Actualizar un reporte existente
func (s *ReporteService) UpdateReporte(id string, reporteDTO models.ReporteDTO) (models.Reporte, error) {
	var reporte models.Reporte
    if err := s.db.First(&reporte, "id = ?", id).Error; err != nil {
        return reporte, err
    }

	reporte.Descripcion = reporteDTO.Descripcion
	reporte.Fecha = reporteDTO.Fecha
	reporte.Hora = reporteDTO.Hora
	reporte.RequiereTraslado = reporteDTO.RequiereTraslado
	reporte.AccidenteId = reporteDTO.AccidenteId

	result := s.db.Save(&reporte)
	return reporte, result.Error
}

// Eliminar un reporte existente
func (s *ReporteService) DeleteReporte(id string) error {
    result := s.db.Delete(&models.Reporte{}, "id = ?", id)
    return result.Error
}