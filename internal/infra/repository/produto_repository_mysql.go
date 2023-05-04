package repository

import (
	"context"
	"database/sql"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	database "github.com/jefersonsr05/integrador_pos/internal/infra/repository/db"
)

type ProdutoRepositoryMysql struct {
	DB *sql.DB
}

func NewProdutoRepositoryMysql(db *sql.DB) *ProdutoRepositoryMysql {
	return &ProdutoRepositoryMysql{DB: db}
}

func (r *ProdutoRepositoryMysql) Create(produto *entity.Produto) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.CreateProduto(ctx, database.CreateProdutoParams{
		ID:        produto.ID,
		EmpresaID: produto.EmpresaID,
		CodigoMc:  produto.CodigoMc,
		Descricao: produto.Descricao,
		CodBarras: sql.NullString{String: produto.CodBarras, Valid: true},
		Ncm:       sql.NullString{String: produto.Ncm, Valid: true},
		Cest:      sql.NullString{String: produto.Cest, Valid: true},
		Cbenef:    sql.NullString{String: produto.Cbenef, Valid: true},
		Preco:     produto.Preco,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *ProdutoRepositoryMysql) FindAll() ([]*entity.Produto, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaProduto, err := queries.ListProdutos(ctx)
	if err != nil {
		return nil, err
	}

	var produtos []*entity.Produto
	for _, produto := range listaProduto {
		var newProduto entity.Produto
		newProduto.ID = produto.ID
		newProduto.EmpresaID = produto.EmpresaID
		newProduto.CodigoMc = produto.CodigoMc
		newProduto.Descricao = produto.Descricao
		newProduto.CodBarras = produto.CodigoMc
		newProduto.Ncm = produto.Ncm.String
		newProduto.Cest = produto.Cest.String
		newProduto.Cbenef = produto.Cbenef.String
		newProduto.Preco = produto.Preco
		produtos = append(produtos, &newProduto)
	}
	return produtos, nil
}

func (r *ProdutoRepositoryMysql) Update(produto *entity.Produto) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.UpdateProduto(ctx, database.UpdateProdutoParams{
		ID:        produto.ID,
		EmpresaID: produto.EmpresaID,
		CodigoMc:  produto.CodigoMc,
		Descricao: produto.Descricao,
		CodBarras: sql.NullString{String: produto.CodBarras, Valid: true},
		Ncm:       sql.NullString{String: produto.Ncm, Valid: true},
		Cest:      sql.NullString{String: produto.Cest, Valid: true},
		Cbenef:    sql.NullString{String: produto.Cbenef, Valid: true},
		Preco:     produto.Preco,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *ProdutoRepositoryMysql) GetProdutoByID(id string) (*entity.Produto, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	// log.Printf("passei aqui com o ID: " + id)

	produto, err := queries.GetProduto(ctx, id)
	if err != nil {
		return nil, err
	}

	var newProduto entity.Produto
	newProduto.ID = produto.ID
	newProduto.EmpresaID = produto.EmpresaID
	newProduto.CodigoMc = produto.CodigoMc
	newProduto.Descricao = produto.Descricao
	newProduto.CodBarras = produto.CodigoMc
	newProduto.Ncm = produto.Ncm.String
	newProduto.Cest = produto.Cest.String
	newProduto.Cbenef = produto.Cbenef.String
	newProduto.Preco = produto.Preco

	return &newProduto, nil
}

func (r *ProdutoRepositoryMysql) GetProdutoByCodBarras(codbarras sql.NullString) (*entity.Produto, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	// log.Printf("passei aqui com o ID: " + id)

	produto, err := queries.GetProdutoByCodBarras(ctx, codbarras)
	if err != nil {
		return nil, err
	}

	var newProduto entity.Produto
	newProduto.EmpresaID = produto.EmpresaID
	newProduto.ID = produto.ID
	newProduto.CodigoMc = produto.CodigoMc
	newProduto.Descricao = produto.Descricao
	newProduto.CodBarras = produto.CodigoMc
	newProduto.Ncm = produto.Ncm.String
	newProduto.Cest = produto.Cest.String
	newProduto.Cbenef = produto.Cbenef.String
	newProduto.Preco = produto.Preco

	return &newProduto, nil
}

func (r *ProdutoRepositoryMysql) GetProdutoByCodigoMC(codigomc string) (*entity.Produto, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	// log.Printf("passei aqui com o ID: " + id)

	produto, err := queries.GetProdutoByCodigoMC(ctx, codigomc)
	if err != nil {
		return nil, err
	}

	var newProduto entity.Produto
	newProduto.ID = produto.ID
	newProduto.EmpresaID = produto.EmpresaID
	newProduto.CodigoMc = produto.CodigoMc
	newProduto.Descricao = produto.Descricao
	newProduto.CodBarras = produto.CodigoMc
	newProduto.Ncm = produto.Ncm.String
	newProduto.Cest = produto.Cest.String
	newProduto.Cbenef = produto.Cbenef.String
	newProduto.Preco = produto.Preco

	return &newProduto, nil
}

func (r *ProdutoRepositoryMysql) GetProdutoByEmpresa(empresa string) ([]*entity.Produto, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaProduto, err := queries.GetProdutoByEmpresa(ctx, empresa)
	if err != nil {
		return nil, err
	}

	var produtos []*entity.Produto
	for _, produto := range listaProduto {
		var newProduto entity.Produto
		newProduto.ID = produto.ID
		newProduto.EmpresaID = produto.EmpresaID
		newProduto.CodigoMc = produto.CodigoMc
		newProduto.Descricao = produto.Descricao
		newProduto.CodBarras = produto.CodigoMc
		newProduto.Ncm = produto.Ncm.String
		newProduto.Cest = produto.Cest.String
		newProduto.Cbenef = produto.Cbenef.String
		newProduto.Preco = produto.Preco
		produtos = append(produtos, &newProduto)
	}
	return produtos, nil
}

func (r *ProdutoRepositoryMysql) Delete(id string) error {
	ctx := context.Background()

	queries := database.New(r.DB)

	err := queries.DeleteProduto(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
