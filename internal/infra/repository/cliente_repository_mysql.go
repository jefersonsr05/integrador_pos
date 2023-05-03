package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	database "github.com/jefersonsr05/integrador_pos/internal/infra/repository/db"
)

type ClienteRepositoryMysql struct {
	DB *sql.DB
}

func NewClienteRepositoryMysql(db *sql.DB) *ClienteRepositoryMysql {
	return &ClienteRepositoryMysql{DB: db}
}

func (r *ClienteRepositoryMysql) Create(cliente *entity.Cliente) error {
	ctx := context.Background()
	queries := database.New(r.DB)
	log.Printf("Create Repository")
	err := queries.CreateCliente(ctx, database.CreateClienteParams{
		ID:          cliente.ID,
		CodigoMc:    int64(cliente.CodigoMc),
		EmpresaID:   cliente.EmpresaID,
		Nome:        cliente.Nome,
		Cep:         sql.NullString{String: cliente.Cep, Valid: true},
		CidadeID:    sql.NullString{String: cliente.CidadeID, Valid: true},
		Endereco:    sql.NullString{String: cliente.Endereco, Valid: true},
		Numero:      sql.NullString{String: cliente.Numero, Valid: true},
		Complemento: sql.NullString{String: cliente.Complemento, Valid: true},
	})

	if err != nil {
		log.Printf("Erro Create Repository")
		return err
	}

	return nil
}

func (r *ClienteRepositoryMysql) FindAll() ([]*entity.Cliente, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaCliente, err := queries.ListClientes(ctx)
	if err != nil {
		return nil, err
	}

	var clientes []*entity.Cliente
	for _, cliente := range listaCliente {
		var newCliente entity.Cliente
		newCliente.ID = cliente.ID
		newCliente.CodigoMc = int64(cliente.CodigoMc)
		newCliente.EmpresaID = cliente.EmpresaID
		newCliente.Nome = cliente.Nome
		newCliente.Cep = cliente.Cep.String
		newCliente.CidadeID = cliente.CidadeID.String
		newCliente.Endereco = cliente.Endereco.String
		newCliente.Numero = cliente.Numero.String
		newCliente.Complemento = cliente.Complemento.String
		clientes = append(clientes, &newCliente)
	}
	return clientes, nil
}

func (r *ClienteRepositoryMysql) Update(cliente *entity.Cliente) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.UpdateCliente(ctx, database.UpdateClienteParams{
		ID:          cliente.ID,
		CodigoMc:    int64(cliente.CodigoMc),
		EmpresaID:   cliente.EmpresaID,
		Nome:        cliente.Nome,
		Cep:         sql.NullString{String: cliente.Cep, Valid: true},
		CidadeID:    sql.NullString{String: cliente.CidadeID, Valid: true},
		Endereco:    sql.NullString{String: cliente.Endereco, Valid: true},
		Numero:      sql.NullString{String: cliente.Numero, Valid: true},
		Complemento: sql.NullString{String: cliente.Complemento, Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *ClienteRepositoryMysql) GetClienteByEmpresa(empresa string) ([]*entity.Cliente, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaCliente, err := queries.GetClienteByEmpresa(ctx, empresa)
	if err != nil {
		return nil, err
	}

	var clientes []*entity.Cliente
	for _, cliente := range listaCliente {
		var newCliente entity.Cliente
		newCliente.ID = cliente.ID
		newCliente.CodigoMc = int64(cliente.CodigoMc)
		newCliente.EmpresaID = cliente.EmpresaID
		newCliente.Nome = cliente.Nome
		newCliente.Cep = cliente.Cep.String
		newCliente.CidadeID = cliente.CidadeID.String
		newCliente.Endereco = cliente.Endereco.String
		newCliente.Numero = cliente.Numero.String
		newCliente.Complemento = cliente.Complemento.String

		clientes = append(clientes, &newCliente)
	}
	return clientes, nil
}

func (r *ClienteRepositoryMysql) GetClienteByID(id string) (*entity.Cliente, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	cliente, err := queries.GetCliente(ctx, id)
	if err != nil {
		return nil, err
	}

	var newCliente entity.Cliente
	newCliente.ID = cliente.ID
	newCliente.CodigoMc = int64(cliente.CodigoMc)
	newCliente.EmpresaID = cliente.EmpresaID
	newCliente.Nome = cliente.Nome
	newCliente.Cep = cliente.Cep.String
	newCliente.CidadeID = cliente.CidadeID.String
	newCliente.Endereco = cliente.Endereco.String
	newCliente.Numero = cliente.Numero.String
	newCliente.Complemento = cliente.Complemento.String

	return &newCliente, nil
}

func (r *ClienteRepositoryMysql) GetClienteByCodigoMC(codigomc int64) (*entity.Cliente, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	cliente, err := queries.GetClienteByCodigoMC(ctx, codigomc)
	if err != nil {
		return nil, err
	}

	var newCliente entity.Cliente
	newCliente.ID = cliente.ID
	newCliente.CodigoMc = int64(cliente.CodigoMc)
	newCliente.EmpresaID = cliente.EmpresaID
	newCliente.Nome = cliente.Nome
	newCliente.Cep = cliente.Cep.String
	newCliente.CidadeID = cliente.CidadeID.String
	newCliente.Endereco = cliente.Endereco.String
	newCliente.Numero = cliente.Numero.String
	newCliente.Complemento = cliente.Complemento.String

	return &newCliente, nil
}

func (r *ClienteRepositoryMysql) Delete(id string) error {
	ctx := context.Background()

	queries := database.New(r.DB)

	err := queries.DeleteCliente(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
