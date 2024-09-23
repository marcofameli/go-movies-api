package usecase

import (
	"awesomeProject/model"
	"awesomeProject/repository"
)

type FilmeUseCase struct {
	repository repository.FilmeRepository
}

func NewFilmeUseCase(repo repository.FilmeRepository) FilmeUseCase {
	return FilmeUseCase{
		repository: repo,
	}
}

func (pu *FilmeUseCase) GetFilmes() ([]model.Filme, error) {
	return pu.repository.GetFilmes()
}

func (pu *FilmeUseCase) CreateFilme(filme model.Filme) (model.Filme, error) {
	filmeId, err := pu.repository.CreateFilme(filme)
	if err != nil {
		return model.Filme{}, err
	}

	filme.ID = filmeId
	return filme, nil
}
