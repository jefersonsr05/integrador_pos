package entity

import "github.com/google/uuid"

type EstadoRepository interface {
	Create(estado *Estado) error
	Update(estado *Estado) error
	GetEstadoByID(id string) (*Estado, error)
	GetEstadoByUF(uf string) (*Estado, error)
	GetEstadoByCodIBGE(ibge int32) (*Estado, error)
	FindAll() ([]*Estado, error)
	Delete(id string) error
}

type Estado struct {
	ID        string
	Descricao string
	UF        string
	CodIbge   int32
	Cidades   []Cidade
}

func NewEstado(descricao string, uf string, codigoIbge int32) *Estado {
	return &Estado{
		ID:        uuid.New().String(),
		Descricao: descricao,
		UF:        uf,
		CodIbge:   codigoIbge,
	}
}
