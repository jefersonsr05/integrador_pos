package repository

import (
	"context"
	"database/sql"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	database "github.com/jefersonsr05/integrador_pos/internal/infra/repository/db"
)

type VendaRepositoryMysql struct {
	DB *sql.DB
}

func NewVendaRepositoryMysql(db *sql.DB) *VendaRepositoryMysql {
	return &VendaRepositoryMysql{DB: db}
}

func (r *VendaRepositoryMysql) Create(venda *entity.Venda) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.CreateVenda(ctx, database.CreateVendaParams{
		ID:        venda.ID,
		EmpresaID: venda.EmpresaID,
		ClienteID: venda.ClienteID,
		VendaMc:   venda.VendaMc,
		Total:     venda.Total,
		Data:      sql.NullTime{Time: venda.Data, Valid: true},
		// ChaveRegistro: sql.NullString{String: "114719062017110014", Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *VendaRepositoryMysql) FindAll() ([]*entity.Venda, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaVenda, err := queries.ListVendas(ctx)
	if err != nil {
		return nil, err
	}

	var vendas []*entity.Venda
	for _, venda := range listaVenda {
		var newVenda entity.Venda
		newVenda.ID = venda.ID
		newVenda.EmpresaID = venda.EmpresaID
		newVenda.ClienteID = venda.ClienteID
		newVenda.VendaMc = venda.VendaMc
		newVenda.Total = venda.Total
		newVenda.Data = venda.Data.Time
		vendas = append(vendas, &newVenda)
	}
	return vendas, nil
}

func (r *VendaRepositoryMysql) Update(venda *entity.Venda) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.UpdateVenda(ctx, database.UpdateVendaParams{
		ID:        venda.ID,
		EmpresaID: venda.EmpresaID,
		ClienteID: venda.ClienteID,
		VendaMc:   venda.VendaMc,
		Total:     venda.Total,
		Data:      sql.NullTime{Time: venda.Data, Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *VendaRepositoryMysql) GetVendaByID(id string) (*entity.Venda, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	// log.Printf("passei aqui com o ID: " + id)

	venda, err := queries.GetVenda(ctx, id)
	if err != nil {
		return nil, err
	}

	var newVenda entity.Venda
	newVenda.ID = venda.ID
	newVenda.EmpresaID = venda.EmpresaID
	newVenda.ClienteID = venda.ClienteID
	newVenda.VendaMc = venda.VendaMc
	newVenda.Total = venda.Total
	newVenda.Data = venda.Data.Time

	return &newVenda, nil
}

func (r *VendaRepositoryMysql) GetVendaByEmpresa(empresa string) ([]*entity.Venda, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaVenda, err := queries.GetVendaByEmpresa(ctx, empresa)
	if err != nil {
		return nil, err
	}

	var vendas []*entity.Venda
	for _, venda := range listaVenda {
		var newVenda entity.Venda
		newVenda.ID = venda.ID
		newVenda.EmpresaID = venda.EmpresaID
		newVenda.ClienteID = venda.ClienteID
		newVenda.VendaMc = venda.VendaMc
		newVenda.Total = venda.Total
		newVenda.Data = venda.Data.Time
		vendas = append(vendas, &newVenda)
	}
	return vendas, nil
}

func (r *VendaRepositoryMysql) GetVendaByStatus(status bool) ([]*entity.Venda, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaVenda, err := queries.GetVendaByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	var vendas []*entity.Venda
	for _, venda := range listaVenda {
		var newVenda entity.Venda
		newVenda.ID = venda.ID
		newVenda.EmpresaID = venda.EmpresaID
		newVenda.ClienteID = venda.ClienteID
		newVenda.VendaMc = venda.VendaMc
		newVenda.Total = venda.Total
		newVenda.Data = venda.Data.Time
		vendas = append(vendas, &newVenda)
	}
	return vendas, nil
}

func (r *VendaRepositoryMysql) Delete(id string) error {
	ctx := context.Background()

	queries := database.New(r.DB)

	err := queries.DeleteVenda(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
