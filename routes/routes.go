package routes

import (
	"awesomeProject/controller"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func SetupFilmeRoutes(router *gin.RouterGroup, filmesController *controller.FilmesController) {
	authorized := router.Group("")
	authorized.Use(middleware.BasicAuth())
	{
		authorized.GET("", filmesController.GetFilmes)                // Rota: /filmes
		authorized.POST("", filmesController.CreateFilme)             // Rota: /filmes
		authorized.GET("/:filmeId", filmesController.GetFilmesById)   // Rota: /filmes/:filmeId
		authorized.PUT("/:filmeId", filmesController.AtualizarFilme)  // Rota: /filmes/:filmeId
		authorized.DELETE("/:filmeId", filmesController.DeletarFilme) // Rota: /filmes/:filmeId
	}
}
