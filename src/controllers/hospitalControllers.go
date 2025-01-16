package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/mateopolci/AmbulanciaYa/src/services"
    "github.com/mateopolci/AmbulanciaYa/src/models"
)

type HospitalController struct {
	service *services.HospitalService
}

func NewHospitalController(service *services.HospitalService) *HospitalController {
	return &HospitalController{service: service}
}

func (c *HospitalController) GetHospitales(ctx *gin.Context) {
	hospitalesDTO, err := c.service.GetAllHospitales()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, hospitalesDTO)
}

func (c *HospitalController) GetHospital(ctx *gin.Context) {
	id := ctx.Param("id")
	hospitalDTO, err := c.service.GetHospitalById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Hospital no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, hospitalDTO)
}

func (c *HospitalController) PostHospital(ctx *gin.Context) {
	var hospitalDTO models.HospitalDTO
	if err := ctx.ShouldBindJSON(&hospitalDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hospital, err := c.service.CreateHospital(hospitalDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, hospital)
}

func (c *HospitalController) PutHospital(ctx *gin.Context) {
	id := ctx.Param("id")
	var hospitalDTO models.HospitalDTO
	if err := ctx.ShouldBindJSON(&hospitalDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hospital, err := c.service.UpdateHospital(id, hospitalDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, hospital)
}

func (c *HospitalController) DeleteHospital(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteHospital(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hospital eliminado"})
}