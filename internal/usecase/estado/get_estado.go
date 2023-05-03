package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	estadodto "github.com/jefersonsr05/integrador_pos/internal/usecase/estado/dto"
)

type GetEstadoUseCase struct {
	EstadoRepository entity.EstadoRepository
}

func NewGetEstadoUseCase(estadoRepository entity.EstadoRepository) *GetEstadoUseCase {
	return &GetEstadoUseCase{EstadoRepository: estadoRepository}

}

func (u *GetEstadoUseCase) GetEstadoByID(id string) (*estadodto.EstadoOutputDTO, error) {
	estado, err := u.EstadoRepository.GetEstadoByID(id)
	if err != nil {
		return nil, err
	}
	return &estadodto.EstadoOutputDTO{
		ID:        estado.ID,
		Descricao: estado.Descricao,
		UF:        estado.UF,
		CodIbge:   estado.CodIbge,
	}, nil
}

func (u *GetEstadoUseCase) GetEstadoByUF(uf string) (*estadodto.EstadoOutputDTO, error) {
	estado, err := u.EstadoRepository.GetEstadoByUF(uf)
	if err != nil {
		return nil, err
	}
	return &estadodto.EstadoOutputDTO{
		ID:        estado.ID,
		Descricao: estado.Descricao,
		UF:        estado.UF,
		CodIbge:   estado.CodIbge,
	}, nil
}

// func (u *GetEstadoUseCase) GetEstadoByIBGE(ibge string) (*estadodto.EstadoOutputDTO, error) {
// 	estado, err := u.EstadoRepository.GetEstadoByIBGE(ibge)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &estadodto.EstadoOutputDTO{
// 		ID:        estado.ID,
// 		Descricao: estado.Descricao,
// 		UF:        estado.UF,
// 		CodIbge:   estado.CodIbge,
// 	}, nil
// }
