package services

import (
	"time"

	"github.com/mateopolci/AmbulanciaYa/src/models"
	"gorm.io/gorm"
)

type AmbulanciaService struct {
	db               *gorm.DB
	pacienteService  *PacienteService
	accidenteService *AccidenteService
}

// Constructor del servicio
func NewAmbulanciaService(db *gorm.DB, pacienteService *PacienteService, accidenteService *AccidenteService) *AmbulanciaService {
	return &AmbulanciaService{
		db:               db,
		pacienteService:  pacienteService,
		accidenteService: accidenteService,
	}
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
			Id:              res.Id,
			Patente:         res.Patente,
			Inventario:      res.Inventario,
			Vtv:             res.Vtv,
			Seguro:          res.Seguro,
			Chofer:          res.ChoferNombre,
			Paramedico:      res.ParamedicoNombre,
			Base:            res.Base,
			Cadenas:         res.Cadenas,
			Antinieblas:     res.Antinieblas,
			CubiertasLluvia: res.CubiertasLluvia,
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

// Obtener el id de la primera ambulancia disponible
func (s *AmbulanciaService) GetAmbulanciaDisp(descripcion string) (models.AmbulanciaDTO, error) {

	var ambulancia models.Ambulancia

	// Validación de ambulancia para  "Veloway"
	if descripcion == "Veloway" {

		datos := GetDatosVeloway()

		query := s.db

		query = query.Where("vtv = ? AND seguro = ? AND base = ?",
			true, true, true)

		if datos.EnfermedadCardiaca != nil || datos.EnfermedadRespiratoria != nil || datos.Alergias != nil {
			query = query.Where("inventario = ?", true)
		}

		result := query.First(&ambulancia)
		if result.Error != nil {
			return models.AmbulanciaDTO{}, result.Error
		}
		return ambulancia.AmbulanciaToDTO(), nil
	}

	// Validación de ambulancia para "Los Pinos"
	if descripcion == "Los Pinos" {
		datos := GetDatosLosPinos()

		query := s.db

		query = query.Where("vtv = ? AND seguro = ? AND base = ?",
			true, true, true)

		if datos.Nieve >= 30 {
			query = query.Where("cadenas = ?", true)
		}
		if datos.Lluvia >= 40 {
			query = query.Where("cubiertaslluvia = ?", true)
		}
		if datos.Visibilidad <= 50 {
			query = query.Where("antinieblas = ?", true)
		}

		result := query.First(&ambulancia)
		if result.Error != nil {
			return models.AmbulanciaDTO{}, result.Error
		}
		return ambulancia.AmbulanciaToDTO(), nil
	}

	result := s.db.Where(
		"vtv = ? AND seguro = ? AND base = ?",
		true, true, true,
	).First(&ambulancia)

	if result.Error != nil {
		return models.AmbulanciaDTO{}, result.Error
	}

	return ambulancia.AmbulanciaToDTO(), nil
}

// Crear una nueva ambulancia
func (s *AmbulanciaService) CreateAmbulancia(ambulanciaDTO models.AmbulanciaDTO) (models.Ambulancia, error) {
	ambulancia := models.Ambulancia{
		Patente:         ambulanciaDTO.Patente,
		Inventario:      ambulanciaDTO.Inventario,
		Vtv:             ambulanciaDTO.Vtv,
		Seguro:          ambulanciaDTO.Seguro,
		ChoferId:        ambulanciaDTO.ChoferId,
		ParamedicoId:    ambulanciaDTO.ParamedicoId,
		Base:            ambulanciaDTO.Base,
		Cadenas:         ambulanciaDTO.Cadenas,
		Antinieblas:     ambulanciaDTO.Antinieblas,
		CubiertasLluvia: ambulanciaDTO.CubiertasLluvia,
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
	ambulancia.Cadenas = ambulanciaDTO.Cadenas
	ambulancia.Antinieblas = ambulanciaDTO.Antinieblas
	ambulancia.CubiertasLluvia = ambulanciaDTO.CubiertasLluvia

	result := s.db.Save(&ambulancia)
	return ambulancia, result.Error
}

// Eliminar una ambulancia por su ID
func (s *AmbulanciaService) DeleteAmbulancia(id string) error {
	result := s.db.Delete(&models.Ambulancia{}, "id = ?", id)
	return result.Error
}

// Pedido de ambulancia
func (s *AmbulanciaService) PedidoAmbulancia(pedido models.AmbulanciaPedidoDTO) (string, error) {
	// Recuperar ambulancia disponible
	var ambulanciaDisp models.AmbulanciaDTO
	var err error
	maxIntentos := 7

	for intento := 0; intento < maxIntentos; intento++ {
		ambulanciaDisp, err = s.GetAmbulanciaDisp(pedido.Descripcion)
		if err == nil && ambulanciaDisp.Id != "" {
			break
		}
		if intento < maxIntentos-1 {
			time.Sleep(10 * time.Second)
		}
	}

	if err != nil || ambulanciaDisp.Id == "" {
		return "No se encuentran ambulancias disponibles", err
	}
	idAmbulanciaEncontrada := ambulanciaDisp.Id

	// Inicializar pacienteId
	var pacienteId *string

	// Solo procesar paciente si se proporcionan nombre y teléfono
	if pedido.Nombre != "" && pedido.Telefono != "" {
		paciente, err := s.pacienteService.GetByTelefono(pedido.Telefono)
		if err != nil || paciente.Id == "" {
			nuevoPaciente, err := s.pacienteService.Create(models.PacienteDTO{
				NombreCompleto: pedido.Nombre,
				Telefono:       pedido.Telefono,
			})
			if err != nil {
				return "Error al crear paciente", err
			}
			pacienteId = &nuevoPaciente.Id
		} else {
			pacienteId = &paciente.Id
		}
	}

	// Crear accidente y enviar ambulancia
	loc, err := time.LoadLocation("America/Argentina/Buenos_Aires")
	if err != nil {
		return "Error al configurar zona horaria", err
	}
	now := time.Now().In(loc)
	accidente := models.AccidenteDTO{
		Direccion:    pedido.Direccion,
		Descripcion:  pedido.Descripcion,
		Fecha:        now.Format("2006-01-02"),
		Hora:         now.Format("15:04"),
		AmbulanciaId: idAmbulanciaEncontrada,
		PacienteId:   pacienteId,
	}

	_, err = s.accidenteService.CreateAccidenteAndSendAmbulancia(accidente)
	if err != nil {
		return "Error al registrar el accidente", err
	}

	return "La ambulancia ha sido enviada", nil
}
