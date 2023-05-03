package repository

import (
	"context"
	"database/sql"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	database "github.com/jefersonsr05/integrador_pos/internal/infra/repository/db"
)

type CidadeRepositoryMysql struct {
	DB *sql.DB
}

func NewCidadeRepositoryMysql(db *sql.DB) *CidadeRepositoryMysql {
	return &CidadeRepositoryMysql{DB: db}
}

func (r *CidadeRepositoryMysql) Create(cidade *entity.Cidade) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.CreateCidade(ctx, database.CreateCidadeParams{
		ID:        cidade.ID,
		Descricao: cidade.Descricao,
		EstadoID:  cidade.EstadoID,
		CodIbge:   cidade.CodIbge,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *CidadeRepositoryMysql) FindAll() ([]*entity.Cidade, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaCidade, err := queries.ListCidades(ctx)
	if err != nil {
		return nil, err
	}

	var cidades []*entity.Cidade
	for _, cidade := range listaCidade {
		var newCidade entity.Cidade
		newCidade.ID = cidade.ID
		newCidade.Descricao = cidade.Descricao
		newCidade.EstadoID = cidade.EstadoID
		newCidade.CodIbge = cidade.CodIbge
		cidades = append(cidades, &newCidade)
	}
	return cidades, nil
}

func (r *CidadeRepositoryMysql) Update(cidade *entity.Cidade) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.UpdateCidade(ctx, database.UpdateCidadeParams{
		ID:        cidade.ID,
		Descricao: cidade.Descricao,
		EstadoID:  cidade.EstadoID,
		CodIbge:   cidade.CodIbge,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *CidadeRepositoryMysql) GetCidadeByID(id string) (*entity.Cidade, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	// log.Printf("passei aqui com o ID: " + id)

	cidade, err := queries.GetCidade(ctx, id)
	if err != nil {
		return nil, err
	}

	var newCidade entity.Cidade
	newCidade.ID = cidade.ID
	newCidade.Descricao = cidade.Descricao
	newCidade.EstadoID = cidade.EstadoID
	newCidade.CodIbge = cidade.CodIbge

	return &newCidade, nil
}

func (r *CidadeRepositoryMysql) GetCidadeByIBGE(ibge string) (*entity.Cidade, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	cidade, err := queries.GetCidadeByIBGE(ctx, ibge)
	if err != nil {
		return nil, err
	}

	var newCidade entity.Cidade
	newCidade.ID = cidade.ID
	newCidade.Descricao = cidade.Descricao
	newCidade.EstadoID = cidade.EstadoID
	newCidade.CodIbge = cidade.CodIbge

	return &newCidade, nil
}

func (r *CidadeRepositoryMysql) GetCidadeByUF(uf string) ([]*entity.Cidade, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaCidade, err := queries.GetCidadeByUF(ctx, uf)
	if err != nil {
		return nil, err
	}

	var cidades []*entity.Cidade
	for _, cidade := range listaCidade {
		var newCidade entity.Cidade
		newCidade.ID = cidade.ID
		newCidade.Descricao = cidade.Descricao
		newCidade.EstadoID = cidade.EstadoID
		newCidade.CodIbge = cidade.CodIbge
		cidades = append(cidades, &newCidade)
	}
	return cidades, nil
}

func (r *CidadeRepositoryMysql) GetCidadeByIbgeUF(ibgeuf int32) ([]*entity.Cidade, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listaCidade, err := queries.GetCidadeByIbgeUF(ctx, ibgeuf)
	if err != nil {
		return nil, err
	}

	var cidades []*entity.Cidade
	for _, cidade := range listaCidade {
		var newCidade entity.Cidade
		newCidade.ID = cidade.ID
		newCidade.Descricao = cidade.Descricao
		newCidade.EstadoID = cidade.EstadoID
		newCidade.CodIbge = cidade.CodIbge
		cidades = append(cidades, &newCidade)
	}
	return cidades, nil
}

func (r *CidadeRepositoryMysql) Delete(id string) error {
	ctx := context.Background()

	queries := database.New(r.DB)

	err := queries.DeleteCidade(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
