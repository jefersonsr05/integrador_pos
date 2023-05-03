package entity

import "github.com/google/uuid"

type CidadeRepository interface {
	Create(cidade *Cidade) error
	Update(cidade *Cidade) error
	GetCidadeByID(id string) (*Cidade, error)
	GetCidadeByIBGE(ibgecidade string) (*Cidade, error)
	GetCidadeByIbgeUF(ibgeuf int32) ([]*Cidade, error)
	GetCidadeByUF(uf string) ([]*Cidade, error)
	FindAll() ([]*Cidade, error)
	Delete(id string) error
}

type Cidade struct {
	ID        string
	Descricao string
	EstadoID  string
	CodIbge   string
}

func NewCidade(descricao string, estadoid string, codigoibge string) *Cidade {
	return &Cidade{
		ID:        uuid.New().String(),
		Descricao: descricao,
		EstadoID:  estadoid,
		CodIbge:   codigoibge,
	}
}
