package repository

import (
	"context"
	"database/sql"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	database "github.com/jefersonsr05/integrador_pos/internal/infra/repository/db"
)

type EmpresaRepositoryMysql struct {
	DB *sql.DB
}

func NewEmpresaRepositoryMysql(db *sql.DB) *EmpresaRepositoryMysql {
	return &EmpresaRepositoryMysql{DB: db}
}

func (r *EmpresaRepositoryMysql) Create(empresa *entity.Empresa) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	// log.Printf(empresa.ID)
	// log.Printf(empresa.Descricao)
	// log.Printf(empresa.Cnpj)
	// log.Printf(empresa.ChaveRegistro)

	err := queries.CreateEmpresa(ctx, database.CreateEmpresaParams{
		ID:            empresa.ID,
		Descricao:     empresa.Descricao,
		Cnpj:          sql.NullString{String: empresa.Cnpj, Valid: true},
		ChaveRegistro: sql.NullString{String: empresa.ChaveRegistro, Valid: true},
		// ChaveRegistro: sql.NullString{String: "114719062017110014", Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *EmpresaRepositoryMysql) FindAll() ([]*entity.Empresa, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaEmpresa, err := queries.ListEmpresas(ctx)
	if err != nil {
		return nil, err
	}

	var empresas []*entity.Empresa
	for _, empresa := range listaEmpresa {
		var newEmpresa entity.Empresa
		newEmpresa.ID = empresa.ID
		newEmpresa.Descricao = empresa.Descricao
		newEmpresa.Cnpj = empresa.Cnpj.String
		newEmpresa.ChaveRegistro = empresa.ChaveRegistro.String
		empresas = append(empresas, &newEmpresa)
	}
	return empresas, nil
}

func (r *EmpresaRepositoryMysql) Update(empresa *entity.Empresa) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.UpdateEmpresa(ctx, database.UpdateEmpresaParams{
		ID:            empresa.ID,
		Descricao:     empresa.Descricao,
		Cnpj:          sql.NullString{String: empresa.Cnpj, Valid: true},
		ChaveRegistro: sql.NullString{String: empresa.ChaveRegistro, Valid: true},
		// ChaveRegistro: sql.NullString{String: "114719062017110014", Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *EmpresaRepositoryMysql) GetEmpresaByID(id string) (*entity.Empresa, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	// log.Printf("passei aqui com o ID: " + id)

	empresa, err := queries.GetEmpresa(ctx, id)
	if err != nil {
		return nil, err
	}

	var newEmpresa entity.Empresa
	newEmpresa.ID = empresa.ID
	newEmpresa.Descricao = empresa.Descricao
	newEmpresa.Cnpj = empresa.Cnpj.String
	newEmpresa.ChaveRegistro = empresa.ChaveRegistro.String

	return &newEmpresa, nil
}

func (r *EmpresaRepositoryMysql) GetEmpresaByCNPJ(cnpj string) (*entity.Empresa, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	empresa, err := queries.GetEmpresaByCNPJ(ctx, sql.NullString{String: cnpj, Valid: true})
	if err != nil {
		return nil, err
	}

	var newEmpresa entity.Empresa
	newEmpresa.ID = empresa.ID
	newEmpresa.Descricao = empresa.Descricao
	newEmpresa.Cnpj = empresa.Cnpj.String
	newEmpresa.ChaveRegistro = empresa.ChaveRegistro.String

	return &newEmpresa, nil
}

func (r *EmpresaRepositoryMysql) GetEmpresaByChaveRegistro(chave string) (*entity.Empresa, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	empresa, err := queries.GetEmpresaByChaveRegistro(ctx, sql.NullString{String: chave, Valid: true})
	if err != nil {
		return nil, err
	}

	var newEmpresa entity.Empresa
	newEmpresa.ID = empresa.ID
	newEmpresa.Descricao = empresa.Descricao
	newEmpresa.Cnpj = empresa.Cnpj.String
	newEmpresa.ChaveRegistro = empresa.ChaveRegistro.String

	return &newEmpresa, nil
}

func (r *EmpresaRepositoryMysql) Delete(id string) error {
	ctx := context.Background()

	queries := database.New(r.DB)

	err := queries.DeleteEmpresa(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
