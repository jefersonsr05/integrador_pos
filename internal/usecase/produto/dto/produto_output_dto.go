package produtodto

type ProdutoOutputDTO struct {
	ID        string  `json:"id"`
	EmpresaID string  `json:"empresa_id"`
	CodigoMc  string  `json:"codigo_mc"`
	Descricao string  `json:"descricao"`
	CodBarras string  `json:"codigo_barras"`
	Ncm       string  `json:"ncm"`
	Cest      string  `json:"cest"`
	Cbenef    string  `json:"cbenef"`
	Preco     float64 `json:"preco"`
}
