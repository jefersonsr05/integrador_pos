package web

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	"github.com/jefersonsr05/integrador_pos/internal/infra/repository"
	usecase_cidade "github.com/jefersonsr05/integrador_pos/internal/usecase/cidade"
	cidadedto "github.com/jefersonsr05/integrador_pos/internal/usecase/cidade/dto"
)

type CidadeHandlers struct {
	CreateCidadeUseCase *usecase_cidade.CreateCidadeUseCase
	ListCidadeUseCase   *usecase_cidade.ListCidadeUseCase
	DeleteCidadeUseCase *usecase_cidade.DeleteCidadeUseCase
	GetCidadeUseCase    *usecase_cidade.GetCidadeUseCase
	UpdateCidadeUseCase *usecase_cidade.UpdateCidadeUseCase
}

func NewCidadeHandlers(
// createCidadeUseCase *usecase.CreateCidadeUseCase,
// listCidadeUseCase *usecase.ListCidadeUseCase,
// deleteCidadeUseCase *usecase.DeleteCidadeUseCase,
// getCidadeUseCase *usecase.GetCidadeUseCase,
// updateCidadeUseCase *usecase.UpdateCidadeUseCase
) *CidadeHandlers {
	dbConn, _ := db.Conectar()
	// defer dbConn.Close()

	repositoryCidade := repository.NewCidadeRepositoryMysql(dbConn)
	createCidadeUseCase := usecase_cidade.NewCreateCidadeUseCase(repositoryCidade)
	listCidadeUseCase := usecase_cidade.NewListCidadeUseCase(repositoryCidade)
	getCidadeUseCase := usecase_cidade.NewGetCidadeUseCase(repositoryCidade)
	deleteCidadeUseCase := usecase_cidade.NewDeleteCidadeUseCase(repositoryCidade)
	updateCidadeUseCase := usecase_cidade.NewUpdateCidadeUseCase(repositoryCidade)

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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
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
	var cidadesInput cidadedto.CidadesInputDTO
	err := json.NewDecoder(r.Body).Decode(&cidadesInput)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	// Itera sobre a lista de cidades e executa a operação de criação em cada uma
	var outputs []*cidadedto.CidadeOutputDTO
	for _, input := range cidadesInput.Cidades {
		log.Printf("Cadastrando empresa: " + input.Descricao)
		output, err := p.CreateCidadeUseCase.Execute(input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		outputs = append(outputs, output)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(outputs)
}

// func (p *CidadeHandlers) CreateCidadeHandler(w http.ResponseWriter, r *http.Request) {
// 	var input cidadedto.CidadeInputDTO
// 	// log.Printf("Entrando no cadastro")
// 	err := json.NewDecoder(r.Body).Decode(&input)
// 	if err != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(err)
// 	}
// 	// log.Printf("Decode JSON SUCCESS")
// 	output, err := p.CreateCidadeUseCase.Execute(input)
// 	if err != nil {
// 		// log.Printf("Erro Execute Cadastro JSON")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(output)
// }

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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("opção inválida")
	}

}
