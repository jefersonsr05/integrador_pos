package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	usecase "github.com/jefersonsr05/integrador_pos/internal/usecase/venda"
	vendadto "github.com/jefersonsr05/integrador_pos/internal/usecase/venda/dto"
)

type VendaHandlers struct {
	CreateVendaUseCase *usecase.CreateVendaUseCase
	ListVendaUseCase   *usecase.ListVendaUseCase
	DeleteVendaUseCase *usecase.DeleteVendaUseCase
	GetVendaUseCase    *usecase.GetVendaUseCase
	UpdateVendaUseCase *usecase.UpdateVendaUseCase
}

func NewVendaHandlers(
	createVendaUseCase *usecase.CreateVendaUseCase,
	listVendaUseCase *usecase.ListVendaUseCase,
	deleteVendaUseCase *usecase.DeleteVendaUseCase,
	getVendaUseCase *usecase.GetVendaUseCase,
	updateVendaUseCase *usecase.UpdateVendaUseCase) *VendaHandlers {
	return &VendaHandlers{
		CreateVendaUseCase: createVendaUseCase,
		ListVendaUseCase:   listVendaUseCase,
		DeleteVendaUseCase: deleteVendaUseCase,
		GetVendaUseCase:    getVendaUseCase,
		UpdateVendaUseCase: updateVendaUseCase}
}

func (p *VendaHandlers) UpdateVendaHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := p.GetVendaUseCase.GetVendaByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	var input vendadto.VendaInputDTO
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.UpdateVendaUseCase.Execute(id, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *VendaHandlers) CreateVendaHandler(w http.ResponseWriter, r *http.Request) {
	var input vendadto.VendaInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.CreateVendaUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *VendaHandlers) ListVendaHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListVendaUseCase.Execute()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (p *VendaHandlers) DeleteVendaHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := p.GetVendaUseCase.GetVendaByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = p.DeleteVendaUseCase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *VendaHandlers) GetVendaHandler(w http.ResponseWriter, r *http.Request) {

	var id string
	var empresa string

	if chi.URLParam(r, "id") != "" {
		id = chi.URLParam(r, "id")
		//log.Printf("parametros ID:" + chi.URLParam(r, "id"))

		output, err := p.GetVendaUseCase.GetVendaByID(id)
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

		output, err := p.GetVendaUseCase.GetVendaByEmpresa(empresa)
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

		output, err := p.GetVendaUseCase.GetVendaByStatus(status)
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
