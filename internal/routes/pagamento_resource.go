package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
	usecase_pagamento "github.com/jefersonsr05/integrador_pos/internal/usecase/pagamento"
)

func RouterPagamento(r chi.Router) {

	dbConn, _ := db.Conectar()

	log.Printf("iniciando rotas de PAGAMENTOS")
	repositoryPagamento := repository.NewPagamentoRepositoryMysql(dbConn)
	createPagamentoUsecase := usecase_pagamento.NewCreatePagamentoUseCase(repositoryPagamento)
	listPagamentosUsecase := usecase_pagamento.NewListPagamentoUseCase(repositoryPagamento)
	getPagamentoUseCase := usecase_pagamento.NewGetPagamentoUseCase(repositoryPagamento)
	deletePagamentoUseCase := usecase_pagamento.NewDeletePagamentoUseCase(repositoryPagamento)
	updatePagamentoUseCase := usecase_pagamento.NewUpdatePagamentoUseCase(repositoryPagamento)
	PagamentoHandlers := web.NewPagamentoHandlers(
		createPagamentoUsecase,
		listPagamentosUsecase,
		deletePagamentoUseCase,
		getPagamentoUseCase,
		updatePagamentoUseCase)

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
