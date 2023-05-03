package entity

import (
	"time"

	"github.com/google/uuid"
)

type VendaRepository interface {
	Create(venda *Venda) error
	Update(venda *Venda) error
	GetVendaByID(id string) (*Venda, error)
	GetVendaByEmpresa(id string) ([]*Venda, error)
	GetVendaByStatus(status bool) ([]*Venda, error)
	FindAll() ([]*Venda, error)
	Delete(id string) error
}

type Venda struct {
	ID        string
	EmpresaID string
	ClienteID string
	VendaMc   string
	Total     float64
	Data      time.Time
}

func NewVenda(empresaID string, clienteID string, vendaMc string, total float64, data time.Time) *Venda {
	return &Venda{
		ID:        uuid.New().String(),
		EmpresaID: empresaID,
		ClienteID: clienteID,
		VendaMc:   vendaMc,
		Total:     total,
		Data:      data,
	}
}
