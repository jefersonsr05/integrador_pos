package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	produtodto "github.com/jefersonsr05/integrador_pos/internal/usecase/produto/dto"
)

type ListProdutoUseCase struct {
	ProdutoRepository entity.ProdutoRepository
}

func NewListProdutoUseCase(produtoRepository entity.ProdutoRepository) *ListProdutoUseCase {
	return &ListProdutoUseCase{ProdutoRepository: produtoRepository}

}

func (u *ListProdutoUseCase) Execute() ([]*produtodto.ProdutoOutputDTO, error) {
	produtos, err := u.ProdutoRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var produtosOutput []*produtodto.ProdutoOutputDTO
	for _, produto := range produtos {
		produtosOutput = append(produtosOutput, &produtodto.ProdutoOutputDTO{
			ID:        produto.ID,
			CodigoMc:  produto.CodigoMc,
			EmpresaID: produto.EmpresaID,
			Descricao: produto.Descricao,
			CodBarras: produto.CodBarras,
			Ncm:       produto.Ncm,
			Cest:      produto.Cest,
			Cbenef:    produto.Cbenef,
			Preco:     produto.Preco,
		})
	}
	return produtosOutput, nil
}
