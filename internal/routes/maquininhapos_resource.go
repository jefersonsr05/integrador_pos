package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
	usecase_maquininhapos "github.com/jefersonsr05/integrador_pos/internal/usecase/maquininhapos"
)

func RouterMaquininhaPOS(r chi.Router) {

	dbConn, _ := db.Conectar()

	log.Printf("iniciando rotas de MAQUININHA_POS")
	repositoryMaquininhaPos := repository.NewMaquininhaPosRepositoryMysql(dbConn)
	createMaquininhaPosUsecase := usecase_maquininhapos.NewCreateMaquininhaPosUseCase(repositoryMaquininhaPos)
	listMaquininhaPossUsecase := usecase_maquininhapos.NewListMaquininhaPosUseCase(repositoryMaquininhaPos)
	getMaquininhaPosUseCase := usecase_maquininhapos.NewGetMaquininhaPosUseCase(repositoryMaquininhaPos)
	deleteMaquininhaPosUseCase := usecase_maquininhapos.NewDeleteMaquininhaPosUseCase(repositoryMaquininhaPos)
	updateMaquininhaPosUseCase := usecase_maquininhapos.NewUpdateMaquininhaPosUseCase(repositoryMaquininhaPos)
	MaquininhaPosHandlers := web.NewMaquininhaPosHandlers(
		createMaquininhaPosUsecase,
		listMaquininhaPossUsecase,
		deleteMaquininhaPosUseCase,
		getMaquininhaPosUseCase,
		updateMaquininhaPosUseCase)

	r.Post("/", MaquininhaPosHandlers.CreateMaquininhaPosHandler)
	r.Get("/", MaquininhaPosHandlers.ListMaquininhaPosHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Delete("/", MaquininhaPosHandlers.DeleteMaquininhaPosHandler)
		r.Get("/", MaquininhaPosHandlers.GetMaquininhaPosHandler)
		r.Put("/", MaquininhaPosHandlers.UpdateMaquininhaPosHandler)
	})
	r.Get("/codigomc/{codigomc}", MaquininhaPosHandlers.GetMaquininhaPosHandler)
	r.Get("/empresa/{empresa}", MaquininhaPosHandlers.GetMaquininhaPosHandler)

}
