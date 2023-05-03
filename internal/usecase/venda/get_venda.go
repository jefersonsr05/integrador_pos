package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	vendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/venda/dto"
)

type GetVendaUseCase struct {
	VendaRepository entity.VendaRepository
}

func NewGetVendaUseCase(vendaRepository entity.VendaRepository) *GetVendaUseCase {
	return &GetVendaUseCase{VendaRepository: vendaRepository}

}

func (u *GetVendaUseCase) GetVendaByID(id string) (*vendadto.VendaOutputDTO, error) {
	venda, err := u.VendaRepository.GetVendaByID(id)
	if err != nil {
		return nil, err
	}
	return &vendadto.VendaOutputDTO{
		ID:        venda.ID,
		EmpresaID: venda.EmpresaID,
		ClienteID: venda.ClienteID,
		VendaMc:   venda.VendaMc,
		Total:     venda.Total,
		Data:      venda.Data,
	}, nil
}

func (u *GetVendaUseCase) GetVendaByEmpresa(id string) ([]*vendadto.VendaOutputDTO, error) {
	vendas, err := u.VendaRepository.GetVendaByEmpresa(id)
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

func (u *GetVendaUseCase) GetVendaByStatus(status bool) ([]*vendadto.VendaOutputDTO, error) {
	vendas, err := u.VendaRepository.GetVendaByStatus(status)
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
