package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mateopolci/AmbulanciaYa/src/controllers"
	"github.com/mateopolci/AmbulanciaYa/src/middleware"
	"github.com/mateopolci/AmbulanciaYa/src/services"
)

func SetupParamedicoRoutes(router *gin.Engine, service *services.ParamedicoService) {
    ParamedicoController := controllers.NewParamedicoController(service)

    // Ruta pública de login
    router.POST("/login", ParamedicoController.Login)
    router.POST("/logout", ParamedicoController.Logout)

    // Rutas que requieren autenticación pero NO requieren ser admin
    paramedicoAuth := router.Group("/paramedicos/me")
    paramedicoAuth.Use(middleware.AuthMiddleware())
    {
        paramedicoAuth.PATCH("/email", ParamedicoController.UpdateEmail)
        paramedicoAuth.PATCH("/password", ParamedicoController.UpdatePassword)
    }

    // Rutas que requieren ser admin
    paramedicoAdmin := router.Group("/paramedicos")
    paramedicoAdmin.Use(middleware.AuthMiddleware(), middleware.IsAdminMiddleware())
    {
        paramedicoAdmin.GET("", ParamedicoController.GetParamedicos)
        paramedicoAdmin.GET("", ParamedicoController.GetParamedicosDisp)
        paramedicoAdmin.GET("/:id", ParamedicoController.GetParamedico)
        paramedicoAdmin.POST("", ParamedicoController.PostParamedico)
        paramedicoAdmin.PUT("/:id", ParamedicoController.PutParamedico)
        paramedicoAdmin.DELETE("/:id", ParamedicoController.DeleteParamedico)
    }
}