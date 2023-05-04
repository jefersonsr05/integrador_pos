package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	usecase_empresa "github.com/jefersonsr05/integrador_pos/internal/usecase/empresa"
	empresadto "github.com/jefersonsr05/integrador_pos/internal/usecase/empresa/dto"
)

type EmpresaHandlers struct {
	CreateEmpresaUseCase *usecase_empresa.CreateEmpresaUseCase
	ListEmpresasUseCase  *usecase_empresa.ListEmpresasUseCase
	DeleteEmpresaUseCase *usecase_empresa.DeleteEmpresaUseCase
	GetEmpresaUseCase    *usecase_empresa.GetEmpresaUseCase
	UpdateEmpresaUseCase *usecase_empresa.UpdateEmpresaUseCase
}

func NewEmpresaHandlers(
// createEmpresaUseCase *usecase.CreateEmpresaUseCase,
// listEmpresasUseCase *usecase.ListEmpresasUseCase,
// deleteEmpresaUseCase *usecase.DeleteEmpresaUseCase,
// getEmpresaUseCase *usecase.GetEmpresaUseCase,
// updateEmpresaUseCase *usecase.UpdateEmpresaUseCase
) *EmpresaHandlers {
	dbConn, _ := db.Conectar()

	repositoryEmpresa := repository.NewEmpresaRepositoryMysql(dbConn)
	createEmpresaUseCase := usecase_empresa.NewCreateEmpresaUseCase(repositoryEmpresa)
	listEmpresasUseCase := usecase_empresa.NewListEmpresasUseCase(repositoryEmpresa)
	getEmpresaUseCase := usecase_empresa.NewGetEmpresaUseCase(repositoryEmpresa)
	deleteEmpresaUseCase := usecase_empresa.NewDeleteEmpresaUseCase(repositoryEmpresa)
	updateEmpresaUseCase := usecase_empresa.NewUpdateEmpresaUseCase(repositoryEmpresa)

	return &EmpresaHandlers{
		CreateEmpresaUseCase: createEmpresaUseCase,
		ListEmpresasUseCase:  listEmpresasUseCase,
		DeleteEmpresaUseCase: deleteEmpresaUseCase,
		GetEmpresaUseCase:    getEmpresaUseCase,
		UpdateEmpresaUseCase: updateEmpresaUseCase}
}

func (p *EmpresaHandlers) UpdateEmpresaHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err := p.GetEmpresaUseCase.GetEmpresaByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	var input empresadto.EmpresaInputDTO
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	output, err := p.UpdateEmpresaUseCase.Execute(id, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *EmpresaHandlers) CreateEmpresaHandler(w http.ResponseWriter, r *http.Request) {
	var input empresadto.EmpresaInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	output, err := p.CreateEmpresaUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *EmpresaHandlers) ListEmpresasHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListEmpresasUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (p *EmpresaHandlers) DeleteEmpresaHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err := p.GetEmpresaUseCase.GetEmpresaByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = p.DeleteEmpresaUseCase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *EmpresaHandlers) GetEmpresaHandler(w http.ResponseWriter, r *http.Request) {

	var id string
	var cnpj string
	var chave string

	if chi.URLParam(r, "id") != "" {
		id = chi.URLParam(r, "id")
		//log.Printf("parametros ID:" + chi.URLParam(r, "id"))

		output, err := p.GetEmpresaUseCase.GetEmpresaByID(id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(output)
	} else if chi.URLParam(r, "cnpj") != "" {
		cnpj = chi.URLParam(r, "cnpj")
		//log.Printf("parametros CNPJ:" + chi.URLParam(r, "cnpj"))

		output, err := p.GetEmpresaUseCase.GetEmpresaByCNPJ(cnpj)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(output)
	} else if chi.URLParam(r, "chave") != "" {
		chave = chi.URLParam(r, "chave")
		//log.Printf("parametros Chave:" + chi.URLParam(r, "chave"))
		output, err := p.GetEmpresaUseCase.GetEmpresaByChaveRegistro(chave)
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
	}

}
