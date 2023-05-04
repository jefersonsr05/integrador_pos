package entity

import "github.com/google/uuid"

type ClienteRepository interface {
	Create(cliente *Cliente) error
	Update(cliente *Cliente) error
	GetClienteByID(id string) (*Cliente, error)
	GetClienteByEmpresa(empresaid string) ([]*Cliente, error)
	GetClienteByCodigoMC(codigomc int64) (*Cliente, error)
	FindAll() ([]*Cliente, error)
	Delete(id string) error
}

type Cliente struct {
	ID          string
	EmpresaID   string
	CodigoMc    int64
	Nome        string
	Cep         string
	CidadeID    string
	Endereco    string
	Numero      string
	Complemento string
}

func NewCliente(empresaid string, codigomc int64, nome string, cep string, cidadeid string, endereco string, numero string, complemento string) *Cliente {
	return &Cliente{
		ID:          uuid.New().String(),
		EmpresaID:   empresaid,
		CodigoMc:    codigomc,
		Nome:        nome,
		Cep:         cep,
		CidadeID:    cidadeid,
		Endereco:    endereco,
		Numero:      numero,
		Complemento: complemento,
	}
}
