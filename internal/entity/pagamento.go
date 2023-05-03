package entity

import "github.com/google/uuid"

type PagamentoRepository interface {
	Create(pagamento *Pagamento) error
	Update(pagamento *Pagamento) error
	GetPagamentoByID(id string) (*Pagamento, error)
	GetPagamentoByEmpresaID(id string) ([]*Pagamento, error)
	FindAll() ([]*Pagamento, error)
	Delete(id string) error
}

type Pagamento struct {
	ID           string
	EmpresaID    string
	Descricao    string
	TpPagamento  string
	IndPagamento string
	PosExclusivo bool
	IDPos        string
}

func NewPagamento(empresaid string, descricao string, tppagamento string, indpagamento string, posexclusivo bool, idpos string) *Pagamento {
	return &Pagamento{
		ID:           uuid.New().String(),
		EmpresaID:    empresaid,
		Descricao:    descricao,
		TpPagamento:  tppagamento,
		IndPagamento: indpagamento,
		PosExclusivo: posexclusivo,
		IDPos:        idpos,
	}
}
