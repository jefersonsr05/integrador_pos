package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
)

type DeleteVendaUseCase struct {
	VendaRepository entity.VendaRepository
}

func NewDeleteVendaUseCase(vendaRepository entity.VendaRepository) *DeleteVendaUseCase {
	return &DeleteVendaUseCase{VendaRepository: vendaRepository}
}

func (u *DeleteVendaUseCase) Execute(id string) error {

	err := u.VendaRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
