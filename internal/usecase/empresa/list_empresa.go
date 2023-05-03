package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	empresadto "github.com/jefersonsr05/integrador_pos/internal/usecase/empresa/dto"
)

type ListEmpresasUseCase struct {
	EmpresaRepository entity.EmpresaRepository
}

func NewListEmpresasUseCase(empresaRepository entity.EmpresaRepository) *ListEmpresasUseCase {
	return &ListEmpresasUseCase{EmpresaRepository: empresaRepository}

}

func (u *ListEmpresasUseCase) Execute() ([]*empresadto.EmpresaOutputDTO, error) {
	empresas, err := u.EmpresaRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var empresasOutput []*empresadto.EmpresaOutputDTO
	for _, empresa := range empresas {
		empresasOutput = append(empresasOutput, &empresadto.EmpresaOutputDTO{
			ID:            empresa.ID,
			Descricao:     empresa.Descricao,
			Cnpj:          empresa.Cnpj,
			ChaveRegistro: empresa.ChaveRegistro,
		})
	}
	return empresasOutput, nil
}
