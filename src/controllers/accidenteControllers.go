package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/models"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

type AccidenteController struct {
    service *services.AccidenteService
}

func NewAccidenteController(service *services.AccidenteService) *AccidenteController {
    return &AccidenteController{service: service}
}

func (c *AccidenteController) GetAccidentes(ctx *gin.Context) {
    accidentesDTO, err := c.service.GetAllAccidentes()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, accidentesDTO)
}

func (c *AccidenteController) GetAccidente(ctx *gin.Context) {
    id := ctx.Param("id")
    accidenteDTO, err := c.service.GetAccidenteById(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Accidente no encontrado"})
        return
    }
    ctx.JSON(http.StatusOK, accidenteDTO)
}

func (c *AccidenteController) PostAccidente(ctx *gin.Context) {
    var accidenteDTO models.AccidenteDTO
    if err := ctx.ShouldBindJSON(&accidenteDTO); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    accidente, err := c.service.CreateAccidente(accidenteDTO)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, accidente.AccidenteToDTO())
}

func (c *AccidenteController) PutAccidente(ctx *gin.Context) {
    id := ctx.Param("id")
    var accidenteDTO models.AccidenteDTO
    if err := ctx.ShouldBindJSON(&accidenteDTO); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    accidente, err := c.service.UpdateAccidente(id, accidenteDTO)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, accidente.AccidenteToDTO())
}

func (c *AccidenteController) DeleteAccidente(ctx *gin.Context) {
    id := ctx.Param("id")
    if err := c.service.DeleteAccidente(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Accidente eliminado"})
}