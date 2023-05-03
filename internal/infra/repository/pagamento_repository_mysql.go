package repository

import (
	"context"
	"database/sql"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	database "github.com/jefersonsr05/integrador_pos/internal/infra/repository/db"
)

type PagamentoRepositoryMysql struct {
	DB *sql.DB
}

func NewPagamentoRepositoryMysql(db *sql.DB) *PagamentoRepositoryMysql {
	return &PagamentoRepositoryMysql{DB: db}
}

func (r *PagamentoRepositoryMysql) Create(pagamento *entity.Pagamento) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.CreatePagamento(ctx, database.CreatePagamentoParams{
		ID:           pagamento.ID,
		Descricao:    pagamento.Descricao,
		TpPagamento:  pagamento.TpPagamento,
		IndPagamento: pagamento.IndPagamento,
		PosExclusivo: pagamento.PosExclusivo,
		IDPos:        sql.NullString{String: pagamento.IDPos, Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *PagamentoRepositoryMysql) FindAll() ([]*entity.Pagamento, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaPagamento, err := queries.ListPagamentos(ctx)
	if err != nil {
		return nil, err
	}

	var pagamentos []*entity.Pagamento
	for _, pagamento := range listaPagamento {
		var newPagamento entity.Pagamento
		newPagamento.ID = pagamento.ID
		newPagamento.Descricao = pagamento.Descricao
		newPagamento.TpPagamento = pagamento.TpPagamento
		newPagamento.IndPagamento = pagamento.IndPagamento
		newPagamento.PosExclusivo = pagamento.PosExclusivo
		newPagamento.IDPos = pagamento.IDPos.String

		pagamentos = append(pagamentos, &newPagamento)
	}
	return pagamentos, nil
}

func (r *PagamentoRepositoryMysql) Update(pagamento *entity.Pagamento) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.UpdatePagamento(ctx, database.UpdatePagamentoParams{
		ID:           pagamento.ID,
		Descricao:    pagamento.Descricao,
		TpPagamento:  pagamento.TpPagamento,
		IndPagamento: pagamento.IndPagamento,
		PosExclusivo: pagamento.PosExclusivo,
		IDPos:        sql.NullString{String: pagamento.IDPos, Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *PagamentoRepositoryMysql) GetPagamentoByID(id string) (*entity.Pagamento, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	// log.Printf("passei aqui com o ID: " + id)

	pagamento, err := queries.GetPagamento(ctx, id)
	if err != nil {
		return nil, err
	}

	var newPagamento entity.Pagamento
	newPagamento.ID = pagamento.ID
	newPagamento.Descricao = pagamento.Descricao
	newPagamento.TpPagamento = pagamento.TpPagamento
	newPagamento.IndPagamento = pagamento.IndPagamento
	newPagamento.PosExclusivo = pagamento.PosExclusivo
	newPagamento.IDPos = pagamento.IDPos.String

	return &newPagamento, nil
}

func (r *PagamentoRepositoryMysql) GetPagamentoByEmpresaID(empresaid string) ([]*entity.Pagamento, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaPagamento, err := queries.GetPagamentoByEmpresaID(ctx, empresaid)
	if err != nil {
		return nil, err
	}

	var pagamentos []*entity.Pagamento
	for _, pagamento := range listaPagamento {
		var newPagamento entity.Pagamento
		newPagamento.ID = pagamento.ID
		newPagamento.Descricao = pagamento.Descricao
		newPagamento.TpPagamento = pagamento.TpPagamento
		newPagamento.IndPagamento = pagamento.IndPagamento
		newPagamento.PosExclusivo = pagamento.PosExclusivo
		newPagamento.IDPos = pagamento.IDPos.String

		pagamentos = append(pagamentos, &newPagamento)
	}
	return pagamentos, nil
}

func (r *PagamentoRepositoryMysql) Delete(id string) error {
	ctx := context.Background()

	queries := database.New(r.DB)

	err := queries.DeletePagamento(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
