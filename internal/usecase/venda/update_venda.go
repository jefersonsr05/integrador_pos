package usecase

import (
	"log"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	vendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/venda/dto"
)

type UpdateVendaUseCase struct {
	VendaRepository entity.VendaRepository
}

func NewUpdateVendaUseCase(vendaRepository entity.VendaRepository) *UpdateVendaUseCase {
	return &UpdateVendaUseCase{VendaRepository: vendaRepository}
}

func (u *UpdateVendaUseCase) Execute(id string, input vendadto.VendaInputDTO) (*vendadto.VendaOutputDTO, error) {
	venda, err := u.VendaRepository.GetVendaByID(id)
	if err != nil {
		log.Printf("não achei resultado")
		return nil, err
	}

	venda.EmpresaID = input.EmpresaID
	venda.ClienteID = input.ClienteID
	venda.VendaMc = input.VendaMc
	venda.Total = input.Total
	venda.Data = input.Data

	err = u.VendaRepository.Update(venda)
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
