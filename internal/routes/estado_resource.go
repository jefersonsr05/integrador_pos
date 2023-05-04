package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
)

func RouterEstado(r chi.Router) {

	log.Printf("iniciando rotas de ESTADOS")

	EstadoHandlers := web.NewEstadoHandlers(
	// createEstadoUsecase,
	// listEstadosUsecase,
	// deleteEstadoUseCase,
	// getEstadoUseCase,
	// updateEstadoUseCase
	)

	r.Post("/", EstadoHandlers.CreateEstadoHandler)
	r.Get("/", EstadoHandlers.ListEstadoHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Delete("/", EstadoHandlers.DeleteEstadoHandler)
		r.Get("/", EstadoHandlers.GetEstadoHandler)
		r.Put("/", EstadoHandlers.UpdateEstadoHandler)
	})
	r.Get("/uf/{uf}", EstadoHandlers.GetEstadoHandler)
	r.Get("/ibge/{ibge}", EstadoHandlers.GetEstadoHandler)

}
