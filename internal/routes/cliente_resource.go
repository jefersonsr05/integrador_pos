package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
	usecase_cliente "github.com/jefersonsr05/integrador_pos/internal/usecase/cliente"
)

func RouterCliente(r chi.Router) {
	dbConn, _ := db.Conectar()

	log.Printf("iniciando rotas de CLIENTE")
	repositoryCliente := repository.NewClienteRepositoryMysql(dbConn)
	createClienteUsecase := usecase_cliente.NewCreateClienteUseCase(repositoryCliente)
	listClientesUsecase := usecase_cliente.NewListClienteUseCase(repositoryCliente)
	getClienteUseCase := usecase_cliente.NewGetClienteUseCase(repositoryCliente)
	deleteClienteUseCase := usecase_cliente.NewDeleteClienteUseCase(repositoryCliente)
	updateClienteUseCase := usecase_cliente.NewUpdateClienteUseCase(repositoryCliente)
	clienteHandlers := web.NewClienteHandlers(
		createClienteUsecase,
		listClientesUsecase,
		deleteClienteUseCase,
		getClienteUseCase,
		updateClienteUseCase)

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
