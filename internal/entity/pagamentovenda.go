package entity

import (
	"time"

	"github.com/google/uuid"
)

type PagamentoVendaRepository interface {
	Create(pagamentosvenda *PagamentoVenda) error
	Update(pagamentosvenda *PagamentoVenda) error
	GetPagamentoVendaByID(id string) (*PagamentoVenda, error)
	GetPagamentoVendaByVendaID(vendaid string) ([]*PagamentoVenda, error)
	GetPagamentoVendaByStatus(status bool) ([]*PagamentoVenda, error)
	FindAll() ([]*PagamentoVenda, error)
	Delete(id string) error
}

type PagamentoVenda struct {
	ID          string
	EmpresaID   string
	VendaID     string
	PagamentoID string
	Valor       float64
	Vencimento  time.Time
	Status      bool
}

func NewPagamentoVenda(empresaid string, vendaid string, pagamentoid string, valor float64, vencimento time.Time, status bool) *PagamentoVenda {
	return &PagamentoVenda{
		ID:          uuid.New().String(),
		EmpresaID:   empresaid,
		VendaID:     vendaid,
		PagamentoID: pagamentoid,
		Valor:       valor,
		Vencimento:  vencimento,
		Status:      status,
	}
}
