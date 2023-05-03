package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
)

type DeleteMaquininhaPosUseCase struct {
	MaquininhaPosRepository entity.MaquininhaPosRepository
}

func NewDeleteMaquininhaPosUseCase(maquininhaPosRepository entity.MaquininhaPosRepository) *DeleteMaquininhaPosUseCase {
	return &DeleteMaquininhaPosUseCase{MaquininhaPosRepository: maquininhaPosRepository}
}

func (u *DeleteMaquininhaPosUseCase) Execute(id string) error {

	err := u.MaquininhaPosRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
