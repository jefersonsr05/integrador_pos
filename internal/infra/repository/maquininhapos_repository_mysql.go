package repository

import (
	"context"
	"database/sql"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	database "github.com/jefersonsr05/integrador_pos/internal/infra/repository/db"
)

type MaquininhaPosRepositoryMysql struct {
	DB *sql.DB
}

func NewMaquininhaPosRepositoryMysql(db *sql.DB) *MaquininhaPosRepositoryMysql {
	return &MaquininhaPosRepositoryMysql{DB: db}
}

func (r *MaquininhaPosRepositoryMysql) Create(maquininhapos *entity.MaquininhaPos) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.CreateMaquininhaPOS(ctx, database.CreateMaquininhaPOSParams{
		ID:             maquininhapos.ID,
		IDEmpresa:      maquininhapos.EmpresaID,
		Administradora: sql.NullString{String: maquininhapos.Administradora, Valid: true},
		Cnpj:           sql.NullString{String: maquininhapos.Cnpj, Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *MaquininhaPosRepositoryMysql) FindAll() ([]*entity.MaquininhaPos, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listMaquininhaPos, err := queries.ListMaquininhaPOS(ctx)
	if err != nil {
		return nil, err
	}

	var maquininhas []*entity.MaquininhaPos
	for _, maquininha := range listMaquininhaPos {
		var newMaquininha entity.MaquininhaPos
		newMaquininha.ID = maquininha.ID
		newMaquininha.EmpresaID = maquininha.IDEmpresa
		newMaquininha.Administradora = maquininha.Administradora.String
		newMaquininha.Cnpj = maquininha.Cnpj.String

		maquininhas = append(maquininhas, &newMaquininha)
	}
	return maquininhas, nil
}

func (r *MaquininhaPosRepositoryMysql) GetMaquininhasPosByEmpresa(empresaid string) ([]*entity.MaquininhaPos, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	listMaquininhaPos, err := queries.GetMaquininhasPosByEmpresa(ctx, empresaid)
	if err != nil {
		return nil, err
	}

	var maquininhas []*entity.MaquininhaPos
	for _, maquininha := range listMaquininhaPos {
		var newMaquininha entity.MaquininhaPos
		newMaquininha.ID = maquininha.ID
		newMaquininha.EmpresaID = maquininha.IDEmpresa
		newMaquininha.Administradora = maquininha.Administradora.String
		newMaquininha.Cnpj = maquininha.Cnpj.String

		maquininhas = append(maquininhas, &newMaquininha)
	}
	return maquininhas, nil
}

func (r *MaquininhaPosRepositoryMysql) Update(maquininhapos *entity.MaquininhaPos) error {
	ctx := context.Background()
	queries := database.New(r.DB)

	err := queries.UpdateMaquininhaPOS(ctx, database.UpdateMaquininhaPOSParams{
		ID:             maquininhapos.ID,
		IDEmpresa:      maquininhapos.EmpresaID,
		Administradora: sql.NullString{String: maquininhapos.Administradora, Valid: true},
		Cnpj:           sql.NullString{String: maquininhapos.Cnpj, Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *MaquininhaPosRepositoryMysql) GetMaquininhasPOS(id string) (*entity.MaquininhaPos, error) {
	ctx := context.Background()

	queries := database.New(r.DB)

	// log.Printf("passei aqui com o ID: " + id)

	maquininha, err := queries.GetMaquininhasPOS(ctx, id)
	if err != nil {
		return nil, err
	}

	var newMaquininha entity.MaquininhaPos
	newMaquininha.ID = maquininha.ID
	newMaquininha.EmpresaID = maquininha.IDEmpresa
	newMaquininha.Administradora = maquininha.Administradora.String
	newMaquininha.Cnpj = maquininha.Cnpj.String

	return &newMaquininha, nil
}

func (r *MaquininhaPosRepositoryMysql) Delete(id string) error {
	ctx := context.Background()

	queries := database.New(r.DB)

	err := queries.DeleteMaquininhaPOS(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
