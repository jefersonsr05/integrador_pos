package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
)

type DeleteItemVendaUseCase struct {
	ItemVendaRepository entity.ItemVendaRepository
}

func NewDeleteItemVendaUseCase(itemVendaRepository entity.ItemVendaRepository) *DeleteItemVendaUseCase {
	return &DeleteItemVendaUseCase{ItemVendaRepository: itemVendaRepository}
}

func (u *DeleteItemVendaUseCase) Execute(id string) error {

	err := u.ItemVendaRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
