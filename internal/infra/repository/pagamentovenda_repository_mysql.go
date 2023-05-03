package repository

import (
	"context"
	"database/sql"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	database "github.com/jefersonsr05/integrador_pos/internal/infra/repository/db"
)

type PagamentoVendaRepositoryMysql struct {
	DB *sql.DB
}

func NewPagamentoVendaRepositoryMysql(db *sql.DB) *PagamentoVendaRepositoryMysql {
	return &PagamentoVendaRepositoryMysql{DB: db}
}

func (r *PagamentoVendaRepositoryMysql) Create(pagamentovenda *entity.PagamentoVenda) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.CreatePagamentoVenda(ctx, database.CreatePagamentoVendaParams{
		ID:          pagamentovenda.ID,
		VendaID:     pagamentovenda.VendaID,
		PagamentoID: pagamentovenda.PagamentoID,
		Valor:       pagamentovenda.Valor,
		Vencimento:  sql.NullTime{Time: pagamentovenda.Vencimento, Valid: true},
		Status:      pagamentovenda.Status,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *PagamentoVendaRepositoryMysql) FindAll() ([]*entity.PagamentoVenda, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaPagamentoVenda, err := queries.ListPagamentoVenda(ctx)
	if err != nil {
		return nil, err
	}

	var pagamentovendas []*entity.PagamentoVenda
	for _, pagamentovenda := range listaPagamentoVenda {
		var newPagamentoVenda entity.PagamentoVenda
		newPagamentoVenda.ID = pagamentovenda.ID
		newPagamentoVenda.VendaID = pagamentovenda.VendaID
		newPagamentoVenda.PagamentoID = pagamentovenda.PagamentoID
		newPagamentoVenda.Valor = pagamentovenda.Valor
		newPagamentoVenda.Vencimento = pagamentovenda.Vencimento.Time
		newPagamentoVenda.Status = pagamentovenda.Status

		pagamentovendas = append(pagamentovendas, &newPagamentoVenda)
	}
	return pagamentovendas, nil
}

func (r *PagamentoVendaRepositoryMysql) Update(pagamentovenda *entity.PagamentoVenda) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.UpdatePagamentoVenda(ctx, database.UpdatePagamentoVendaParams{
		ID:          pagamentovenda.ID,
		VendaID:     pagamentovenda.VendaID,
		PagamentoID: pagamentovenda.PagamentoID,
		Valor:       pagamentovenda.Valor,
		Vencimento:  sql.NullTime{Time: pagamentovenda.Vencimento, Valid: true},
		Status:      pagamentovenda.Status,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *PagamentoVendaRepositoryMysql) GetPagamentoVendaByID(id string) (*entity.PagamentoVenda, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	// log.Printf("passei aqui com o ID: " + id)

	pagamentovenda, err := queries.GetPagamentoVenda(ctx, id)
	if err != nil {
		return nil, err
	}

	var newPagamentoVenda entity.PagamentoVenda
	newPagamentoVenda.ID = pagamentovenda.ID
	newPagamentoVenda.VendaID = pagamentovenda.VendaID
	newPagamentoVenda.PagamentoID = pagamentovenda.PagamentoID
	newPagamentoVenda.Valor = pagamentovenda.Valor
	newPagamentoVenda.Vencimento = pagamentovenda.Vencimento.Time
	newPagamentoVenda.Status = pagamentovenda.Status

	return &newPagamentoVenda, nil
}

func (r *PagamentoVendaRepositoryMysql) GetPagamentoVendaByVendaID(venda string) ([]*entity.PagamentoVenda, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaPagamentoVenda, err := queries.GetPagamentoVendaByVendaID(ctx, venda)
	if err != nil {
		return nil, err
	}

	var pagamentovendas []*entity.PagamentoVenda
	for _, pagamentovenda := range listaPagamentoVenda {
		var newPagamentoVenda entity.PagamentoVenda
		newPagamentoVenda.ID = pagamentovenda.ID
		newPagamentoVenda.VendaID = pagamentovenda.VendaID
		newPagamentoVenda.PagamentoID = pagamentovenda.PagamentoID
		newPagamentoVenda.Valor = pagamentovenda.Valor
		newPagamentoVenda.Vencimento = pagamentovenda.Vencimento.Time
		newPagamentoVenda.Status = pagamentovenda.Status

		pagamentovendas = append(pagamentovendas, &newPagamentoVenda)
	}
	return pagamentovendas, nil
}

func (r *PagamentoVendaRepositoryMysql) GetPagamentoVendaByStatus(status bool) ([]*entity.PagamentoVenda, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaPagamentoVenda, err := queries.GetPagamentoVendaByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	var pagamentovendas []*entity.PagamentoVenda
	for _, pagamentovenda := range listaPagamentoVenda {
		var newPagamentoVenda entity.PagamentoVenda
		newPagamentoVenda.ID = pagamentovenda.ID
		newPagamentoVenda.VendaID = pagamentovenda.VendaID
		newPagamentoVenda.PagamentoID = pagamentovenda.PagamentoID
		newPagamentoVenda.Valor = pagamentovenda.Valor
		newPagamentoVenda.Vencimento = pagamentovenda.Vencimento.Time
		newPagamentoVenda.Status = pagamentovenda.Status

		pagamentovendas = append(pagamentovendas, &newPagamentoVenda)
	}
	return pagamentovendas, nil
}

func (r *PagamentoVendaRepositoryMysql) Delete(id string) error {
	ctx := context.Background()

	queries := database.New(r.DB)

	err := queries.DeletePagamentoVenda(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
