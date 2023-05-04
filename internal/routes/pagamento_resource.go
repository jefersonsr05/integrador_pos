package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
)

func RouterPagamento(r chi.Router) {

	log.Printf("iniciando rotas de PAGAMENTOS")

	PagamentoHandlers := web.NewPagamentoHandlers(
	// createPagamentoUsecase,
	// listPagamentosUsecase,
	// deletePagamentoUseCase,
	// getPagamentoUseCase,
	// updatePagamentoUseCase
	)

	r.Post("/", PagamentoHandlers.CreatePagamentoHandler)
	r.Get("/", PagamentoHandlers.ListPagamentoHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Delete("/", PagamentoHandlers.DeletePagamentoHandler)
		r.Get("/", PagamentoHandlers.GetPagamentoHandler)
		r.Put("/", PagamentoHandlers.UpdatePagamentoHandler)
	})
	r.Get("/codigomc/{codigomc}", PagamentoHandlers.GetPagamentoHandler)
	r.Get("/empresa/{empresa}", PagamentoHandlers.GetPagamentoHandler)

}
