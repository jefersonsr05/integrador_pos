package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	pagamentovendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/pagamentovenda/dto"
)

type GetPagamentoVendaUseCase struct {
	PagamentoVendaRepository entity.PagamentoVendaRepository
}

func NewGetPagamentoVendaUseCase(pagamentoVendaRepository entity.PagamentoVendaRepository) *GetPagamentoVendaUseCase {
	return &GetPagamentoVendaUseCase{PagamentoVendaRepository: pagamentoVendaRepository}

}

func (u *GetPagamentoVendaUseCase) GetPagamentoVendaByID(id string) (*pagamentovendadto.PagamentoVendaOutputDTO, error) {
	pagamentovenda, err := u.PagamentoVendaRepository.GetPagamentoVendaByID(id)
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

func (u *GetPagamentoVendaUseCase) GetPagamentoVendaByVendaID(id string) ([]*pagamentovendadto.PagamentoVendaOutputDTO, error) {
	pagamentosvenda, err := u.PagamentoVendaRepository.GetPagamentoVendaByVendaID(id)
	if err != nil {
		return nil, err
	}
	var pagamentovendaOutput []*pagamentovendadto.PagamentoVendaOutputDTO
	for _, pagamentovenda := range pagamentosvenda {
		pagamentovendaOutput = append(pagamentovendaOutput, &pagamentovendadto.PagamentoVendaOutputDTO{
			ID:          pagamentovenda.ID,
			EmpresaID:   pagamentovenda.EmpresaID,
			VendaID:     pagamentovenda.VendaID,
			PagamentoID: pagamentovenda.PagamentoID,
			Valor:       pagamentovenda.Valor,
			Vencimento:  pagamentovenda.Vencimento,
			Status:      pagamentovenda.Status,
		})
	}
	return pagamentovendaOutput, nil
}

func (u *GetPagamentoVendaUseCase) GetPagamentoVendaByStatus(status bool) ([]*pagamentovendadto.PagamentoVendaOutputDTO, error) {
	pagamentosvenda, err := u.PagamentoVendaRepository.GetPagamentoVendaByStatus(status)
	if err != nil {
		return nil, err
	}
	var pagamentovendaOutput []*pagamentovendadto.PagamentoVendaOutputDTO
	for _, pagamentovenda := range pagamentosvenda {
		pagamentovendaOutput = append(pagamentovendaOutput, &pagamentovendadto.PagamentoVendaOutputDTO{
			ID:          pagamentovenda.ID,
			EmpresaID:   pagamentovenda.EmpresaID,
			VendaID:     pagamentovenda.VendaID,
			PagamentoID: pagamentovenda.PagamentoID,
			Valor:       pagamentovenda.Valor,
			Vencimento:  pagamentovenda.Vencimento,
			Status:      pagamentovenda.Status,
		})
	}
	return pagamentovendaOutput, nil
}
