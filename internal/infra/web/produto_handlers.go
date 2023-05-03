package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	usecase "github.com/jefersonsr05/integrador_pos/internal/usecase/produto"
	produtodto "github.com/jefersonsr05/integrador_pos/internal/usecase/produto/dto"
)

type ProdutoHandlers struct {
	CreateProdutoUseCase *usecase.CreateProdutoUseCase
	ListProdutoUseCase   *usecase.ListProdutoUseCase
	DeleteProdutoUseCase *usecase.DeleteProdutoUseCase
	GetProdutoUseCase    *usecase.GetProdutoUseCase
	UpdateProdutoUseCase *usecase.UpdateProdutoUseCase
}

func NewProdutoHandlers(
	createProdutoUseCase *usecase.CreateProdutoUseCase,
	listProdutoUseCase *usecase.ListProdutoUseCase,
	deleteProdutoUseCase *usecase.DeleteProdutoUseCase,
	getProdutoUseCase *usecase.GetProdutoUseCase,
	updateProdutoUseCase *usecase.UpdateProdutoUseCase) *ProdutoHandlers {
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
		w.WriteHeader(http.StatusBadRequest)
		return
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
		w.WriteHeader(http.StatusBadRequest)
		return
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
		w.WriteHeader(http.StatusBadRequest)
		return
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
		w.WriteHeader(http.StatusBadRequest)
		return
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
