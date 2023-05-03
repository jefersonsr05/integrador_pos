package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
	usecase_cidade "github.com/jefersonsr05/integrador_pos/internal/usecase/cidade"
)

func RouterCidade(r chi.Router) {

	dbConn, _ := db.Conectar()

	log.Printf("iniciando rotas de CIDADES")
	repositoryCidade := repository.NewCidadeRepositoryMysql(dbConn)
	createCidadeUsecase := usecase_cidade.NewCreateCidadeUseCase(repositoryCidade)
	listCidadesUsecase := usecase_cidade.NewListCidadeUseCase(repositoryCidade)
	getCidadeUseCase := usecase_cidade.NewGetCidadeUseCase(repositoryCidade)
	deleteCidadeUseCase := usecase_cidade.NewDeleteCidadeUseCase(repositoryCidade)
	updateCidadeUseCase := usecase_cidade.NewUpdateCidadeUseCase(repositoryCidade)
	CidadeHandlers := web.NewCidadeHandlers(
		createCidadeUsecase,
		listCidadesUsecase,
		deleteCidadeUseCase,
		getCidadeUseCase,
		updateCidadeUseCase)

	r.Post("/", CidadeHandlers.CreateCidadeHandler)
	r.Get("/", CidadeHandlers.ListCidadeHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Delete("/", CidadeHandlers.DeleteCidadeHandler)
		r.Get("/", CidadeHandlers.GetCidadeHandler)
		r.Put("/", CidadeHandlers.UpdateCidadeHandler)
	})
	r.Get("/codigomc/{codigomc}", CidadeHandlers.GetCidadeHandler)
	r.Get("/empresa/{empresa}", CidadeHandlers.GetCidadeHandler)

}
