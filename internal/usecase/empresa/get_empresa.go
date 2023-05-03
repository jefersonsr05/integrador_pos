package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	empresadto "github.com/jefersonsr05/integrador_pos/internal/usecase/empresa/dto"
)

type GetEmpresaUseCase struct {
	EmpresaRepository entity.EmpresaRepository
}

func NewGetEmpresaUseCase(empresaRepository entity.EmpresaRepository) *GetEmpresaUseCase {
	return &GetEmpresaUseCase{EmpresaRepository: empresaRepository}

}

func (u *GetEmpresaUseCase) GetEmpresaByID(id string) (*empresadto.EmpresaOutputDTO, error) {
	empresa, err := u.EmpresaRepository.GetEmpresaByID(id)
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

func (u *GetEmpresaUseCase) GetEmpresaByCNPJ(cnpj string) (*empresadto.EmpresaOutputDTO, error) {
	empresa, err := u.EmpresaRepository.GetEmpresaByCNPJ(cnpj)
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

func (u *GetEmpresaUseCase) GetEmpresaByChaveRegistro(chave string) (*empresadto.EmpresaOutputDTO, error) {
	empresa, err := u.EmpresaRepository.GetEmpresaByChaveRegistro(chave)
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
