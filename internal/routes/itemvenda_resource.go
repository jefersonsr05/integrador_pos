package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
	usecase_itemvenda "github.com/jefersonsr05/integrador_pos/internal/usecase/itemvenda"
)

func RouterItemVenda(r chi.Router) {

	dbConn, _ := db.Conectar()

	log.Printf("iniciando rotas de ITEM_VENDA")
	repositoryItemVenda := repository.NewItemVendaRepositoryMysql(dbConn)
	createItemVendaUsecase := usecase_itemvenda.NewCreateItemVendaUseCase(repositoryItemVenda)
	listItemVendasUsecase := usecase_itemvenda.NewListItemVendaUseCase(repositoryItemVenda)
	getItemVendaUseCase := usecase_itemvenda.NewGetItemVendaUseCase(repositoryItemVenda)
	deleteItemVendaUseCase := usecase_itemvenda.NewDeleteItemVendaUseCase(repositoryItemVenda)
	updateItemVendaUseCase := usecase_itemvenda.NewUpdateItemVendaUseCase(repositoryItemVenda)
	ItemVendaHandlers := web.NewItemVendaHandlers(
		createItemVendaUsecase,
		listItemVendasUsecase,
		deleteItemVendaUseCase,
		getItemVendaUseCase,
		updateItemVendaUseCase)

	r.Post("/", ItemVendaHandlers.CreateItemVendaHandler)
	r.Get("/", ItemVendaHandlers.ListItemVendaHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Delete("/", ItemVendaHandlers.DeleteItemVendaHandler)
		r.Get("/", ItemVendaHandlers.GetItemVendaHandler)
		r.Put("/", ItemVendaHandlers.UpdateItemVendaHandler)
	})
	r.Get("/codigomc/{codigomc}", ItemVendaHandlers.GetItemVendaHandler)
	r.Get("/empresa/{empresa}", ItemVendaHandlers.GetItemVendaHandler)

}
