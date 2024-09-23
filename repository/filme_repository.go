package repository

import (
	"awesomeProject/model"
	"database/sql"
	"fmt"
)

type FilmeRepository struct {
	connection *sql.DB
}

func NewFilmeRepository(connection *sql.DB) FilmeRepository {
	return FilmeRepository{
		connection: connection,
	}
}

func (fm *FilmeRepository) GetFilmes() ([]model.Filme, error) {
	query := "SELECT id,titulo,diretor,ano_lancamento,genero,duracao FROM filmes"
	rows, err := fm.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Filme{}, err
	}
	var filmeList []model.Filme
	var filmeObj model.Filme

	for rows.Next() {
		err = rows.Scan(
			&filmeObj.ID,
			&filmeObj.Titulo,
			&filmeObj.Diretor,
			&filmeObj.AnoLancamento,
			&filmeObj.Genero,
			&filmeObj.Duracao)

		if err != nil {
			fmt.Println(err)
			return []model.Filme{}, err
		}

		filmeList = append(filmeList, filmeObj)
	}
	rows.Close()

	return filmeList, nil
}

func (fm *FilmeRepository) CreateFilme(filme model.Filme) (int, error) {
	var id int
	query, err := fm.connection.Prepare("INSERT INTO filmes" +
		"(titulo,diretor,ano_lancamento,genero,duracao)" +
		"VALUES($1,$2,$3,$4,$5 RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(filme.Titulo, filme.Diretor, filme.AnoLancamento, filme.Genero, filme.Duracao).Scan(&filme.ID)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}
