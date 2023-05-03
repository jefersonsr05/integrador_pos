package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	produtodto "github.com/jefersonsr05/integrador_pos/internal/usecase/produto/dto"
)

type GetProdutoUseCase struct {
	ProdutoRepository entity.ProdutoRepository
}

func NewGetProdutoUseCase(produtoRepository entity.ProdutoRepository) *GetProdutoUseCase {
	return &GetProdutoUseCase{ProdutoRepository: produtoRepository}

}

func (u *GetProdutoUseCase) GetProdutoByID(id string) (*produtodto.ProdutoOutputDTO, error) {
	produto, err := u.ProdutoRepository.GetProdutoByID(id)
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

func (u *GetProdutoUseCase) GetProdutoByEmpresa(id string) ([]*produtodto.ProdutoOutputDTO, error) {
	produtos, err := u.ProdutoRepository.GetProdutoByEmpresa(id)
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
