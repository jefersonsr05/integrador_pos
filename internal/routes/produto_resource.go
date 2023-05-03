package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
	usecase_produto "github.com/jefersonsr05/integrador_pos/internal/usecase/produto"
)

func RouterProduto(r chi.Router) {

	dbConn, _ := db.Conectar()

	log.Printf("iniciando rotas de PRODUTOS")
	repositoryProduto := repository.NewProdutoRepositoryMysql(dbConn)
	createProdutoUsecase := usecase_produto.NewCreateProdutoUseCase(repositoryProduto)
	listProdutosUsecase := usecase_produto.NewListProdutoUseCase(repositoryProduto)
	getProdutoUseCase := usecase_produto.NewGetProdutoUseCase(repositoryProduto)
	deleteProdutoUseCase := usecase_produto.NewDeleteProdutoUseCase(repositoryProduto)
	updateProdutoUseCase := usecase_produto.NewUpdateProdutoUseCase(repositoryProduto)
	ProdutoHandlers := web.NewProdutoHandlers(
		createProdutoUsecase,
		listProdutosUsecase,
		deleteProdutoUseCase,
		getProdutoUseCase,
		updateProdutoUseCase)

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
