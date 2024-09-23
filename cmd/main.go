package main

import (
	"awesomeProject/controller"
	"awesomeProject/db"
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

	server.GET("/filmes", FilmeController.GetFilmes)
	server.POST("/filme", FilmeController.CreateFilme) // Testar requisição de criar

	server.Run(":8080")

}
