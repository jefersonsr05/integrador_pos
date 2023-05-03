package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	pagamentovendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/pagamentovenda/dto"
)

type ListPagamentoVendaUseCase struct {
	PagamentoVendaRepository entity.PagamentoVendaRepository
}

func NewListPagamentoVendaUseCase(pagamentoVendaRepository entity.PagamentoVendaRepository) *ListPagamentoVendaUseCase {
	return &ListPagamentoVendaUseCase{PagamentoVendaRepository: pagamentoVendaRepository}

}

func (u *ListPagamentoVendaUseCase) Execute() ([]*pagamentovendadto.PagamentoVendaOutputDTO, error) {
	pagamentosvenda, err := u.PagamentoVendaRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var pagamentosvendaOutput []*pagamentovendadto.PagamentoVendaOutputDTO
	for _, pagamentovenda := range pagamentosvenda {
		pagamentosvendaOutput = append(pagamentosvendaOutput, &pagamentovendadto.PagamentoVendaOutputDTO{
			ID:          pagamentovenda.ID,
			EmpresaID:   pagamentovenda.EmpresaID,
			VendaID:     pagamentovenda.VendaID,
			PagamentoID: pagamentovenda.PagamentoID,
			Valor:       pagamentovenda.Valor,
			Vencimento:  pagamentovenda.Vencimento,
			Status:      pagamentovenda.Status,
		})
	}
	return pagamentosvendaOutput, nil
}
