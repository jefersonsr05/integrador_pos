package vendadto

import "time"

type VendaInputDTO struct {
	EmpresaID string    `json:"empresa_id"`
	ClienteID string    `json:"cliente_id"`
	VendaMc   string    `json:"venda_mc"`
	Total     float64   `json:"total"`
	Data      time.Time `json:"data"`
}

type VendasInputDTO struct {
	Vendas VendaInputDTO `json:"vendas"`
}
