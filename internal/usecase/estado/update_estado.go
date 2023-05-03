package usecase

import (
	"log"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	estadodto "github.com/jefersonsr05/integrador_pos/internal/usecase/estado/dto"
)

type UpdateEstadoUseCase struct {
	EstadoRepository entity.EstadoRepository
}

func NewUpdateEstadoUseCase(estadoRepository entity.EstadoRepository) *UpdateEstadoUseCase {
	return &UpdateEstadoUseCase{EstadoRepository: estadoRepository}
}

func (u *UpdateEstadoUseCase) Execute(id string, input estadodto.EstadoInputDTO) (*estadodto.EstadoOutputDTO, error) {
	estado, err := u.EstadoRepository.GetEstadoByID(id)
	if err != nil {
		log.Printf("não achei resultado")
		return nil, err
	}

	estado.Descricao = input.Descricao
	estado.UF = input.UF
	estado.CodIbge = input.CodIbge

	err = u.EstadoRepository.Update(estado)
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
