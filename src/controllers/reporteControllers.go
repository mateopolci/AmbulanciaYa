package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/mateopolci/AmbulanciaYa/src/services"
    "github.com/mateopolci/AmbulanciaYa/src/models"
)

type ReporteController struct {
    service *services.ReporteService
}

func NewReporteController(service *services.ReporteService) *ReporteController {
	return &ReporteController{service: service}
}

func (c *ReporteController) GetReportes(ctx *gin.Context) {
    reportesDTO, err := c.service.GetAllReportes()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, reportesDTO)
}

func (c *ReporteController) GetReporte(ctx *gin.Context) {
    id := ctx.Param("id")
    reporteDTO, err := c.service.GetReporteById(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Reporte no encontrado"})
        return
    }
    ctx.JSON(http.StatusOK, reporteDTO)
}

func (c *ReporteController) GetReporteByAccidente(ctx *gin.Context) {
    accidenteId := ctx.Param("accidenteId")
    reporteDTO, err := c.service.GetReporteByAccidenteId(accidenteId)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Reporte no encontrado"})
        return
    }
    ctx.JSON(http.StatusOK, reporteDTO)
}

func (c *ReporteController) PostReporte(ctx *gin.Context) {
	var reporteDTO models.ReporteDTO
	if err := ctx.ShouldBindJSON(&reporteDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reporte, err := c.service.CreateReporte(reporteDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, reporte.ReporteToDTO())
}

func (c *ReporteController) CreateReporteAndUpdateHospital(ctx *gin.Context) {
    accidenteId := ctx.Param("accidenteId")
    var postDTO models.ReportePostDTO
    
    if err := ctx.ShouldBindJSON(&postDTO); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    reporte, err := c.service.CreateReporteAndUpdateHospital(accidenteId, postDTO)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, reporte.ReporteToDTO())
}

func (c *ReporteController) PutReporte(ctx *gin.Context) {
	id := ctx.Param("id")
	var reporteDTO models.ReporteDTO
	if err := ctx.ShouldBindJSON(&reporteDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reporte, err := c.service.UpdateReporte(id, reporteDTO)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, reporte.ReporteToDTO())
}

func (c *ReporteController) UpdateReporteAndHospital(ctx *gin.Context) {
    id := ctx.Param("reporteId")
    var updateDTO models.ReporteUpdateDTO
    
    if err := ctx.ShouldBindJSON(&updateDTO); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    reporte, err := c.service.UpdateReporteAndHospital(id, updateDTO)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, reporte.ReporteToDTO())
}

func (c *ReporteController) DeleteReporte(ctx *gin.Context) {
    id := ctx.Param("id")
    if err := c.service.DeleteReporte(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Reporte eliminado"})
}

func (c *ReporteController) DeleteReporteByAccidente(ctx *gin.Context) {
    accidenteId := ctx.Param("accidenteId")
    if err := c.service.DeleteReporteByAccidenteId(accidenteId); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Reporte eliminado"})
}