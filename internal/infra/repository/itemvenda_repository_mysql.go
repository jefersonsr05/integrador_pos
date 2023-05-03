package repository

import (
	"context"
	"database/sql"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	database "github.com/jefersonsr05/integrador_pos/internal/infra/repository/db"
)

type ItemVendaRepositoryMysql struct {
	DB *sql.DB
}

func NewItemVendaRepositoryMysql(db *sql.DB) *ItemVendaRepositoryMysql {
	return &ItemVendaRepositoryMysql{DB: db}
}

func (r *ItemVendaRepositoryMysql) Create(itemvenda *entity.ItemVenda) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.CreateItemVenda(ctx, database.CreateItemVendaParams{
		ID:         itemvenda.ID,
		EmpresaID:  itemvenda.EmpresaID,
		VendaID:    itemvenda.VendaID,
		ProdutoID:  itemvenda.ProdutoID,
		Quantidade: itemvenda.Quantidade,
		Valor:      itemvenda.Valor,
		Total:      itemvenda.Total,

		// ChaveRegistro: sql.NullString{String: "114719062017110014", Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *ItemVendaRepositoryMysql) Update(itemvenda *entity.ItemVenda) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.UpdateItemVenda(ctx, database.UpdateItemVendaParams{
		ID:         itemvenda.ID,
		EmpresaID:  itemvenda.EmpresaID,
		VendaID:    itemvenda.VendaID,
		ProdutoID:  itemvenda.ProdutoID,
		Quantidade: itemvenda.Quantidade,
		Valor:      itemvenda.Valor,
		Total:      itemvenda.Total,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *ItemVendaRepositoryMysql) GetItemVendaByID(id string) (*entity.ItemVenda, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	// log.Printf("passei aqui com o ID: " + id)

	itemvenda, err := queries.GetItemVenda(ctx, id)
	if err != nil {
		return nil, err
	}

	var newItemVenda entity.ItemVenda
	newItemVenda.ID = itemvenda.ID
	newItemVenda.EmpresaID = itemvenda.EmpresaID
	newItemVenda.VendaID = itemvenda.VendaID
	newItemVenda.ProdutoID = itemvenda.ProdutoID
	newItemVenda.Quantidade = itemvenda.Quantidade
	newItemVenda.Valor = itemvenda.Valor
	newItemVenda.Total = itemvenda.Total

	return &newItemVenda, nil
}

func (r *ItemVendaRepositoryMysql) FindAll() ([]*entity.ItemVenda, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaItemVenda, err := queries.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var itensvenda []*entity.ItemVenda
	for _, itemvenda := range listaItemVenda {
		var newItemVenda entity.ItemVenda
		newItemVenda.ID = itemvenda.ID
		newItemVenda.EmpresaID = itemvenda.EmpresaID
		newItemVenda.VendaID = itemvenda.VendaID
		newItemVenda.ProdutoID = itemvenda.ProdutoID
		newItemVenda.Quantidade = itemvenda.Quantidade
		newItemVenda.Valor = itemvenda.Valor
		newItemVenda.Total = itemvenda.Total
		itensvenda = append(itensvenda, &newItemVenda)
	}
	return itensvenda, nil
}

func (r *ItemVendaRepositoryMysql) GetItemVendaByVendaID(venda string) ([]*entity.ItemVenda, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaItemVenda, err := queries.GetItensVenda(ctx, venda)
	if err != nil {
		return nil, err
	}

	var itensvenda []*entity.ItemVenda
	for _, itemvenda := range listaItemVenda {
		var newItemVenda entity.ItemVenda
		newItemVenda.ID = itemvenda.ID
		newItemVenda.EmpresaID = itemvenda.EmpresaID
		newItemVenda.VendaID = itemvenda.VendaID
		newItemVenda.ProdutoID = itemvenda.ProdutoID
		newItemVenda.Quantidade = itemvenda.Quantidade
		newItemVenda.Valor = itemvenda.Valor
		newItemVenda.Total = itemvenda.Total
		itensvenda = append(itensvenda, &newItemVenda)
	}
	return itensvenda, nil
}

func (r *ItemVendaRepositoryMysql) Delete(id string) error {
	ctx := context.Background()

	queries := database.New(r.DB)

	err := queries.DeleteItemVenda(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
