package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	produtodto "github.com/jefersonsr05/integrador_pos/internal/usecase/produto/dto"
)

type CreateProdutoUseCase struct {
	ProdutoRepository entity.ProdutoRepository
}

func NewCreateProdutoUseCase(produtoRepository entity.ProdutoRepository) *CreateProdutoUseCase {
	return &CreateProdutoUseCase{ProdutoRepository: produtoRepository}
}

func (u *CreateProdutoUseCase) Execute(input produtodto.ProdutoInputDTO) (*produtodto.ProdutoOutputDTO, error) {
	produto := entity.NewProduto(input.CodigoMc, input.EmpresaID, input.Descricao, input.CodBarras, input.Ncm, input.Cest, input.Cbenef, input.Preco)
	err := u.ProdutoRepository.Create(produto)
	if err != nil {
		return nil, err
	}
	return &produtodto.ProdutoOutputDTO{
		ID:        produto.ID,
		CodigoMc:  produto.CodigoMc,
		EmpresaID: produto.EmpresaID,
		Descricao: produto.Descricao,
		CodBarras: produto.CodBarras,
		Ncm:       produto.Ncm,
		Cest:      produto.Cest,
		Cbenef:    produto.Cbenef,
		Preco:     produto.Preco,
	}, nil
}
