package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/models"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

type PacienteController struct {
	service *services.PacienteService
}

func NewPacienteController(service *services.PacienteService) *PacienteController {
	return &PacienteController{service: service}
}

func (c *PacienteController) GetPacientes(ctx *gin.Context) {
	pacientesDTO, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, pacientesDTO)
}

func (c *PacienteController) GetPaciente(ctx *gin.Context) {
	id := ctx.Param("id")
	pacienteDTO, err := c.service.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Paciente no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, pacienteDTO)
}

func (c *PacienteController) PostPaciente(ctx *gin.Context) {
	var pacienteDTO models.PacienteDTO
	if err := ctx.ShouldBindJSON(&pacienteDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paciente, err := c.service.Create(pacienteDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, paciente)
}

func (c *PacienteController) PutPaciente(ctx *gin.Context) {
	id := ctx.Param("id")
	var pacienteDTO models.PacienteDTO
	if err := ctx.ShouldBindJSON(&pacienteDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paciente, err := c.service.Update(id, pacienteDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, paciente)
}

func (c *PacienteController) DeletePaciente(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Paciente eliminado"})
}
