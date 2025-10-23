package model

type ProdutosRequest struct {
	IdRoupa     int64  `json:"id_roupa"`
	Categoria   string `json:"categoria"`
	Tamanho     string `json:"tamanho"`
	Cores       string `json:"cores"`
	TempoValor  int64  `json:"tempoValor"`
	Status      string `json:"status"`
	Localizacao string `json:"localizacao"`
	ImagemUrl   string `json:"imagem_url"`
}

type ProdutosResponse struct {
	IdRoupa     int64  `json:"id_roupa"`
	Nome        string `json:"nome"`
	Categoria   string `json:"categoria"`
	Tamanho     string `json:"tamanho"`
	Cores       string `json:"cores"`
	TempoValor  int64  `json:"tempoValor"`
	Status      string `json:"status"`
	Localizacao string `json:"localizacao"`
	ImagemUrl   string `json:"imagem_url"`
	Ativo       bool   `json:"ativo"`
}

type GetProdutosByIdResponse struct {
	Id          int64  `json:"id_roupa"`
	Categoria   string `json:"categoria"`
	Tamanho     string `json:"tamanho"`
	Cores       string `json:"cores"`
	TempoValor  int64  `json:"tempoValor"`
	Status      string `json:"status"`
	Localizacao string `json:"localizacao"`
	ImagemUrl   string `json:"imagem_url"`
	Ativo       bool   `json:"ativo"`
}
