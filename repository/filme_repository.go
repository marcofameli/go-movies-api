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
	query, err := fm.connection.Prepare("INSERT INTO filmes" +
		"(titulo,diretor,ano_lancamento,genero,duracao)" +
		"VALUES($1,$2,$3,$4,$5) RETURNING id")
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
	return filme.ID, nil
}

func (fm *FilmeRepository) GetFilmesById(id_filme int) (*model.Filme, error) {
	query, err := fm.connection.Prepare("SELECT * FROM filmes WHERE id = $1")
	if err != nil {
		fmt.Println("Erro ao preparar query: ", err)
		return nil, err
	}

	var filme model.Filme
	err = query.QueryRow(id_filme).Scan(
		&filme.ID,
		&filme.Titulo,
		&filme.Diretor,
		&filme.AnoLancamento,
		&filme.Genero,
		&filme.Duracao)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Nenhum filme encontrado para o ID: ", id_filme)
			return nil, err
		}
		fmt.Println("Erro ao buscar filme", err)
		return nil, err
	}
	query.Close()
	return &filme, nil

}

func (fm *FilmeRepository) AtualizarFilme(filme model.Filme) error {
	query, err := fm.connection.Prepare("UPDATE filmes SET titulo = $1, diretor = $2, ano_lancamento = $3, genero = $4, duracao = $5 WHERE id = $6")
	if err != nil {
		fmt.Println("Erro ao preparar query", err)
		return err
	}

	_, err = query.Exec(filme.Titulo, filme.Diretor, filme.AnoLancamento, filme.Genero, filme.Duracao, filme.ID)
	if err != nil {
		fmt.Println("Erro ao executar query", err)
		return err
	}
	query.Close()
	return nil

}

func (fm *FilmeRepository) DeletarFilme(id int) error {

	query, err := fm.connection.Prepare("DELETE FROM filmes WHERE id = $1")
	if err != nil {
		fmt.Println("Erro ao preparar query", err)
		return err
	}
	_, err = query.Exec(id)
	if err != nil {
		fmt.Println("Erro ao executar query", err)
		return err
	}
	query.Close()
	return nil
}
