package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	maquininhaposdto "github.com/jefersonsr05/integrador_pos/internal/usecase/maquininhapos/dto"
)

type GetMaquininhaPosUseCase struct {
	MaquininhaPosRepository entity.MaquininhaPosRepository
}

func NewGetMaquininhaPosUseCase(maquininhaPosRepository entity.MaquininhaPosRepository) *GetMaquininhaPosUseCase {
	return &GetMaquininhaPosUseCase{MaquininhaPosRepository: maquininhaPosRepository}

}

func (u *GetMaquininhaPosUseCase) GetMaquininhaPOS(id string) (*maquininhaposdto.MaquininhaPosOutputDTO, error) {
	maquininhapos, err := u.MaquininhaPosRepository.GetMaquininhasPOS(id)
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

func (u *GetMaquininhaPosUseCase) GetMaquininhaPosByEmpresa(id string) ([]*maquininhaposdto.MaquininhaPosOutputDTO, error) {
	maquininhaspos, err := u.MaquininhaPosRepository.GetMaquininhasPosByEmpresa(id)
	if err != nil {
		return nil, err
	}
	var maquininhasposOutput []*maquininhaposdto.MaquininhaPosOutputDTO
	for _, maquininhapos := range maquininhaspos {
		maquininhasposOutput = append(maquininhasposOutput, &maquininhaposdto.MaquininhaPosOutputDTO{
			ID:             maquininhapos.ID,
			EmpresaID:      maquininhapos.EmpresaID,
			Descricao:      maquininhapos.Descricao,
			Administradora: maquininhapos.Administradora,
			Cnpj:           maquininhapos.Cnpj,
		})
	}
	return maquininhasposOutput, nil
}
