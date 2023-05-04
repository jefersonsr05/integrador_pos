package itemvendadto

type ItemVendaOutputDTO struct {
	ID         string  `json:"id"`
	EmpresaID  string  `json:"empresa_id"`
	VendaID    string  `json:"venda_id"`
	ProdutoID  string  `json:"produto_id"`
	Quantidade float64 `json:"quantidade"`
	Valor      float64 `json:"valor"`
	Total      float64 `json:"total"`
}

type ItensOutputDTO struct {
	ItensVenda ItemVendaOutputDTO `json:"itens_venda"`
}
