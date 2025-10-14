package model

import "time"

type SaleRequest struct {
	ProdutoID  int64 `json:"produtoId"`
	Quantidade int32 `json:"quantidade"`
	TempoValor int64 `json:"tempoValor"`
}

type SaleResponse struct {
	ProdutoID  int64     `json:"produtoId" db:"produto_id"` //
	Quantidade int       `json:"quantidade" db:"quantidade"`
	PrecoUnit  float64   `json:"precoUnit" db:"preco_unit"`
	Total      float64   `json:"total" db:"total"`
	DataVenda  time.Time `json:"dataVenda" db:"data_venda"`
}
