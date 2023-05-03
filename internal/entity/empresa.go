package entity

import "github.com/google/uuid"

type EmpresaRepository interface {
	Create(empresa *Empresa) error
	Update(empresa *Empresa) error
	GetEmpresaByID(id string) (*Empresa, error)
	GetEmpresaByCNPJ(cnpj string) (*Empresa, error)
	GetEmpresaByChaveRegistro(chave string) (*Empresa, error)
	FindAll() ([]*Empresa, error)
	Delete(id string) error
}

type Empresa struct {
	ID            string
	Descricao     string
	Cnpj          string
	ChaveRegistro string
}

func NewEmpresa(descricao string, cnpj string, chaveregistro string) *Empresa {
	return &Empresa{
		ID:            uuid.New().String(),
		Descricao:     descricao,
		Cnpj:          cnpj,
		ChaveRegistro: chaveregistro,
	}
}
