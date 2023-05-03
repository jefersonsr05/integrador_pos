package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	usecase "github.com/jefersonsr05/integrador_pos/internal/usecase/cliente"
	clientedto "github.com/jefersonsr05/integrador_pos/internal/usecase/cliente/dto"
)

type ClienteHandlers struct {
	CreateClienteUseCase *usecase.CreateClienteUseCase
	ListClienteUseCase   *usecase.ListClienteUseCase
	DeleteClienteUseCase *usecase.DeleteClienteUseCase
	GetClienteUseCase    *usecase.GetClienteUseCase
	UpdateClienteUseCase *usecase.UpdateClienteUseCase
}

func NewClienteHandlers(
	createClienteUseCase *usecase.CreateClienteUseCase,
	listClienteUseCase *usecase.ListClienteUseCase,
	deleteClienteUseCase *usecase.DeleteClienteUseCase,
	getClienteUseCase *usecase.GetClienteUseCase,
	updateClienteUseCase *usecase.UpdateClienteUseCase) *ClienteHandlers {
	return &ClienteHandlers{
		CreateClienteUseCase: createClienteUseCase,
		ListClienteUseCase:   listClienteUseCase,
		DeleteClienteUseCase: deleteClienteUseCase,
		GetClienteUseCase:    getClienteUseCase,
		UpdateClienteUseCase: updateClienteUseCase}
}

func (p *ClienteHandlers) UpdateClienteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := p.GetClienteUseCase.GetClienteByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	var input clientedto.ClienteInputDTO
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.UpdateClienteUseCase.Execute(id, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *ClienteHandlers) CreateClienteHandler(w http.ResponseWriter, r *http.Request) {
	var input clientedto.ClienteInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Printf("Erro Create UseCase" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.CreateClienteUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *ClienteHandlers) ListClienteHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListClienteUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (p *ClienteHandlers) DeleteClienteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := p.GetClienteUseCase.GetClienteByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = p.DeleteClienteUseCase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *ClienteHandlers) GetClienteHandler(w http.ResponseWriter, r *http.Request) {

	var id string
	var empresa string
	// var codigomc string

	if chi.URLParam(r, "id") != "" {
		id = chi.URLParam(r, "id")
		//log.Printf("parametros ID:" + chi.URLParam(r, "id"))

		output, err := p.GetClienteUseCase.GetClienteByID(id)
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

		output, err := p.GetClienteUseCase.GetClienteByEmpresa(empresa)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(output)
		// } else if chi.URLParam(r, "codigomc") != "" {
		// 	codigomc = chi.URLParam(r, "codigomc")
		// 	//log.Printf("parametros codigomc:" + chi.URLParam(r, "codigomc"))

		// 	output, err := p.GetClienteUseCase.GetClienteByCodigoMC(codigomc)
		// 	if err != nil {
		// 		w.Header().Set("Content-Type", "application/json")
		// 		w.WriteHeader(http.StatusNotFound)
		// 		json.NewEncoder(w).Encode(err)
		// 		return
		// 	}
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusOK)
		// 	json.NewEncoder(w).Encode(output)
		// }
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
