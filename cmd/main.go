package main

import (
	"awesomeProject/controller"
	"awesomeProject/db"
	"awesomeProject/middleware"
	"awesomeProject/repository"
	"awesomeProject/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	//Camada de Repository
	FilmeRepository := repository.NewFilmeRepository(dbConnection)
	//Camada useCase
	FilmeUseCase := usecase.NewFilmeUseCase(FilmeRepository)
	//Camada de controllers
	FilmeController := controller.NewFilmeController(FilmeUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authorized := server.Group("/filmes")
	authorized.Use(middleware.BasicAuth())
	{

		authorized.GET("", FilmeController.GetFilmes)                // Rota: /filmes
		authorized.POST("", FilmeController.CreateFilme)             // Rota: /filmes
		authorized.GET("/:filmeId", FilmeController.GetFilmesById)   // Rota: /filmes/:filmeId
		authorized.PUT("/:filmeId", FilmeController.AtualizarFilme)  // Rota: /filmes/:filmeId
		authorized.DELETE("/:filmeId", FilmeController.DeletarFilme) // Rota: /filmes/:filmeId
	}
	server.Run(":8080")

}
