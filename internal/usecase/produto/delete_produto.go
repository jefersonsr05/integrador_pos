package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
)

type DeleteProdutoUseCase struct {
	ProdutoRepository entity.ProdutoRepository
}

func NewDeleteProdutoUseCase(produtoRepository entity.ProdutoRepository) *DeleteProdutoUseCase {
	return &DeleteProdutoUseCase{ProdutoRepository: produtoRepository}
}

func (u *DeleteProdutoUseCase) Execute(id string) error {

	err := u.ProdutoRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
