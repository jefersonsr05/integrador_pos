package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
)

type DeletePagamentoVendaUseCase struct {
	PagamentoVendaRepository entity.PagamentoVendaRepository
}

func NewDeletePagamentoVendaUseCase(pagamentoVendaRepository entity.PagamentoVendaRepository) *DeletePagamentoVendaUseCase {
	return &DeletePagamentoVendaUseCase{PagamentoVendaRepository: pagamentoVendaRepository}
}

func (u *DeletePagamentoVendaUseCase) Execute(id string) error {

	err := u.PagamentoVendaRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
