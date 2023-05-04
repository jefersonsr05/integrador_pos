package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
)

func RouterCidade(r chi.Router) {

	// dbConn, _ := db.Conectar()

	log.Printf("iniciando rotas de CIDADES")

	CidadeHandlers := web.NewCidadeHandlers()

	r.Post("/", CidadeHandlers.CreateCidadeHandler)
	r.Get("/", CidadeHandlers.ListCidadeHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Delete("/", CidadeHandlers.DeleteCidadeHandler)
		r.Get("/", CidadeHandlers.GetCidadeHandler)
		r.Put("/", CidadeHandlers.UpdateCidadeHandler)
	})
	r.Get("/uf/{uf}", CidadeHandlers.GetCidadeHandler)
	r.Get("/ibge/{ibge}", CidadeHandlers.GetCidadeHandler)
	r.Get("/ufibge/{ufibge}", CidadeHandlers.GetCidadeHandler)

}
