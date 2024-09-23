package model

type Filme struct {
	ID            int    `json:"id_filme"`       // Identificador único do filme
	Titulo        string `json:"titulo"`         // Título do filme
	Diretor       string `json:"diretor"`        // Nome do diretor
	AnoLancamento int    `json:"ano_lancamento"` // Ano de lançamento
	Genero        string `json:"genero"`         // Gênero do filme
	Duracao       int    `json:"duracao"`        // Duração em minutos
}
