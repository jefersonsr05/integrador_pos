package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
)

type DeleteEstadoUseCase struct {
	EstadoRepository entity.EstadoRepository
}

func NewDeleteEstadoUseCase(estadoRepository entity.EstadoRepository) *DeleteEstadoUseCase {
	return &DeleteEstadoUseCase{EstadoRepository: estadoRepository}
}

func (u *DeleteEstadoUseCase) Execute(id string) error {

	err := u.EstadoRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
