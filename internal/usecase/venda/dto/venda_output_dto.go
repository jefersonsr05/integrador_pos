package vendadto

import "time"

type VendaOutputDTO struct {
	ID        string    `json:"id"`
	EmpresaID string    `json:"empresa_id"`
	ClienteID string    `json:"cliente_id"`
	VendaMc   string    `json:"venda_mc"`
	Total     float64   `json:"total"`
	Data      time.Time `json:"data"`
}
