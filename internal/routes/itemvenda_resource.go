package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
)

func RouterItemVenda(r chi.Router) {

	log.Printf("iniciando rotas de ITEM_VENDA")

	ItemVendaHandlers := web.NewItemVendaHandlers(
	// createItemVendaUsecase,
	// listItemVendasUsecase,
	// deleteItemVendaUseCase,
	// getItemVendaUseCase,
	// updateItemVendaUseCase
	)

	r.Post("/", ItemVendaHandlers.CreateItemVendaHandler)
	r.Get("/", ItemVendaHandlers.ListItemVendaHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Delete("/", ItemVendaHandlers.DeleteItemVendaHandler)
		r.Get("/", ItemVendaHandlers.GetItemVendaHandler)
		r.Put("/", ItemVendaHandlers.UpdateItemVendaHandler)
	})
	r.Get("/codigomc/{codigomc}", ItemVendaHandlers.GetItemVendaHandler)
	r.Get("/empresa/{empresa}", ItemVendaHandlers.GetItemVendaHandler)

}
