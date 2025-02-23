package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mateopolci/AmbulanciaYa/src/middleware"
	"github.com/mateopolci/AmbulanciaYa/src/models"
	"golang.org/x/crypto/bcrypt"
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
		paramedicosDTO[i] = param.ParamedicoToDTO()
	}
	return paramedicosDTO, nil
}

// GetAll obtiene todos los paramedicos que no esten asociados a una ambulancia
func (s *ParamedicoService) GetAllDisp() ([]models.ParamedicoDTO, error) {
	var paramedicos []models.Paramedico
	result := s.db.
		Joins("LEFT JOIN ambulancias ON ambulancias.paramedicoid = paramedicos.id").
		Where("ambulancias.id IS NULL").
		Find(&paramedicos)

	if result.Error != nil {
		return nil, result.Error
	}

	paramedicosDTO := make([]models.ParamedicoDTO, len(paramedicos))
	for i, param := range paramedicos {
		paramedicosDTO[i] = param.ParamedicoToDTO()
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
	return paramedico.ParamedicoToDTO(), nil
}

// Create crea un nuevo paramedico
func (s *ParamedicoService) Create(paramedicoDTO models.ParamedicoDTO) (models.Paramedico, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(paramedicoDTO.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.Paramedico{}, err
	}

	paramedico := models.Paramedico{
		NombreCompleto: paramedicoDTO.NombreCompleto,
		Dni:            paramedicoDTO.Dni,
		Email:          paramedicoDTO.Email,
		Password:       string(hashedPassword),
		IsAdmin:        paramedicoDTO.IsAdmin,
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
	paramedico.Email = paramedicoDTO.Email
	paramedico.IsAdmin = paramedicoDTO.IsAdmin

	if paramedicoDTO.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(paramedicoDTO.Password), bcrypt.DefaultCost)
		if err != nil {
			return paramedico, err
		}
		paramedico.Password = string(hashedPassword)
	}

	result := s.db.Save(&paramedico)
	return paramedico, result.Error
}

// Delete elimina un paramedico por su ID
func (s *ParamedicoService) Delete(id string) error {
	result := s.db.Delete(&models.Paramedico{}, "id = ?", id)
	return result.Error
}

// Login paramedico/admin
func (s *ParamedicoService) Login(email, password string) (*models.LoginResponse, error) {
	var paramedico models.Paramedico
	if err := s.db.Where("email = ?", email).First(&paramedico).Error; err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(paramedico.Password), []byte(password)); err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	// Creacion del payload del token
	claims := jwt.MapClaims{
		"id":      paramedico.Id,
		"email":   paramedico.Email,
		"isAdmin": paramedico.IsAdmin,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// Creacion del JWT con header, payload y signature
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firma del token con la clave secreta
	tokenString, err := token.SignedString([]byte(middleware.GetSecretKey()))

	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token:   tokenString,
		IsAdmin: paramedico.IsAdmin,
	}, nil
}

func (s *ParamedicoService) UpdateEmail(paramedicoId string, currentPassword string, newEmail string) error {
	var paramedico models.Paramedico
	if err := s.db.First(&paramedico, "id = ?", paramedicoId).Error; err != nil {
		return err
	}

	// Verificamos la contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(paramedico.Password), []byte(currentPassword)); err != nil {
		return errors.New("contraseña actual incorrecta")
	}

	// Actualizamos el email
	result := s.db.Model(&paramedico).Update("email", newEmail)
	return result.Error
}

func (s *ParamedicoService) UpdatePassword(paramedicoId string, currentPassword string, newPassword string) error {
	// Verificamos el id del paramedico
	var paramedico models.Paramedico
	if err := s.db.First(&paramedico, "id = ?", paramedicoId).Error; err != nil {
		return err
	}

	// Verificamos la contraseña actual
	if err := bcrypt.CompareHashAndPassword([]byte(paramedico.Password), []byte(currentPassword)); err != nil {
		return errors.New("contraseña actual incorrecta")
	}

	// Hashear la nueva contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Actualizar la contraseña
	result := s.db.Model(&paramedico).Update("password", string(hashedPassword))
	return result.Error
}
