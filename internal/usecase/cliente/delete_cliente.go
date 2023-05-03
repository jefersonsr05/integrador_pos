package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
)

type DeleteClienteUseCase struct {
	ClienteRepository entity.ClienteRepository
}

func NewDeleteClienteUseCase(clienteRepository entity.ClienteRepository) *DeleteClienteUseCase {
	return &DeleteClienteUseCase{ClienteRepository: clienteRepository}
}

func (u *DeleteClienteUseCase) Execute(id string) error {

	err := u.ClienteRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
