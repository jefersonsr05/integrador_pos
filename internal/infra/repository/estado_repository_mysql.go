package repository

import (
	"context"
	"database/sql"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	database "github.com/jefersonsr05/integrador_pos/internal/infra/repository/db"
)

type EstadoRepositoryMysql struct {
	DB *sql.DB
}

func NewEstadoRepositoryMysql(db *sql.DB) *EstadoRepositoryMysql {
	return &EstadoRepositoryMysql{DB: db}
}

func (r *EstadoRepositoryMysql) Create(estado *entity.Estado) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.CreateEstado(ctx, database.CreateEstadoParams{
		ID:        estado.ID,
		Descricao: estado.Descricao,
		Uf:        estado.UF,
		CodIbge:   estado.CodIbge,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *EstadoRepositoryMysql) FindAll() ([]*entity.Estado, error) {
	ctx := context.Background()
	// log.Printf(fmt.Sprint(r.DB.Stats().OpenConnections, r.DB.Stats().InUse, r.DB.Stats().Idle, r.DB.Stats().MaxOpenConnections))
	queries := database.New(r.DB)

	listaEstado, err := queries.ListEstados(ctx)
	if err != nil {
		return nil, err
	}

	var estados []*entity.Estado
	for _, estado := range listaEstado {
		var newEstado entity.Estado
		newEstado.ID = estado.ID
		newEstado.Descricao = estado.Descricao
		newEstado.UF = estado.Uf
		newEstado.CodIbge = estado.CodIbge
		estados = append(estados, &newEstado)
	}
	return estados, nil
}

func (r *EstadoRepositoryMysql) Update(estado *entity.Estado) error {
	ctx := context.Background()
	// log.Printf(fmt.Sprint(r.DB.Stats().OpenConnections, r.DB.Stats().InUse, r.DB.Stats().Idle, r.DB.Stats().MaxOpenConnections))
	queries := database.New(r.DB)

	err := queries.UpdateEstado(ctx, database.UpdateEstadoParams{
		ID:        estado.ID,
		Descricao: estado.Descricao,
		Uf:        estado.UF,
		CodIbge:   estado.CodIbge,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *EstadoRepositoryMysql) GetEstadoByID(id string) (*entity.Estado, error) {
	ctx := context.Background()
	// log.Printf(fmt.Sprint(r.DB.Stats().OpenConnections, r.DB.Stats().InUse, r.DB.Stats().Idle, r.DB.Stats().MaxOpenConnections))
	queries := database.New(r.DB)

	// log.Printf("passei aqui com o ID: " + id)

	estado, err := queries.GetEstado(ctx, id)
	if err != nil {
		return nil, err
	}

	var newEstado entity.Estado
	newEstado.ID = estado.ID
	newEstado.Descricao = estado.Descricao
	newEstado.UF = estado.Uf
	newEstado.CodIbge = estado.CodIbge

	return &newEstado, nil
}

func (r *EstadoRepositoryMysql) GetEstadoByCodIBGE(ibge int32) (*entity.Estado, error) {
	ctx := context.Background()
	// log.Printf(fmt.Sprint(r.DB.Stats().OpenConnections, r.DB.Stats().InUse, r.DB.Stats().Idle, r.DB.Stats().MaxOpenConnections))
	queries := database.New(r.DB)

	estado, err := queries.GetEstadoByCodIBGE(ctx, ibge)
	if err != nil {
		return nil, err
	}

	var newEstado entity.Estado
	newEstado.ID = estado.ID
	newEstado.Descricao = estado.Descricao
	newEstado.UF = estado.Uf
	newEstado.CodIbge = estado.CodIbge

	return &newEstado, nil
}

func (r *EstadoRepositoryMysql) GetEstadoByUF(uf string) (*entity.Estado, error) {
	ctx := context.Background()
	// log.Printf(fmt.Sprint(r.DB.Stats().OpenConnections, r.DB.Stats().InUse, r.DB.Stats().Idle, r.DB.Stats().MaxOpenConnections))
	queries := database.New(r.DB)

	estado, err := queries.GetEstadoByUF(ctx, uf)
	if err != nil {
		return nil, err
	}

	var newEstado entity.Estado
	newEstado.ID = estado.ID
	newEstado.Descricao = estado.Descricao
	newEstado.UF = estado.Uf
	newEstado.CodIbge = estado.CodIbge

	return &newEstado, nil
}

func (r *EstadoRepositoryMysql) Delete(id string) error {
	ctx := context.Background()

	queries := database.New(r.DB)

	err := queries.DeleteEstado(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
