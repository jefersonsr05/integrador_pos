package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	usecase_estado "github.com/jefersonsr05/integrador_pos/internal/usecase/estado"
	estadodto "github.com/jefersonsr05/integrador_pos/internal/usecase/estado/dto"
)

type EstadoHandlers struct {
	CreateEstadoUseCase *usecase_estado.CreateEstadoUseCase
	ListEstadoUseCase   *usecase_estado.ListEstadoUseCase
	DeleteEstadoUseCase *usecase_estado.DeleteEstadoUseCase
	GetEstadoUseCase    *usecase_estado.GetEstadoUseCase
	UpdateEstadoUseCase *usecase_estado.UpdateEstadoUseCase
}

func NewEstadoHandlers(
// createEstadoUseCase *usecase.CreateEstadoUseCase,
// listEstadoUseCase *usecase.ListEstadoUseCase,
// deleteEstadoUseCase *usecase.DeleteEstadoUseCase,
// getEstadoUseCase *usecase.GetEstadoUseCase,
// updateEstadoUseCase *usecase.UpdateEstadoUseCase
) *EstadoHandlers {
	dbConn, _ := db.Conectar()

	repositoryEstado := repository.NewEstadoRepositoryMysql(dbConn)
	createEstadoUseCase := usecase_estado.NewCreateEstadoUseCase(repositoryEstado)
	listEstadoUseCase := usecase_estado.NewListEstadoUseCase(repositoryEstado)
	getEstadoUseCase := usecase_estado.NewGetEstadoUseCase(repositoryEstado)
	deleteEstadoUseCase := usecase_estado.NewDeleteEstadoUseCase(repositoryEstado)
	updateEstadoUseCase := usecase_estado.NewUpdateEstadoUseCase(repositoryEstado)

	return &EstadoHandlers{
		CreateEstadoUseCase: createEstadoUseCase,
		ListEstadoUseCase:   listEstadoUseCase,
		DeleteEstadoUseCase: deleteEstadoUseCase,
		GetEstadoUseCase:    getEstadoUseCase,
		UpdateEstadoUseCase: updateEstadoUseCase}
}

func (p *EstadoHandlers) UpdateEstadoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err := p.GetEstadoUseCase.GetEstadoByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	var input estadodto.EstadoInputDTO
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	output, err := p.UpdateEstadoUseCase.Execute(id, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *EstadoHandlers) CreateEstadoHandler(w http.ResponseWriter, r *http.Request) {
	var input estadodto.EstadoInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	output, err := p.CreateEstadoUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *EstadoHandlers) ListEstadoHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListEstadoUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (p *EstadoHandlers) DeleteEstadoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err := p.GetEstadoUseCase.GetEstadoByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = p.DeleteEstadoUseCase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *EstadoHandlers) GetEstadoHandler(w http.ResponseWriter, r *http.Request) {

	var id string
	var uf string

	if chi.URLParam(r, "id") != "" {
		id = chi.URLParam(r, "id")
		//log.Printf("parametros ID:" + chi.URLParam(r, "id"))

		output, err := p.GetEstadoUseCase.GetEstadoByID(id)
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
		//log.Printf("parametros uf:" + chi.URLParam(r, "uf"))

		output, err := p.GetEstadoUseCase.GetEstadoByUF(uf)
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
		ibge, _ := strconv.ParseInt(chi.URLParam(r, "ibge"), 10, 32)
		//log.Printf("parametros ibge:" + chi.URLParam(r, "ibge"))

		output, err := p.GetEstadoUseCase.GetEstadoByCodIBGE(int32(ibge))
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
