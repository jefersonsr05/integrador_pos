package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	estadodto "github.com/jefersonsr05/integrador_pos/internal/usecase/estado/dto"
)

type ListEstadoUseCase struct {
	EstadoRepository entity.EstadoRepository
}

func NewListEstadoUseCase(estadoRepository entity.EstadoRepository) *ListEstadoUseCase {
	return &ListEstadoUseCase{EstadoRepository: estadoRepository}

}

func (u *ListEstadoUseCase) Execute() ([]*estadodto.EstadoOutputDTO, error) {
	estados, err := u.EstadoRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var estadosOutput []*estadodto.EstadoOutputDTO
	for _, estado := range estados {
		estadosOutput = append(estadosOutput, &estadodto.EstadoOutputDTO{
			ID:        estado.ID,
			Descricao: estado.Descricao,
			UF:        estado.UF,
			CodIbge:   estado.CodIbge,
		})
	}
	return estadosOutput, nil
}
