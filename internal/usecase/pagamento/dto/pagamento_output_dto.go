package pagamentodto

type PagamentoOutputDTO struct {
	ID           string `json:"id"`
	EmpresaID    string `json:"empresa_id"`
	Descricao    string `json:"descricao"`
	TpPagamento  string `json:"tp_pagamento"`
	IndPagamento string `json:"ind_pagamento"`
	PosExclusivo bool   `json:"pos_exclusivo"`
	IDPos        string `json:"id_pos"`
}

type PagamentosOutputDTO struct {
	Pagamentos PagamentoOutputDTO `json:"pagamentos"`
}
