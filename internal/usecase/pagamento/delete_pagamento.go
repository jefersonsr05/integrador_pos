package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
)

type DeletePagamentoUseCase struct {
	PagamentoRepository entity.PagamentoRepository
}

func NewDeletePagamentoUseCase(pagamentoRepository entity.PagamentoRepository) *DeletePagamentoUseCase {
	return &DeletePagamentoUseCase{PagamentoRepository: pagamentoRepository}
}

func (u *DeletePagamentoUseCase) Execute(id string) error {

	err := u.PagamentoRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
