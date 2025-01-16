package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/mateopolci/AmbulanciaYa/src/services"
    "github.com/mateopolci/AmbulanciaYa/src/models"
)

type AccidenteController struct {
    service *services.AccidenteService
}

func NewAccidenteController(service *services.AccidenteService) *AccidenteController {
    return &AccidenteController{service: service}
}

func (c *AccidenteController) GetAccidentes(ctx *gin.Context) {
    accidentesDTO, err := c.service.GetAll()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, accidentesDTO)
}

func (c *AccidenteController) GetAccidente(ctx *gin.Context) {
    id := ctx.Param("id")
    accidenteDTO, err := c.service.GetById(id)
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

    accidente, err := c.service.Create(accidenteDTO)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, accidente)
}

func (c *AccidenteController) PutAccidente(ctx *gin.Context) {
    id := ctx.Param("id")
    var accidenteDTO models.AccidenteDTO
    if err := ctx.ShouldBindJSON(&accidenteDTO); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    accidente, err := c.service.Update(id, accidenteDTO)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, accidente)
}

func (c *AccidenteController) DeleteAccidente(ctx *gin.Context) {
    id := ctx.Param("id")
    if err := c.service.Delete(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Accidente eliminado"})
}