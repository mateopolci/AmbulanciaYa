package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/models"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

type AmbulanciaController struct {
	service *services.AmbulanciaService
}

func NewAmbulanciaController(service *services.AmbulanciaService) *AmbulanciaController {
	return &AmbulanciaController{service: service}
}

func (c *AmbulanciaController) GetAmbulancias(ctx *gin.Context) {
	ambulanciasDTO, err := c.service.GetAllAmbulancias()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ambulanciasDTO)
}

func (c *AmbulanciaController) GetAmbulanciasDesc(ctx *gin.Context) {
    ambulanciasDTO, err := c.service.GetAllAmbulanciasDesc()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, ambulanciasDTO)
}

func (c *AmbulanciaController) GetAmbulancia(ctx *gin.Context) {
	id := ctx.Param("id")
	ambulanciaDTO, err := c.service.GetAmbulanciaById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Ambulancia no encontrada"})
		return
	}
	ctx.JSON(http.StatusOK, ambulanciaDTO)
}

func (c *AmbulanciaController) GetAmbulanciaDisponible(ctx *gin.Context) {
	var descripcion string
    ambulanciaDTO, err := c.service.GetAmbulanciaDisp(descripcion)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "No hay ambulancias disponibles"})
        return
    }
    ctx.JSON(http.StatusOK, ambulanciaDTO)
}

func (c *AmbulanciaController) PostAmbulancia(ctx *gin.Context) {
	var ambulanciaDTO models.AmbulanciaDTO
	if err := ctx.ShouldBindJSON(&ambulanciaDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ambulancia, err := c.service.CreateAmbulancia(ambulanciaDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, ambulancia.AmbulanciaToDTO())
}

func (c *AmbulanciaController) PutAmbulancia(ctx *gin.Context) {
	id := ctx.Param("id")
	var ambulanciaDTO models.AmbulanciaDTO
	if err := ctx.ShouldBindJSON(&ambulanciaDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ambulancia, err := c.service.UpdateAmbulancia(id, ambulanciaDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ambulancia.AmbulanciaToDTO())
}

func (c *AmbulanciaController) DeleteAmbulancia(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteAmbulancia(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Ambulancia eliminada"})
}

func (c *AmbulanciaController) SolicitarAmbulancia(ctx *gin.Context) {
    var pedidoDTO models.AmbulanciaPedidoDTO
    if err := ctx.ShouldBindJSON(&pedidoDTO); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    mensaje, err := c.service.PedidoAmbulancia(pedidoDTO)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": mensaje})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": mensaje})
}