package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/models"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

type ParamedicoController struct {
	service *services.ParamedicoService
}

func NewParamedicoController(service *services.ParamedicoService) *ParamedicoController {
	return &ParamedicoController{service: service}
}

func (c *ParamedicoController) GetParamedicos(ctx *gin.Context) {
	paramedicoDTO, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, paramedicoDTO)
}

func (c *ParamedicoController) GetParamedico(ctx *gin.Context) {
	id := ctx.Param("id")
	paramedicoDTO, err := c.service.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Paramedico no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, paramedicoDTO)
}

func (c *ParamedicoController) PostParamedico(ctx *gin.Context) {
	var paramedicoDTO models.ParamedicoDTO
	if err := ctx.ShouldBindJSON(&paramedicoDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paramedico, err := c.service.Create(paramedicoDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, paramedico.ParamedicoToDTO())
}

func (c *ParamedicoController) PutParamedico(ctx *gin.Context) {
	id := ctx.Param("id")
	var paramedicoDTO models.ParamedicoDTO
	if err := ctx.ShouldBindJSON(&paramedicoDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paramedico, err := c.service.Update(id, paramedicoDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, paramedico.ParamedicoToDTO())
}

func (c *ParamedicoController) DeleteParamedico(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Paramedico eliminado"})
}

func (c *ParamedicoController) Login(ctx *gin.Context) {
	var loginReq models.LoginRequest

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.service.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	// Set HTTP-only cookie
    cfg := &http.Cookie{
        Name:     "jwt",
        Value:    response.Token,
        MaxAge:   3600 * 24,
        Path:     "/",
        Domain:   "",
        Secure:   true,
        HttpOnly: true,
        SameSite: http.SameSiteNoneMode,
    }
    http.SetCookie(ctx.Writer, cfg)

	// Only return isAdmin status
	ctx.JSON(http.StatusOK, gin.H{"isAdmin": response.IsAdmin})
}

func (c *ParamedicoController) Logout(ctx *gin.Context) {
    cfg := &http.Cookie{
        Name:     "jwt",
        Value:    "",
        MaxAge:   -1,
        Path:     "/",
        Domain:   "",
        Secure:   true,
        HttpOnly: true,
        SameSite: http.SameSiteNoneMode,
    }
    http.SetCookie(ctx.Writer, cfg)
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
