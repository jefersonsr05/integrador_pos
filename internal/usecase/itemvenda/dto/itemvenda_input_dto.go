package itemvendadto

type ItemVendaInputDTO struct {
	EmpresaID  string  `json:"empresa_id"`
	VendaID    string  `json:"venda_id"`
	ProdutoID  string  `json:"produto_id"`
	Quantidade float64 `json:"quantidade"`
	Valor      float64 `json:"valor"`
	Total      float64 `json:"total"`
}

type ItensVendaInputDTO struct {
	ItensVenda ItemVendaInputDTO `json:"itens_venda"`
}
