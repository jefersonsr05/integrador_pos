package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	empresadto "github.com/jefersonsr05/integrador_pos/internal/usecase/empresa/dto"
)

type CreateEmpresaUseCase struct {
	EmpresaRepository entity.EmpresaRepository
}

func NewCreateEmpresaUseCase(empresaRepository entity.EmpresaRepository) *CreateEmpresaUseCase {
	return &CreateEmpresaUseCase{EmpresaRepository: empresaRepository}
}

func (u *CreateEmpresaUseCase) Execute(input empresadto.EmpresaInputDTO) (*empresadto.EmpresaOutputDTO, error) {
	empresa := entity.NewEmpresa(input.Descricao, input.Cnpj, input.ChaveRegistro)
	err := u.EmpresaRepository.Create(empresa)
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
