package controller

import (
	"awesomeProject/model"
	"awesomeProject/usecase"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (f *filmesController) GetFilmesById(ctx *gin.Context) {

	id := ctx.Param("filmeId")
	if id == "" {
		response := model.Response{
			Message: "ID do filme não pode ser nulo.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	filmesId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID do filme deve ser um número.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	filmes, err := f.filmeUseCase.GetFilmesById(filmesId)
	if err != nil {
		if err == sql.ErrNoRows {
			response := model.Response{
				Message: "Filme não encontrado na base de dados.",
			}
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		fmt.Println("Erro ao buscar filme no use case:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, filmes)
}

func (f *filmesController) AtualizarFilme(ctx *gin.Context) {
	id := ctx.Param("filmeId")

	if id == "" {
		response := model.Response{
			Message: "ID do filme não pode ser nulo.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	filmesId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID do filme deve ser um número.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var filme model.Filme
	err = ctx.BindJSON(&filme)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	filme.ID = filmesId // Atualiza o ID do filme
	filmeAtualizado, err := f.filmeUseCase.AtualizarFilme(filme)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, filmeAtualizado)
}

func (f *filmesController) DeletarFilme(ctx *gin.Context) {
	id := ctx.Param("filmeId")

	if id == "" {
		response := model.Response{
			Message: "ID do filme não pode ser nulo.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	filmesId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID do filme deve ser um número.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	err = f.filmeUseCase.DeletarFilme(filmesId)
	if err != nil {
		if err == sql.ErrNoRows {
			response := model.Response{
				Message: "Filme não encontrado na base de dados.",
			}
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		fmt.Println("Erro ao deletar filme no use case:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)

}
