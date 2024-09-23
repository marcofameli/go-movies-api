package controller

import (
	"awesomeProject/model"
	"awesomeProject/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type filmesController struct {
	filmeUseCase usecase.FilmeUseCase
}

func NewFilmeController(usecase usecase.FilmeUseCase) filmesController {
	return filmesController{
		filmeUseCase: usecase,
	}
}

func (f *filmesController) GetFilmes(ctx *gin.Context) {
	filmes, err := f.filmeUseCase.GetFilmes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
	}
	ctx.JSON(http.StatusOK, filmes)
}

func (f *filmesController) CreateFilme(ctx *gin.Context) {
	var filme model.Filme
	err := ctx.BindJSON(&filme)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	insertedFilme, err := f.filmeUseCase.CreateFilme(filme)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusCreated, insertedFilme)
}
