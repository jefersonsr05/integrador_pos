package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	usecase_itemvenda "github.com/jefersonsr05/integrador_pos/internal/usecase/itemvenda"
	itemvendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/itemvenda/dto"
)

type ItemVendaHandlers struct {
	CreateItemVendaUseCase *usecase_itemvenda.CreateItemVendaUseCase
	ListItemVendaUseCase   *usecase_itemvenda.ListItemVendaUseCase
	DeleteItemVendaUseCase *usecase_itemvenda.DeleteItemVendaUseCase
	GetItemVendaUseCase    *usecase_itemvenda.GetItemVendaUseCase
	UpdateItemVendaUseCase *usecase_itemvenda.UpdateItemVendaUseCase
}

func NewItemVendaHandlers(
// createItemVendaUseCase *usecase.CreateItemVendaUseCase,
// listItemVendaUseCase *usecase.ListItemVendaUseCase,
// deleteItemVendaUseCase *usecase.DeleteItemVendaUseCase,
// getItemVendaUseCase *usecase.GetItemVendaUseCase,
// updateItemVendaUseCase *usecase.UpdateItemVendaUseCase
) *ItemVendaHandlers {
	dbConn, _ := db.Conectar()

	repositoryItemVenda := repository.NewItemVendaRepositoryMysql(dbConn)
	createItemVendaUseCase := usecase_itemvenda.NewCreateItemVendaUseCase(repositoryItemVenda)
	listItemVendaUseCase := usecase_itemvenda.NewListItemVendaUseCase(repositoryItemVenda)
	getItemVendaUseCase := usecase_itemvenda.NewGetItemVendaUseCase(repositoryItemVenda)
	deleteItemVendaUseCase := usecase_itemvenda.NewDeleteItemVendaUseCase(repositoryItemVenda)
	updateItemVendaUseCase := usecase_itemvenda.NewUpdateItemVendaUseCase(repositoryItemVenda)

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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("opção inválida")
	}

}
