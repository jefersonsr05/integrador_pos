package pagamentovenda

import "time"

type PagamentoVendaOutputDTO struct {
	ID          string    `json:"id"`
	EmpresaID   string    `json:"empresa_id"`
	VendaID     string    `json:"venda_id"`
	PagamentoID string    `json:"pagamento_id"`
	Valor       float64   `json:"valor"`
	Vencimento  time.Time `json:"vencimento"`
	Status      bool      `json:"status"`
}

type PagamentosVendaOutputDTO struct {
	PagamentosVenda PagamentoVendaOutputDTO `json:"pagamentos_venda"`
}
