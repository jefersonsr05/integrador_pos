package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
)

type DeleteEmpresaUseCase struct {
	EmpresaRepository entity.EmpresaRepository
}

func NewDeleteEmpresaUseCase(empresaRepository entity.EmpresaRepository) *DeleteEmpresaUseCase {
	return &DeleteEmpresaUseCase{EmpresaRepository: empresaRepository}
}

func (u *DeleteEmpresaUseCase) Execute(id string) error {

	err := u.EmpresaRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
