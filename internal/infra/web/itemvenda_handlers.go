package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	usecase "github.com/jefersonsr05/integrador_pos/internal/usecase/itemvenda"
	itemvendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/itemvenda/dto"
)

type ItemVendaHandlers struct {
	CreateItemVendaUseCase *usecase.CreateItemVendaUseCase
	ListItemVendaUseCase   *usecase.ListItemVendaUseCase
	DeleteItemVendaUseCase *usecase.DeleteItemVendaUseCase
	GetItemVendaUseCase    *usecase.GetItemVendaUseCase
	UpdateItemVendaUseCase *usecase.UpdateItemVendaUseCase
}

func NewItemVendaHandlers(
	createItemVendaUseCase *usecase.CreateItemVendaUseCase,
	listItemVendaUseCase *usecase.ListItemVendaUseCase,
	deleteItemVendaUseCase *usecase.DeleteItemVendaUseCase,
	getItemVendaUseCase *usecase.GetItemVendaUseCase,
	updateItemVendaUseCase *usecase.UpdateItemVendaUseCase) *ItemVendaHandlers {
	return &ItemVendaHandlers{
		CreateItemVendaUseCase: createItemVendaUseCase,
		ListItemVendaUseCase:   listItemVendaUseCase,
		DeleteItemVendaUseCase: deleteItemVendaUseCase,
		GetItemVendaUseCase:    getItemVendaUseCase,
		UpdateItemVendaUseCase: updateItemVendaUseCase}
}

func (p *ItemVendaHandlers) UpdateItemVendaHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := p.GetItemVendaUseCase.GetItemVendaByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	var input itemvendadto.ItemVendaInputDTO
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.UpdateItemVendaUseCase.Execute(id, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *ItemVendaHandlers) CreateItemVendaHandler(w http.ResponseWriter, r *http.Request) {
	var input itemvendadto.ItemVendaInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.CreateItemVendaUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *ItemVendaHandlers) ListItemVendaHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListItemVendaUseCase.Execute()
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

func (p *ItemVendaHandlers) DeleteItemVendaHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := p.GetItemVendaUseCase.GetItemVendaByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = p.DeleteItemVendaUseCase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *ItemVendaHandlers) GetItemVendaHandler(w http.ResponseWriter, r *http.Request) {

	var id string
	var vendaid string

	if chi.URLParam(r, "id") != "" {
		id = chi.URLParam(r, "id")
		//log.Printf("parametros ID:" + chi.URLParam(r, "id"))

		output, err := p.GetItemVendaUseCase.GetItemVendaByID(id)
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

		output, err := p.GetItemVendaUseCase.GetItemVendaByVendaID(vendaid)
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
