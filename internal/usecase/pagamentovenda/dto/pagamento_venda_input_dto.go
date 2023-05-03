package pagamentovenda

import "time"

type PagamentoVendaInputDTO struct {
	EmpresaID   string    `json:"empresa_id"`
	VendaID     string    `json:"venda_id"`
	PagamentoID string    `json:"pagamento_id"`
	Valor       float64   `json:"valor"`
	Vencimento  time.Time `json:"vencimento"`
	Status      bool      `json:"status"`
}
