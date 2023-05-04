package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	usecase_maquininhapos "github.com/jefersonsr05/integrador_pos/internal/usecase/maquininhapos"
	maquininhaposdto "github.com/jefersonsr05/integrador_pos/internal/usecase/maquininhapos/dto"
)

type MaquininhaPosHandlers struct {
	CreateMaquininhaPosUseCase *usecase_maquininhapos.CreateMaquininhaPosUseCase
	ListMaquininhaPosUseCase   *usecase_maquininhapos.ListMaquininhaPosUseCase
	DeleteMaquininhaPosUseCase *usecase_maquininhapos.DeleteMaquininhaPosUseCase
	GetMaquininhaPosUseCase    *usecase_maquininhapos.GetMaquininhaPosUseCase
	UpdateMaquininhaPosUseCase *usecase_maquininhapos.UpdateMaquininhaPosUseCase
}

func NewMaquininhaPosHandlers(
// createMaquininhaPosUseCase *usecase.CreateMaquininhaPosUseCase,
// listMaquininhaPosUseCase *usecase.ListMaquininhaPosUseCase,
// deleteMaquininhaPosUseCase *usecase.DeleteMaquininhaPosUseCase,
// getMaquininhaPosUseCase *usecase.GetMaquininhaPosUseCase,
// updateMaquininhaPosUseCase *usecase.UpdateMaquininhaPosUseCase
) *MaquininhaPosHandlers {
	dbConn, _ := db.Conectar()

	repositoryMaquininhaPos := repository.NewMaquininhaPosRepositoryMysql(dbConn)
	createMaquininhaPosUseCase := usecase_maquininhapos.NewCreateMaquininhaPosUseCase(repositoryMaquininhaPos)
	listMaquininhaPosUseCase := usecase_maquininhapos.NewListMaquininhaPosUseCase(repositoryMaquininhaPos)
	getMaquininhaPosUseCase := usecase_maquininhapos.NewGetMaquininhaPosUseCase(repositoryMaquininhaPos)
	deleteMaquininhaPosUseCase := usecase_maquininhapos.NewDeleteMaquininhaPosUseCase(repositoryMaquininhaPos)
	updateMaquininhaPosUseCase := usecase_maquininhapos.NewUpdateMaquininhaPosUseCase(repositoryMaquininhaPos)

	return &MaquininhaPosHandlers{
		CreateMaquininhaPosUseCase: createMaquininhaPosUseCase,
		ListMaquininhaPosUseCase:   listMaquininhaPosUseCase,
		DeleteMaquininhaPosUseCase: deleteMaquininhaPosUseCase,
		GetMaquininhaPosUseCase:    getMaquininhaPosUseCase,
		UpdateMaquininhaPosUseCase: updateMaquininhaPosUseCase}
}

func (p *MaquininhaPosHandlers) UpdateMaquininhaPosHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err := p.GetMaquininhaPosUseCase.GetMaquininhaPOS(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	var input maquininhaposdto.MaquininhaPosInputDTO
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	output, err := p.UpdateMaquininhaPosUseCase.Execute(id, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *MaquininhaPosHandlers) CreateMaquininhaPosHandler(w http.ResponseWriter, r *http.Request) {
	var input maquininhaposdto.MaquininhaPosInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	output, err := p.CreateMaquininhaPosUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *MaquininhaPosHandlers) ListMaquininhaPosHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListMaquininhaPosUseCase.Execute()
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

func (p *MaquininhaPosHandlers) DeleteMaquininhaPosHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err := p.GetMaquininhaPosUseCase.GetMaquininhaPOS(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = p.DeleteMaquininhaPosUseCase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *MaquininhaPosHandlers) GetMaquininhaPosHandler(w http.ResponseWriter, r *http.Request) {

	var id string
	var empresa string

	if chi.URLParam(r, "id") != "" {
		id = chi.URLParam(r, "id")
		//log.Printf("parametros ID:" + chi.URLParam(r, "id"))

		output, err := p.GetMaquininhaPosUseCase.GetMaquininhaPOS(id)
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

		output, err := p.GetMaquininhaPosUseCase.GetMaquininhaPosByEmpresa(empresa)
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
