package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
	usecase_venda "github.com/jefersonsr05/integrador_pos/internal/usecase/venda"
)

func RouterVenda(r chi.Router) {

	dbConn, _ := db.Conectar()

	log.Printf("iniciando rotas de VENDAS")
	repositoryVenda := repository.NewVendaRepositoryMysql(dbConn)
	createVendaUsecase := usecase_venda.NewCreateVendaUseCase(repositoryVenda)
	listVendasUsecase := usecase_venda.NewListVendaUseCase(repositoryVenda)
	getVendaUseCase := usecase_venda.NewGetVendaUseCase(repositoryVenda)
	deleteVendaUseCase := usecase_venda.NewDeleteVendaUseCase(repositoryVenda)
	updateVendaUseCase := usecase_venda.NewUpdateVendaUseCase(repositoryVenda)
	VendaHandlers := web.NewVendaHandlers(
		createVendaUsecase,
		listVendasUsecase,
		deleteVendaUseCase,
		getVendaUseCase,
		updateVendaUseCase)

	r.Post("/", VendaHandlers.CreateVendaHandler)
	r.Get("/", VendaHandlers.ListVendaHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Delete("/", VendaHandlers.DeleteVendaHandler)
		r.Get("/", VendaHandlers.GetVendaHandler)
		r.Put("/", VendaHandlers.UpdateVendaHandler)
	})
	r.Get("/codigomc/{codigomc}", VendaHandlers.GetVendaHandler)
	r.Get("/empresa/{empresa}", VendaHandlers.GetVendaHandler)

}
