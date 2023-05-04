package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
)

func RouterVenda(r chi.Router) {

	log.Printf("iniciando rotas de VENDAS")

	VendaHandlers := web.NewVendaHandlers(
	// createVendaUsecase,
	// listVendasUsecase,
	// deleteVendaUseCase,
	// getVendaUseCase,
	// updateVendaUseCase
	)

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
