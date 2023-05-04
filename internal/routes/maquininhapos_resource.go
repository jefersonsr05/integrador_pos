package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
)

func RouterMaquininhaPOS(r chi.Router) {

	log.Printf("iniciando rotas de MAQUININHA_POS")

	MaquininhaPosHandlers := web.NewMaquininhaPosHandlers(
	// createMaquininhaPosUsecase,
	// listMaquininhaPossUsecase,
	// deleteMaquininhaPosUseCase,
	// getMaquininhaPosUseCase,
	// updateMaquininhaPosUseCase
	)

	r.Post("/", MaquininhaPosHandlers.CreateMaquininhaPosHandler)
	r.Get("/", MaquininhaPosHandlers.ListMaquininhaPosHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Delete("/", MaquininhaPosHandlers.DeleteMaquininhaPosHandler)
		r.Get("/", MaquininhaPosHandlers.GetMaquininhaPosHandler)
		r.Put("/", MaquininhaPosHandlers.UpdateMaquininhaPosHandler)
	})
	r.Get("/codigomc/{codigomc}", MaquininhaPosHandlers.GetMaquininhaPosHandler)
	r.Get("/empresa/{empresa}", MaquininhaPosHandlers.GetMaquininhaPosHandler)

}
