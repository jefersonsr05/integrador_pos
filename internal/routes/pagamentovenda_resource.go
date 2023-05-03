package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
	usecase_pagamentovenda "github.com/jefersonsr05/integrador_pos/internal/usecase/pagamentovenda"
)

func RouterPagamentoVenda(r chi.Router) {

	dbConn, _ := db.Conectar()

	log.Printf("iniciando rotas de PAGAMENTOS DA VENDA")
	repositoryPagamentoVenda := repository.NewPagamentoVendaRepositoryMysql(dbConn)
	createPagamentoVendaUsecase := usecase_pagamentovenda.NewCreatePagamentoVendaUseCase(repositoryPagamentoVenda)
	listPagamentoVendasUsecase := usecase_pagamentovenda.NewListPagamentoVendaUseCase(repositoryPagamentoVenda)
	getPagamentoVendaUseCase := usecase_pagamentovenda.NewGetPagamentoVendaUseCase(repositoryPagamentoVenda)
	deletePagamentoVendaUseCase := usecase_pagamentovenda.NewDeletePagamentoVendaUseCase(repositoryPagamentoVenda)
	updatePagamentoVendaUseCase := usecase_pagamentovenda.NewUpdatePagamentoVendaUseCase(repositoryPagamentoVenda)
	PagamentoVendaHandlers := web.NewPagamentoVendaHandlers(
		createPagamentoVendaUsecase,
		listPagamentoVendasUsecase,
		deletePagamentoVendaUseCase,
		getPagamentoVendaUseCase,
		updatePagamentoVendaUseCase)

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
