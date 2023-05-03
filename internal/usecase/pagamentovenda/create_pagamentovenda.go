package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	pagamentovendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/pagamentovenda/dto"
)

type CreatePagamentoVendaUseCase struct {
	PagamentoVendaRepository entity.PagamentoVendaRepository
}

func NewCreatePagamentoVendaUseCase(pagamentoVendaRepository entity.PagamentoVendaRepository) *CreatePagamentoVendaUseCase {
	return &CreatePagamentoVendaUseCase{PagamentoVendaRepository: pagamentoVendaRepository}
}

func (u *CreatePagamentoVendaUseCase) Execute(input pagamentovendadto.PagamentoVendaInputDTO) (*pagamentovendadto.PagamentoVendaOutputDTO, error) {
	pagamentovenda := entity.NewPagamentoVenda(input.EmpresaID, input.VendaID, input.PagamentoID, input.Valor, input.Vencimento, input.Status)
	err := u.PagamentoVendaRepository.Create(pagamentovenda)
	if err != nil {
		return nil, err
	}
	return &pagamentovendadto.PagamentoVendaOutputDTO{
		ID:          pagamentovenda.ID,
		EmpresaID:   pagamentovenda.EmpresaID,
		VendaID:     pagamentovenda.VendaID,
		PagamentoID: pagamentovenda.PagamentoID,
		Valor:       pagamentovenda.Valor,
		Vencimento:  pagamentovenda.Vencimento,
		Status:      pagamentovenda.Status,
	}, nil
}
