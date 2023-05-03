package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	vendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/venda/dto"
)

type CreateVendaUseCase struct {
	VendaRepository entity.VendaRepository
}

func NewCreateVendaUseCase(vendaRepository entity.VendaRepository) *CreateVendaUseCase {
	return &CreateVendaUseCase{VendaRepository: vendaRepository}
}

func (u *CreateVendaUseCase) Execute(input vendadto.VendaInputDTO) (*vendadto.VendaOutputDTO, error) {
	venda := entity.NewVenda(input.EmpresaID, input.ClienteID, input.VendaMc, input.Total, input.Data)
	err := u.VendaRepository.Create(venda)
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
