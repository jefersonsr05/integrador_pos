package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	maquininhaposdto "github.com/jefersonsr05/integrador_pos/internal/usecase/maquininhapos/dto"
)

type UpdateMaquininhaPosUseCase struct {
	MaquininhaPosRepository entity.MaquininhaPosRepository
}

func NewUpdateMaquininhaPosUseCase(maquininhaPosRepository entity.MaquininhaPosRepository) *UpdateMaquininhaPosUseCase {
	return &UpdateMaquininhaPosUseCase{MaquininhaPosRepository: maquininhaPosRepository}
}

func (u *UpdateMaquininhaPosUseCase) Execute(id string, input maquininhaposdto.MaquininhaPosInputDTO) (*maquininhaposdto.MaquininhaPosOutputDTO, error) {
	maquininhapos, err := u.MaquininhaPosRepository.GetMaquininhasPOS(id)
	if err != nil {
		return nil, err
	}

	maquininhapos.EmpresaID = input.EmpresaID
	maquininhapos.Descricao = input.Descricao
	maquininhapos.Administradora = input.Administradora
	maquininhapos.Cnpj = input.Cnpj

	err = u.MaquininhaPosRepository.Update(maquininhapos)
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
