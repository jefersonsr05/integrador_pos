package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
)

type DeleteCidadeUseCase struct {
	CidadeRepository entity.CidadeRepository
}

func NewDeleteCidadeUseCase(cidadeRepository entity.CidadeRepository) *DeleteCidadeUseCase {
	return &DeleteCidadeUseCase{CidadeRepository: cidadeRepository}
}

func (u *DeleteCidadeUseCase) Execute(id string) error {

	err := u.CidadeRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
