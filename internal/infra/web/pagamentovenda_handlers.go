package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	usecase_pagamentovenda "github.com/jefersonsr05/integrador_pos/internal/usecase/pagamentovenda"
	pagamentovendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/pagamentovenda/dto"
)

type PagamentoVendaHandlers struct {
	CreatePagamentoVendaUseCase *usecase_pagamentovenda.CreatePagamentoVendaUseCase
	ListPagamentoVendaUseCase   *usecase_pagamentovenda.ListPagamentoVendaUseCase
	DeletePagamentoVendaUseCase *usecase_pagamentovenda.DeletePagamentoVendaUseCase
	GetPagamentoVendaUseCase    *usecase_pagamentovenda.GetPagamentoVendaUseCase
	UpdatePagamentoVendaUseCase *usecase_pagamentovenda.UpdatePagamentoVendaUseCase
}

func NewPagamentoVendaHandlers(
// createPagamentoVendaUseCase *usecase.CreatePagamentoVendaUseCase,
// listPagamentoVendaUseCase *usecase.ListPagamentoVendaUseCase,
// deletePagamentoVendaUseCase *usecase.DeletePagamentoVendaUseCase,
// getPagamentoVendaUseCase *usecase.GetPagamentoVendaUseCase,
// updatePagamentoVendaUseCase *usecase.UpdatePagamentoVendaUseCase
) *PagamentoVendaHandlers {

	dbConn, _ := db.Conectar()
	repositoryPagamentoVenda := repository.NewPagamentoVendaRepositoryMysql(dbConn)
	createPagamentoVendaUseCase := usecase_pagamentovenda.NewCreatePagamentoVendaUseCase(repositoryPagamentoVenda)
	listPagamentoVendaUseCase := usecase_pagamentovenda.NewListPagamentoVendaUseCase(repositoryPagamentoVenda)
	getPagamentoVendaUseCase := usecase_pagamentovenda.NewGetPagamentoVendaUseCase(repositoryPagamentoVenda)
	deletePagamentoVendaUseCase := usecase_pagamentovenda.NewDeletePagamentoVendaUseCase(repositoryPagamentoVenda)
	updatePagamentoVendaUseCase := usecase_pagamentovenda.NewUpdatePagamentoVendaUseCase(repositoryPagamentoVenda)

	return &PagamentoVendaHandlers{
		CreatePagamentoVendaUseCase: createPagamentoVendaUseCase,
		ListPagamentoVendaUseCase:   listPagamentoVendaUseCase,
		DeletePagamentoVendaUseCase: deletePagamentoVendaUseCase,
		GetPagamentoVendaUseCase:    getPagamentoVendaUseCase,
		UpdatePagamentoVendaUseCase: updatePagamentoVendaUseCase}
}

func (p *PagamentoVendaHandlers) UpdatePagamentoVendaHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err := p.GetPagamentoVendaUseCase.GetPagamentoVendaByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	var input pagamentovendadto.PagamentoVendaInputDTO
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	output, err := p.UpdatePagamentoVendaUseCase.Execute(id, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *PagamentoVendaHandlers) CreatePagamentoVendaHandler(w http.ResponseWriter, r *http.Request) {
	var input pagamentovendadto.PagamentoVendaInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	output, err := p.CreatePagamentoVendaUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *PagamentoVendaHandlers) ListPagamentoVendaHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListPagamentoVendaUseCase.Execute()
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

func (p *PagamentoVendaHandlers) DeletePagamentoVendaHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err := p.GetPagamentoVendaUseCase.GetPagamentoVendaByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = p.DeletePagamentoVendaUseCase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *PagamentoVendaHandlers) GetPagamentoVendaHandler(w http.ResponseWriter, r *http.Request) {

	var id string
	var vendaid string

	if chi.URLParam(r, "id") != "" {
		id = chi.URLParam(r, "id")
		//log.Printf("parametros ID:" + chi.URLParam(r, "id"))

		output, err := p.GetPagamentoVendaUseCase.GetPagamentoVendaByID(id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(output)
	} else if chi.URLParam(r, "status") != "" {
		status, err := strconv.ParseBool(chi.URLParam(r, "status"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		output, err := p.GetPagamentoVendaUseCase.GetPagamentoVendaByStatus(status)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(output)
	} else if chi.URLParam(r, "vendaid") != "" {
		vendaid = chi.URLParam(r, "vendaid")
		//log.Printf("parametros vendaid:" + chi.URLParam(r, "vendaid"))

		output, err := p.GetPagamentoVendaUseCase.GetPagamentoVendaByVendaID(vendaid)
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
