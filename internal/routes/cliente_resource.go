package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
)

func RouterCliente(r chi.Router) {

	log.Printf("iniciando rotas de CLIENTE")

	clienteHandlers := web.NewClienteHandlers()

	r.Post("/", clienteHandlers.CreateClienteHandler)
	r.Get("/", clienteHandlers.ListClienteHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Delete("/", clienteHandlers.DeleteClienteHandler)
		r.Get("/", clienteHandlers.GetClienteHandler)
		r.Put("/", clienteHandlers.UpdateClienteHandler)
	})
	r.Get("/codigomc/{codigomc}", clienteHandlers.GetClienteHandler)
	r.Get("/empresa/{empresa}", clienteHandlers.GetClienteHandler)

}
