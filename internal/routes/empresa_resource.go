package api_routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	"github.com/jefersonsr05/integrador_pos/internal/infra/web"
	usecase_empresa "github.com/jefersonsr05/integrador_pos/internal/usecase/empresa"
)

func RouterEmpresa(r chi.Router) {

	dbConn, _ := db.Conectar()

	log.Printf("iniciando rotas de EMPRESA")
	repositoryEmpresa := repository.NewEmpresaRepositoryMysql(dbConn)
	createEmpresaUsecase := usecase_empresa.NewCreateEmpresaUseCase(repositoryEmpresa)
	listEmpresasUsecase := usecase_empresa.NewListEmpresasUseCase(repositoryEmpresa)
	getEmpresaUseCase := usecase_empresa.NewGetEmpresaUseCase(repositoryEmpresa)
	deleteEmpresaUseCase := usecase_empresa.NewDeleteEmpresaUseCase(repositoryEmpresa)
	updateEmpresaUseCase := usecase_empresa.NewUpdateEmpresaUseCase(repositoryEmpresa)
	empresaHandlers := web.NewEmpresaHandlers(
		createEmpresaUsecase,
		listEmpresasUsecase,
		deleteEmpresaUseCase,
		getEmpresaUseCase,
		updateEmpresaUseCase)

	// r := chi.NewRouter()
	// r.Use() // some middleware..
	// r.Route("/empresas", func(r chi.Router) {
	// r.Use(jwtauth.Verifier(configs.TokenAuth))
	// r.Use(jwtauth.Authenticator)

	r.Post("/", empresaHandlers.CreateEmpresaHandler)
	r.Get("/", empresaHandlers.ListEmpresasHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Delete("/", empresaHandlers.DeleteEmpresaHandler)
		r.Get("/", empresaHandlers.GetEmpresaHandler)
		r.Put("/", empresaHandlers.UpdateEmpresaHandler)
	})
	r.Get("/cnpj/{cnpj}", empresaHandlers.GetEmpresaHandler)
	r.Get("/chave/{chave}", empresaHandlers.GetEmpresaHandler)

	// })

	// return r
}
