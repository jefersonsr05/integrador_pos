package entity

import "github.com/google/uuid"

type ItemVendaRepository interface {
	Create(itemvenda *ItemVenda) error
	Update(itemvenda *ItemVenda) error
	GetItemVendaByID(id string) (*ItemVenda, error)
	GetItemVendaByVendaID(vendaid string) ([]*ItemVenda, error)
	FindAll() ([]*ItemVenda, error)
	Delete(id string) error
}

type ItemVenda struct {
	ID         string
	EmpresaID  string
	VendaID    string
	ProdutoID  string
	Quantidade float64
	Valor      float64
	Total      float64
}

func NewItemVenda(empresaid string, vendaid string, produtoid string, quantidade float64, valor float64, total float64) *ItemVenda {
	return &ItemVenda{
		ID:         uuid.New().String(),
		EmpresaID:  empresaid,
		VendaID:    vendaid,
		ProdutoID:  produtoid,
		Quantidade: quantidade,
		Valor:      valor,
		Total:      total,
	}
}
