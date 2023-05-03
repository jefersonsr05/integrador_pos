package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	itemvendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/itemvenda/dto"
)

type ListItemVendaUseCase struct {
	ItemVendaRepository entity.ItemVendaRepository
}

func NewListItemVendaUseCase(itemVendaRepository entity.ItemVendaRepository) *ListItemVendaUseCase {
	return &ListItemVendaUseCase{ItemVendaRepository: itemVendaRepository}

}

func (u *ListItemVendaUseCase) Execute() ([]*itemvendadto.ItemVendaOutputDTO, error) {
	itensvenda, err := u.ItemVendaRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var itensvendaOutput []*itemvendadto.ItemVendaOutputDTO
	for _, itemvenda := range itensvenda {
		itensvendaOutput = append(itensvendaOutput, &itemvendadto.ItemVendaOutputDTO{
			ID:         itemvenda.ID,
			EmpresaID:  itemvenda.EmpresaID,
			VendaID:    itemvenda.VendaID,
			ProdutoID:  itemvenda.ProdutoID,
			Quantidade: itemvenda.Quantidade,
			Valor:      itemvenda.Valor,
			Total:      itemvenda.Total,
		})
	}
	return itensvendaOutput, nil
}
