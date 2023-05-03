package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
	usecase_estado "github.com/jefersonsr05/integrador_pos/internal/usecase/estado"
)

func RouterEstado(r chi.Router) {

	dbConn, _ := db.Conectar()
	log.Printf("iniciando rotas de ESTADOS")
	repositoryEstado := repository.NewEstadoRepositoryMysql(dbConn)
	createEstadoUsecase := usecase_estado.NewCreateEstadoUseCase(repositoryEstado)
	listEstadosUsecase := usecase_estado.NewListEstadoUseCase(repositoryEstado)
	getEstadoUseCase := usecase_estado.NewGetEstadoUseCase(repositoryEstado)
	deleteEstadoUseCase := usecase_estado.NewDeleteEstadoUseCase(repositoryEstado)
	updateEstadoUseCase := usecase_estado.NewUpdateEstadoUseCase(repositoryEstado)
	EstadoHandlers := web.NewEstadoHandlers(
		createEstadoUsecase,
		listEstadosUsecase,
		deleteEstadoUseCase,
		getEstadoUseCase,
		updateEstadoUseCase)

	r.Post("/", EstadoHandlers.CreateEstadoHandler)
	r.Get("/", EstadoHandlers.ListEstadoHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Delete("/", EstadoHandlers.DeleteEstadoHandler)
		r.Get("/", EstadoHandlers.GetEstadoHandler)
		r.Put("/", EstadoHandlers.UpdateEstadoHandler)
	})
	r.Get("/codigomc/{codigomc}", EstadoHandlers.GetEstadoHandler)
	r.Get("/empresa/{empresa}", EstadoHandlers.GetEstadoHandler)

}
