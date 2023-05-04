package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	usecase_produto "github.com/jefersonsr05/integrador_pos/internal/usecase/produto"
	produtodto "github.com/jefersonsr05/integrador_pos/internal/usecase/produto/dto"
)

type ProdutoHandlers struct {
	CreateProdutoUseCase *usecase_produto.CreateProdutoUseCase
	ListProdutoUseCase   *usecase_produto.ListProdutoUseCase
	DeleteProdutoUseCase *usecase_produto.DeleteProdutoUseCase
	GetProdutoUseCase    *usecase_produto.GetProdutoUseCase
	UpdateProdutoUseCase *usecase_produto.UpdateProdutoUseCase
}

func NewProdutoHandlers(
// createProdutoUseCase *usecase.CreateProdutoUseCase,
// listProdutoUseCase *usecase.ListProdutoUseCase,
// deleteProdutoUseCase *usecase.DeleteProdutoUseCase,
// getProdutoUseCase *usecase.GetProdutoUseCase,
// updateProdutoUseCase *usecase.UpdateProdutoUseCase
) *ProdutoHandlers {
	dbConn, _ := db.Conectar()
	repositoryProduto := repository.NewProdutoRepositoryMysql(dbConn)
	createProdutoUseCase := usecase_produto.NewCreateProdutoUseCase(repositoryProduto)
	listProdutoUseCase := usecase_produto.NewListProdutoUseCase(repositoryProduto)
	getProdutoUseCase := usecase_produto.NewGetProdutoUseCase(repositoryProduto)
	deleteProdutoUseCase := usecase_produto.NewDeleteProdutoUseCase(repositoryProduto)
	updateProdutoUseCase := usecase_produto.NewUpdateProdutoUseCase(repositoryProduto)

	return &ProdutoHandlers{
		CreateProdutoUseCase: createProdutoUseCase,
		ListProdutoUseCase:   listProdutoUseCase,
		DeleteProdutoUseCase: deleteProdutoUseCase,
		GetProdutoUseCase:    getProdutoUseCase,
		UpdateProdutoUseCase: updateProdutoUseCase}
}

func (p *ProdutoHandlers) UpdateProdutoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err := p.GetProdutoUseCase.GetProdutoByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	var input produtodto.ProdutoInputDTO
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	output, err := p.UpdateProdutoUseCase.Execute(id, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *ProdutoHandlers) CreateProdutoHandler(w http.ResponseWriter, r *http.Request) {
	var input produtodto.ProdutoInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	output, err := p.CreateProdutoUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *ProdutoHandlers) ListProdutoHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListProdutoUseCase.Execute()
	if err != nil {
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (p *ProdutoHandlers) DeleteProdutoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err := p.GetProdutoUseCase.GetProdutoByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = p.DeleteProdutoUseCase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *ProdutoHandlers) GetProdutoHandler(w http.ResponseWriter, r *http.Request) {

	var id string
	var empresa string

	if chi.URLParam(r, "id") != "" {
		id = chi.URLParam(r, "id")
		//log.Printf("parametros ID:" + chi.URLParam(r, "id"))

		output, err := p.GetProdutoUseCase.GetProdutoByID(id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(output)
	} else if chi.URLParam(r, "empresa") != "" {
		empresa = chi.URLParam(r, "empresa")
		//log.Printf("parametros Empresa:" + chi.URLParam(r, "empresa"))

		output, err := p.GetProdutoUseCase.GetProdutoByEmpresa(empresa)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(output)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("opção inválida")
	}

}
