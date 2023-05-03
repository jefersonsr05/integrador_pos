package entity

import "github.com/google/uuid"

type ProdutoRepository interface {
	Create(produto *Produto) error
	Update(produto *Produto) error
	GetProdutoByID(id string) (*Produto, error)
	GetProdutoByEmpresa(id string) ([]*Produto, error)
	Delete(id string) error
	FindAll() ([]*Produto, error)
}

type Produto struct {
	ID        string
	EmpresaID string
	CodigoMc  string
	Descricao string
	CodBarras string
	Ncm       string
	Cest      string
	Cbenef    string
	Preco     float64
}

func NewProduto(codigo_mc string, empresaid string, descricao string, cod_barras string, ncm string, cest string, cbenef string, preco float64) *Produto {
	return &Produto{
		ID:        uuid.New().String(),
		EmpresaID: empresaid,
		CodigoMc:  codigo_mc,
		Descricao: descricao,
		CodBarras: cod_barras,
		Ncm:       ncm,
		Cest:      cest,
		Cbenef:    cbenef,
		Preco:     preco,
	}
}
