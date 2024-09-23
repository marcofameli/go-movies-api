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

func (pu *FilmeUseCase) CreateFilme(filmes model.Filme) (model.Filme, error) {
	filmeId, err := pu.repository.CreateFilme(filmes)
	if err != nil {
		return model.Filme{}, err
	}

	filmes.ID = filmeId
	return filmes, nil
}

func (pu *FilmeUseCase) GetFilmesById(id_filme int) (*model.Filme, error) {
	filmes, err := pu.repository.GetFilmesById(id_filme)
	if err != nil {
		return nil, err
	}
	return filmes, nil
}

func (pu *FilmeUseCase) AtualizarFilme(filme model.Filme) (model.Filme, error) {

	err := pu.repository.AtualizarFilme(filme)
	if err != nil {
		return model.Filme{}, err
	}
	return filme, nil
}

func (pu *FilmeUseCase) DeletarFilme(id int) error {
	return pu.repository.DeletarFilme(id)
}
