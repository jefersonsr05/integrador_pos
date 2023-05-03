package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	itemvendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/itemvenda/dto"
)

type CreateItemVendaUseCase struct {
	ItemVendaRepository entity.ItemVendaRepository
}

func NewCreateItemVendaUseCase(itemVendaRepository entity.ItemVendaRepository) *CreateItemVendaUseCase {
	return &CreateItemVendaUseCase{ItemVendaRepository: itemVendaRepository}
}

func (u *CreateItemVendaUseCase) Execute(input itemvendadto.ItemVendaInputDTO) (*itemvendadto.ItemVendaOutputDTO, error) {
	itemvenda := entity.NewItemVenda(input.EmpresaID, input.VendaID, input.ProdutoID, input.Quantidade, input.Valor, input.Total)
	err := u.ItemVendaRepository.Create(itemvenda)
	if err != nil {
		return nil, err
	}
	return &itemvendadto.ItemVendaOutputDTO{
		ID:         itemvenda.ID,
		EmpresaID:  itemvenda.EmpresaID,
		VendaID:    itemvenda.VendaID,
		ProdutoID:  itemvenda.ProdutoID,
		Quantidade: itemvenda.Quantidade,
		Valor:      itemvenda.Valor,
		Total:      itemvenda.Total,
	}, nil
}
