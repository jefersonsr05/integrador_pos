package usecase

import (
	"log"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	empresadto "github.com/jefersonsr05/integrador_pos/internal/usecase/empresa/dto"
)

type UpdateEmpresaUseCase struct {
	EmpresaRepository entity.EmpresaRepository
}

func NewUpdateEmpresaUseCase(empresaRepository entity.EmpresaRepository) *UpdateEmpresaUseCase {
	return &UpdateEmpresaUseCase{EmpresaRepository: empresaRepository}
}

func (u *UpdateEmpresaUseCase) Execute(id string, input empresadto.EmpresaInputDTO) (*empresadto.EmpresaOutputDTO, error) {
	empresa, err := u.EmpresaRepository.GetEmpresaByID(id)
	if err != nil {
		log.Printf("não achei resultado")
		return nil, err
	}

	empresa.Descricao = input.Descricao
	empresa.Cnpj = input.Cnpj
	empresa.ChaveRegistro = input.ChaveRegistro

	err = u.EmpresaRepository.Update(empresa)
	if err != nil {
		return nil, err
	}

	return &empresadto.EmpresaOutputDTO{
		ID:            empresa.ID,
		Descricao:     empresa.Descricao,
		Cnpj:          empresa.Cnpj,
		ChaveRegistro: empresa.ChaveRegistro,
	}, nil
}
