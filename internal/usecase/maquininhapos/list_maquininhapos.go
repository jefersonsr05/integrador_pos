package usecase

import (
	"log"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	maquininhaposdto "github.com/jefersonsr05/integrador_pos/internal/usecase/maquininhapos/dto"
)

type ListMaquininhaPosUseCase struct {
	MaquininhaPosRepository entity.MaquininhaPosRepository
}

func NewListMaquininhaPosUseCase(maquininhaPosRepository entity.MaquininhaPosRepository) *ListMaquininhaPosUseCase {
	return &ListMaquininhaPosUseCase{MaquininhaPosRepository: maquininhaPosRepository}

}

func (u *ListMaquininhaPosUseCase) Execute() ([]*maquininhaposdto.MaquininhaPosOutputDTO, error) {
	maquininhaspos, err := u.MaquininhaPosRepository.FindAll()
	if err != nil {
		log.Printf("Erro maquininha" + err.Error())
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
