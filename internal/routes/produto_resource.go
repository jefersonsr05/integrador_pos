package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
)

func RouterProduto(r chi.Router) {

	log.Printf("iniciando rotas de PRODUTOS")

	ProdutoHandlers := web.NewProdutoHandlers(
	// createProdutoUsecase,
	// listProdutosUsecase,
	// deleteProdutoUseCase,
	// getProdutoUseCase,
	// updateProdutoUseCase
	)

	r.Post("/", ProdutoHandlers.CreateProdutoHandler)
	r.Get("/", ProdutoHandlers.ListProdutoHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Delete("/", ProdutoHandlers.DeleteProdutoHandler)
		r.Get("/", ProdutoHandlers.GetProdutoHandler)
		r.Put("/", ProdutoHandlers.UpdateProdutoHandler)
	})
	r.Get("/codigomc/{codigomc}", ProdutoHandlers.GetProdutoHandler)
	r.Get("/empresa/{empresa}", ProdutoHandlers.GetProdutoHandler)

}
