package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
)

func RouterPagamentoVenda(r chi.Router) {

	log.Printf("iniciando rotas de PAGAMENTOS DA VENDA")

	PagamentoVendaHandlers := web.NewPagamentoVendaHandlers(
	// createPagamentoVendaUsecase,
	// listPagamentoVendasUsecase,
	// deletePagamentoVendaUseCase,
	// getPagamentoVendaUseCase,
	// updatePagamentoVendaUseCase
	)

	r.Post("/", PagamentoVendaHandlers.CreatePagamentoVendaHandler)
	r.Get("/", PagamentoVendaHandlers.ListPagamentoVendaHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Delete("/", PagamentoVendaHandlers.DeletePagamentoVendaHandler)
		r.Get("/", PagamentoVendaHandlers.GetPagamentoVendaHandler)
		r.Put("/", PagamentoVendaHandlers.UpdatePagamentoVendaHandler)
	})
	r.Get("/codigomc/{codigomc}", PagamentoVendaHandlers.GetPagamentoVendaHandler)
	r.Get("/empresa/{empresa}", PagamentoVendaHandlers.GetPagamentoVendaHandler)

}
