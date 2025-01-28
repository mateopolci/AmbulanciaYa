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

// Obtener un reporte por su ID con descripcion
func (s *ReporteService) GetReporteByAccidenteId(accidenteId string) (models.ReporteDescDTO, error) {
    var result struct {
        models.Reporte
        HospitalNombre *string `gorm:"column:hospital_nombre"`
    }

    err := s.db.Table("reportes").
        Select("reportes.*, hospitales.nombre as hospital_nombre").
        Joins("LEFT JOIN accidentes ON reportes.accidenteid = accidentes.id").
        Joins("LEFT JOIN hospitales ON accidentes.hospitalid = hospitales.id").
        Where("reportes.accidenteid = ?", accidenteId).
        First(&result).Error

    if err != nil {
        return models.ReporteDescDTO{}, err
    }

    hospitalNombre := "-"
    if result.HospitalNombre != nil {
        hospitalNombre = *result.HospitalNombre
    }

    return models.ReporteDescDTO{
        Id:               result.Id,
        Descripcion:      result.Descripcion,
        Fecha:           result.Fecha,
        Hora:            result.Hora,
        RequiereTraslado: result.RequiereTraslado,
        AccidenteId:      result.AccidenteId,
        Hospital:         hospitalNombre,
    }, nil
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

// Crear un reporte y establecer el hospital en el accidente asociado

func (s *ReporteService) CreateReporteAndUpdateHospital(accidenteId string, postDTO models.ReportePostDTO) (models.Reporte, error) {
    tx := s.db.Begin()

    // Create new reporte
    reporte := models.Reporte{
        Descripcion:      postDTO.Descripcion,
        Fecha:           postDTO.Fecha,
        Hora:            postDTO.Hora,
        RequiereTraslado: postDTO.RequiereTraslado,
        AccidenteId:      accidenteId,
    }

    if err := tx.Create(&reporte).Error; err != nil {
        tx.Rollback()
        return models.Reporte{}, err
    }

    // Update accidente's hospitalId if provided
    if postDTO.HospitalId != nil {
        if err := tx.Model(&models.Accidente{}).
            Where("id = ?", accidenteId).
            Update("hospitalid", postDTO.HospitalId).Error; err != nil {
            tx.Rollback()
            return models.Reporte{}, err
        }
    }

    if err := tx.Commit().Error; err != nil {
        return models.Reporte{}, err
    }

    return reporte, nil
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

// Modificar un reporte y el hospitalId de su accidente
func (s *ReporteService) UpdateReporteAndHospital(id string, updateDTO models.ReporteUpdateDTO) (models.Reporte, error) {
    tx := s.db.Begin()

    var reporte models.Reporte
    if err := tx.First(&reporte, "id = ?", id).Error; err != nil {
        tx.Rollback()
        return models.Reporte{}, err
    }

    reporte.Descripcion = updateDTO.Descripcion
    reporte.Fecha = updateDTO.Fecha
    reporte.Hora = updateDTO.Hora
    reporte.RequiereTraslado = updateDTO.RequiereTraslado
    reporte.AccidenteId = updateDTO.AccidenteId

    if err := tx.Save(&reporte).Error; err != nil {
        tx.Rollback()
        return models.Reporte{}, err
    }

    // Update accidente's hospitalId (set to NULL if not provided)
    updateQuery := tx.Model(&models.Accidente{}).Where("id = ?", updateDTO.AccidenteId)
    if updateDTO.HospitalId != nil {
        err := updateQuery.Update("hospitalid", updateDTO.HospitalId).Error
        if err != nil {
            tx.Rollback()
            return models.Reporte{}, err
        }
    } else {
        err := updateQuery.Update("hospitalid", nil).Error
        if err != nil {
            tx.Rollback()
            return models.Reporte{}, err
        }
    }

    if err := tx.Commit().Error; err != nil {
        return models.Reporte{}, err
    }

    return reporte, nil
}

// Eliminar un reporte existente
func (s *ReporteService) DeleteReporte(id string) error {
    result := s.db.Delete(&models.Reporte{}, "id = ?", id)
    return result.Error
}