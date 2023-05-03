package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	maquininhaposdto "github.com/jefersonsr05/integrador_pos/internal/usecase/maquininhapos/dto"
)

type CreateMaquininhaPosUseCase struct {
	MaquininhaPosRepository entity.MaquininhaPosRepository
}

func NewCreateMaquininhaPosUseCase(maquininhaPosRepository entity.MaquininhaPosRepository) *CreateMaquininhaPosUseCase {
	return &CreateMaquininhaPosUseCase{MaquininhaPosRepository: maquininhaPosRepository}
}

func (u *CreateMaquininhaPosUseCase) Execute(input maquininhaposdto.MaquininhaPosInputDTO) (*maquininhaposdto.MaquininhaPosOutputDTO, error) {
	maquininhapos := entity.NewMaquininhaPos(input.EmpresaID, input.Descricao, input.Administradora, input.Cnpj)
	err := u.MaquininhaPosRepository.Create(maquininhapos)
	if err != nil {
		return nil, err
	}
	return &maquininhaposdto.MaquininhaPosOutputDTO{
		ID:             maquininhapos.ID,
		EmpresaID:      maquininhapos.EmpresaID,
		Descricao:      maquininhapos.Descricao,
		Administradora: maquininhapos.Administradora,
		Cnpj:           maquininhapos.Cnpj,
	}, nil
}
