package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	usecase "github.com/jefersonsr05/integrador_pos/internal/usecase/pagamento"
	pagamentodto "github.com/jefersonsr05/integrador_pos/internal/usecase/pagamento/dto"
)

type PagamentoHandlers struct {
	CreatePagamentoUseCase *usecase.CreatePagamentoUseCase
	ListPagamentoUseCase   *usecase.ListPagamentoUseCase
	DeletePagamentoUseCase *usecase.DeletePagamentoUseCase
	GetPagamentoUseCase    *usecase.GetPagamentoUseCase
	UpdatePagamentoUseCase *usecase.UpdatePagamentoUseCase
}

func NewPagamentoHandlers(
	createPagamentoUseCase *usecase.CreatePagamentoUseCase,
	listPagamentoUseCase *usecase.ListPagamentoUseCase,
	deletePagamentoUseCase *usecase.DeletePagamentoUseCase,
	getPagamentoUseCase *usecase.GetPagamentoUseCase,
	updatePagamentoUseCase *usecase.UpdatePagamentoUseCase) *PagamentoHandlers {
	return &PagamentoHandlers{
		CreatePagamentoUseCase: createPagamentoUseCase,
		ListPagamentoUseCase:   listPagamentoUseCase,
		DeletePagamentoUseCase: deletePagamentoUseCase,
		GetPagamentoUseCase:    getPagamentoUseCase,
		UpdatePagamentoUseCase: updatePagamentoUseCase}
}

func (p *PagamentoHandlers) UpdatePagamentoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := p.GetPagamentoUseCase.GetPagamentoByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	var input pagamentodto.PagamentoInputDTO
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.UpdatePagamentoUseCase.Execute(id, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *PagamentoHandlers) CreatePagamentoHandler(w http.ResponseWriter, r *http.Request) {
	var input pagamentodto.PagamentoInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.CreatePagamentoUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *PagamentoHandlers) ListPagamentoHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListPagamentoUseCase.Execute()
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

func (p *PagamentoHandlers) DeletePagamentoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := p.GetPagamentoUseCase.GetPagamentoByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = p.DeletePagamentoUseCase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *PagamentoHandlers) GetPagamentoHandler(w http.ResponseWriter, r *http.Request) {

	var id string
	var empresa string

	if chi.URLParam(r, "id") != "" {
		id = chi.URLParam(r, "id")
		//log.Printf("parametros ID:" + chi.URLParam(r, "id"))

		output, err := p.GetPagamentoUseCase.GetPagamentoByID(id)
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

		output, err := p.GetPagamentoUseCase.GetPagamentoByEmpresa(empresa)
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
