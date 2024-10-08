package repository

import (
	"awesomeProject/model"
	"database/sql"
	"log"
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
	query := "SELECT id, titulo, diretor, ano_lancamento, genero, duracao FROM filmes"
	rows, err := fm.connection.Query(query)
	if err != nil {
		log.Println("Erro ao executar query GetFilmes:", err)
		return nil, err // Retorna nil e o erro para a camada superior
	}
	defer rows.Close() // Garante que rows seja fechado no final

	var filmeList []model.Filme
	for rows.Next() {
		var filmeObj model.Filme
		if err := rows.Scan(
			&filmeObj.ID,
			&filmeObj.Titulo,
			&filmeObj.Diretor,
			&filmeObj.AnoLancamento,
			&filmeObj.Genero,
			&filmeObj.Duracao,
		); err != nil {
			log.Println("Erro ao escanear filme:", err)
			return nil, err
		}
		filmeList = append(filmeList, filmeObj)
	}
	if err := rows.Err(); err != nil {
		log.Println("Erro ao iterar sobre as linhas:", err)
		return nil, err
	}

	return filmeList, nil
}

func (fm *FilmeRepository) CreateFilme(filme model.Filme) (int, error) {
	query, err := fm.connection.Prepare("INSERT INTO filmes (titulo, diretor, ano_lancamento, genero, duracao) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		log.Println("Erro ao preparar query CreateFilme:", err)
		return 0, err
	}
	defer query.Close() // Garante que query seja fechado no final

	err = query.QueryRow(filme.Titulo, filme.Diretor, filme.AnoLancamento, filme.Genero, filme.Duracao).Scan(&filme.ID)
	if err != nil {
		log.Println("Erro ao executar query CreateFilme:", err)
		return 0, err
	}

	return filme.ID, nil
}

func (fm *FilmeRepository) GetFilmesById(id_filme int) (*model.Filme, error) {
	query, err := fm.connection.Prepare("SELECT * FROM filmes WHERE id = $1")
	if err != nil {
		log.Println("Erro ao preparar query GetFilmesById:", err)
		return nil, err
	}
	defer query.Close() // Garante que query seja fechado no final

	var filme model.Filme
	err = query.QueryRow(id_filme).Scan(
		&filme.ID,
		&filme.Titulo,
		&filme.Diretor,
		&filme.AnoLancamento,
		&filme.Genero,
		&filme.Duracao,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Nenhum filme encontrado para o ID:", id_filme)
			return nil, nil // Retorna nil para o filme e nil para o erro
		}
		log.Println("Erro ao buscar filme por ID:", err)
		return nil, err
	}

	return &filme, nil
}

func (fm *FilmeRepository) AtualizarFilme(filme model.Filme) error {
	query, err := fm.connection.Prepare("UPDATE filmes SET titulo = $1, diretor = $2, ano_lancamento = $3, genero = $4, duracao = $5 WHERE id = $6")
	if err != nil {
		log.Println("Erro ao preparar query AtualizarFilme:", err)
		return err
	}
	defer query.Close() // Garante que query seja fechado no final

	_, err = query.Exec(filme.Titulo, filme.Diretor, filme.AnoLancamento, filme.Genero, filme.Duracao, filme.ID)
	if err != nil {
		log.Println("Erro ao executar query AtualizarFilme:", err)
		return err
	}

	return nil
}

func (fm *FilmeRepository) DeletarFilme(id int) error {
	query, err := fm.connection.Prepare("DELETE FROM filmes WHERE id = $1")
	if err != nil {
		log.Println("Erro ao preparar query DeletarFilme:", err)
		return err
	}
	defer query.Close() // Garante que query seja fechado no final

	_, err = query.Exec(id)
	if err != nil {
		log.Println("Erro ao executar query DeletarFilme:", err)
		return err
	}

	return nil
}
