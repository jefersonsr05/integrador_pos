package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	itemvendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/itemvenda/dto"
)

type GetItemVendaUseCase struct {
	ItemVendaRepository entity.ItemVendaRepository
}

func NewGetItemVendaUseCase(itemVendaRepository entity.ItemVendaRepository) *GetItemVendaUseCase {
	return &GetItemVendaUseCase{ItemVendaRepository: itemVendaRepository}

}

func (u *GetItemVendaUseCase) GetItemVendaByID(id string) (*itemvendadto.ItemVendaOutputDTO, error) {
	itemvenda, err := u.ItemVendaRepository.GetItemVendaByID(id)
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

func (u *GetItemVendaUseCase) GetItemVendaByVendaID(id string) ([]*itemvendadto.ItemVendaOutputDTO, error) {
	itensvenda, err := u.ItemVendaRepository.GetItemVendaByVendaID(id)
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
