package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	vendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/venda/dto"
)

type ListVendaUseCase struct {
	VendaRepository entity.VendaRepository
}

func NewListVendaUseCase(vendaRepository entity.VendaRepository) *ListVendaUseCase {
	return &ListVendaUseCase{VendaRepository: vendaRepository}

}

func (u *ListVendaUseCase) Execute() ([]*vendadto.VendaOutputDTO, error) {
	vendas, err := u.VendaRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var vendasOutput []*vendadto.VendaOutputDTO
	for _, venda := range vendas {
		vendasOutput = append(vendasOutput, &vendadto.VendaOutputDTO{
			ID:        venda.ID,
			EmpresaID: venda.EmpresaID,
			ClienteID: venda.ClienteID,
			VendaMc:   venda.VendaMc,
			Total:     venda.Total,
			Data:      venda.Data,
		})
	}
	return vendasOutput, nil
}
