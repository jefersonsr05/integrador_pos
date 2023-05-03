package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	usecase "github.com/jefersonsr05/integrador_pos/internal/usecase/cidade"
	cidadedto "github.com/jefersonsr05/integrador_pos/internal/usecase/cidade/dto"
)

type CidadeHandlers struct {
	CreateCidadeUseCase *usecase.CreateCidadeUseCase
	ListCidadeUseCase   *usecase.ListCidadeUseCase
	DeleteCidadeUseCase *usecase.DeleteCidadeUseCase
	GetCidadeUseCase    *usecase.GetCidadeUseCase
	UpdateCidadeUseCase *usecase.UpdateCidadeUseCase
}

func NewCidadeHandlers(
	createCidadeUseCase *usecase.CreateCidadeUseCase,
	listCidadeUseCase *usecase.ListCidadeUseCase,
	deleteCidadeUseCase *usecase.DeleteCidadeUseCase,
	getCidadeUseCase *usecase.GetCidadeUseCase,
	updateCidadeUseCase *usecase.UpdateCidadeUseCase) *CidadeHandlers {
	return &CidadeHandlers{
		CreateCidadeUseCase: createCidadeUseCase,
		ListCidadeUseCase:   listCidadeUseCase,
		DeleteCidadeUseCase: deleteCidadeUseCase,
		GetCidadeUseCase:    getCidadeUseCase,
		UpdateCidadeUseCase: updateCidadeUseCase}
}

func (p *CidadeHandlers) UpdateCidadeHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := p.GetCidadeUseCase.GetCidadeByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	var input cidadedto.CidadeInputDTO
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.UpdateCidadeUseCase.Execute(id, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *CidadeHandlers) CreateCidadeHandler(w http.ResponseWriter, r *http.Request) {
	var input cidadedto.CidadeInputDTO
	// log.Printf("Entrando no cadastro")
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// log.Printf("Decode JSON SUCCESS")
	output, err := p.CreateCidadeUseCase.Execute(input)
	if err != nil {
		// log.Printf("Erro Execute Cadastro JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *CidadeHandlers) ListCidadeHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListCidadeUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (p *CidadeHandlers) DeleteCidadeHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := p.GetCidadeUseCase.GetCidadeByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = p.DeleteCidadeUseCase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *CidadeHandlers) GetCidadeHandler(w http.ResponseWriter, r *http.Request) {

	var id string
	var ibge string
	var uf string

	if chi.URLParam(r, "id") != "" {
		id = chi.URLParam(r, "id")
		//log.Printf("parametros ID:" + chi.URLParam(r, "id"))

		output, err := p.GetCidadeUseCase.GetCidadeByID(id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(output)
	} else if chi.URLParam(r, "ibge") != "" {
		ibge = chi.URLParam(r, "ibge")
		//log.Printf("parametros IBGE:" + chi.URLParam(r, "ibge"))

		output, err := p.GetCidadeUseCase.GetCidadeByIBGE(ibge)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(output)
	} else if chi.URLParam(r, "uf") != "" {
		uf = chi.URLParam(r, "uf")
		//log.Printf("parametros UF:" + chi.URLParam(r, "uf"))
		output, err := p.GetCidadeUseCase.GetCidadeByUF(uf)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(output)
	} else if chi.URLParam(r, "ufibge") != "" {
		ufibge, err := strconv.ParseInt(chi.URLParam(r, "ufibge"), 10, 32)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
		output, err := p.GetCidadeUseCase.GetCidadeByIbgeUF(int32(ufibge))
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
