package model

type ProdutosRequest struct {
	Categoria   string `json:"categoria"`
	Tamanho     string `json:"tamanho"`
	Cores       string `json:"cores"`
	TempoValor  int64  `json:"tempoValor"`
	Status      string `json:"status"`
	Localizacao string `json:"localizacao"`
}

type ProdutosResponse struct {
	IdRoupa int64 `json:"id_roupa"`
}
