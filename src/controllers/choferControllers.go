package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/models"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

type ChoferController struct {
	service *services.ChoferService
}

func NewChoferController(service *services.ChoferService) *ChoferController {
	return &ChoferController{service: service}
}

func (c *ChoferController) GetChoferes(ctx *gin.Context) {
	choferDTO, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, choferDTO)
}

func (c *ChoferController) GetChofer(ctx *gin.Context) {
	id := ctx.Param("id")
	choferDTO, err := c.service.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Chofer no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, choferDTO)
}

func (c *ChoferController) PostChofer(ctx *gin.Context) {
	var choferDTO models.ChoferDTO
	if err := ctx.ShouldBindJSON(&choferDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chofer, err := c.service.Create(choferDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, chofer.ChoferToDTO())
}

func (c *ChoferController) PutChofer(ctx *gin.Context) {
	id := ctx.Param("id")
	var choferDTO models.ChoferDTO
	if err := ctx.ShouldBindJSON(&choferDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chofer, err := c.service.Update(id, choferDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, chofer.ChoferToDTO())
}

func (c *ChoferController) DeleteChofer(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Chofer eliminado"})
}
