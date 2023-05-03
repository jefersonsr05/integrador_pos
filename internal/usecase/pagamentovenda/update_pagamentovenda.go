package usecase

import (
	"log"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	pagamentovendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/pagamentovenda/dto"
)

type UpdatePagamentoVendaUseCase struct {
	PagamentoVendaRepository entity.PagamentoVendaRepository
}

func NewUpdatePagamentoVendaUseCase(pagamentoVendaRepository entity.PagamentoVendaRepository) *UpdatePagamentoVendaUseCase {
	return &UpdatePagamentoVendaUseCase{PagamentoVendaRepository: pagamentoVendaRepository}
}

func (u *UpdatePagamentoVendaUseCase) Execute(id string, input pagamentovendadto.PagamentoVendaInputDTO) (*pagamentovendadto.PagamentoVendaOutputDTO, error) {
	pagamentovenda, err := u.PagamentoVendaRepository.GetPagamentoVendaByID(id)
	if err != nil {
		log.Printf("não achei resultado")
		return nil, err
	}

	pagamentovenda.EmpresaID = input.EmpresaID
	pagamentovenda.VendaID = input.VendaID
	pagamentovenda.PagamentoID = input.PagamentoID
	pagamentovenda.Valor = input.Valor
	pagamentovenda.Vencimento = input.Vencimento
	pagamentovenda.Status = input.Status

	err = u.PagamentoVendaRepository.Update(pagamentovenda)
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
