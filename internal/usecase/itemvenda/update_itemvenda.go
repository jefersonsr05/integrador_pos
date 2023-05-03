package usecase

import (
	"log"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	itemvendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/itemvenda/dto"
)

type UpdateItemVendaUseCase struct {
	ItemVendaRepository entity.ItemVendaRepository
}

func NewUpdateItemVendaUseCase(itemVendaRepository entity.ItemVendaRepository) *UpdateItemVendaUseCase {
	return &UpdateItemVendaUseCase{ItemVendaRepository: itemVendaRepository}
}

func (u *UpdateItemVendaUseCase) Execute(id string, input itemvendadto.ItemVendaInputDTO) (*itemvendadto.ItemVendaOutputDTO, error) {
	itemvenda, err := u.ItemVendaRepository.GetItemVendaByID(id)
	if err != nil {
		log.Printf("não achei resultado")
		return nil, err
	}

	itemvenda.EmpresaID = input.EmpresaID
	itemvenda.VendaID = input.VendaID
	itemvenda.ProdutoID = input.ProdutoID
	itemvenda.Quantidade = input.Quantidade
	itemvenda.Valor = input.Valor
	itemvenda.Total = input.Total

	err = u.ItemVendaRepository.Update(itemvenda)
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
