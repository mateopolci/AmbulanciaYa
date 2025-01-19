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

func (s *ParamedicoService) Login(email, password string) (string, error) {
	var paramedico models.Paramedico
	if err := s.db.Where("email = ?", email).First(&paramedico).Error; err != nil {
		return "", errors.New("credenciales inválidas")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(paramedico.Password), []byte(password)); err != nil {
		return "", errors.New("credenciales inválidas")
	}

	claims := jwt.MapClaims{
		"id":      paramedico.Id,
		"email":   paramedico.Email,
		"isAdmin": paramedico.IsAdmin,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(middleware.GetSecretKey()))
}
