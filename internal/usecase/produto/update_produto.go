package usecase

import (
	"log"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	produtodto "github.com/jefersonsr05/integrador_pos/internal/usecase/produto/dto"
)

type UpdateProdutoUseCase struct {
	ProdutoRepository entity.ProdutoRepository
}

func NewUpdateProdutoUseCase(produtoRepository entity.ProdutoRepository) *UpdateProdutoUseCase {
	return &UpdateProdutoUseCase{ProdutoRepository: produtoRepository}
}

func (u *UpdateProdutoUseCase) Execute(id string, input produtodto.ProdutoInputDTO) (*produtodto.ProdutoOutputDTO, error) {
	produto, err := u.ProdutoRepository.GetProdutoByID(id)
	if err != nil {
		log.Printf("não achei resultado")
		return nil, err
	}

	produto.CodigoMc = input.CodigoMc
	produto.EmpresaID = input.EmpresaID
	produto.Descricao = input.Descricao
	produto.CodBarras = input.CodBarras
	produto.Ncm = input.Ncm
	produto.Cest = input.Cest
	produto.Cbenef = input.Cbenef
	produto.Preco = input.Preco

	err = u.ProdutoRepository.Update(produto)
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
