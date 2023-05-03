package usecase

import (
	"log"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	estadodto "github.com/jefersonsr05/integrador_pos/internal/usecase/estado/dto"
)

type CreateEstadoUseCase struct {
	EstadoRepository entity.EstadoRepository
}

func NewCreateEstadoUseCase(estadoRepository entity.EstadoRepository) *CreateEstadoUseCase {
	return &CreateEstadoUseCase{EstadoRepository: estadoRepository}
}

func (u *CreateEstadoUseCase) Execute(input estadodto.EstadoInputDTO) (*estadodto.EstadoOutputDTO, error) {
	log.Printf("Acessando o Create")
	estado := entity.NewEstado(input.Descricao, input.UF, input.CodIbge)
	log.Printf("Criando Objeto: " + estado.Descricao)
	err := u.EstadoRepository.Create(estado)
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
