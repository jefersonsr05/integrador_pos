package entity

import "github.com/google/uuid"

type MaquininhaPosRepository interface {
	Create(maquininhapos *MaquininhaPos) error
	Update(maquininhapos *MaquininhaPos) error
	GetMaquininhasPOS(id string) (*MaquininhaPos, error)
	GetMaquininhasPosByEmpresa(id string) ([]*MaquininhaPos, error)
	Delete(id string) error
	FindAll() ([]*MaquininhaPos, error)
}

type MaquininhaPos struct {
	ID             string
	EmpresaID      string
	Descricao      string
	Administradora string
	Cnpj           string
}

func NewMaquininhaPos(empresaid string, descricao string, administradora string, cnpj string) *MaquininhaPos {
	return &MaquininhaPos{
		ID:             uuid.New().String(),
		EmpresaID:      empresaid,
		Descricao:      descricao,
		Administradora: administradora,
		Cnpj:           cnpj,
	}
}
