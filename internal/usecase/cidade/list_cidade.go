package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	cidadedto "github.com/jefersonsr05/integrador_pos/internal/usecase/cidade/dto"
)

type ListCidadeUseCase struct {
	CidadeRepository entity.CidadeRepository
}

func NewListCidadeUseCase(cidadeRepository entity.CidadeRepository) *ListCidadeUseCase {
	return &ListCidadeUseCase{CidadeRepository: cidadeRepository}

}

// func (u *ListCidadeUseCase) Execute() ([]*cidadedto.CidadeOutputDTO, error) {
// 	cidades, err := u.CidadeRepository.FindAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var cidadesOutput []*cidadedto.CidadeOutputDTO
// 	for _, cidade := range cidades {
// 		cidadesOutput = append(cidadesOutput, &cidadedto.CidadeOutputDTO{
// 			ID:        cidade.ID,
// 			Descricao: cidade.Descricao,
// 			EstadoID:  cidade.EstadoID,
// 			CodIbge:   cidade.CodIbge,
// 		})
// 	}
// 	return cidadesOutput, nil
// }

func (u *ListCidadeUseCase) Execute() ([]*cidadedto.CidadesOutputDTO, error) {
	cidades, err := u.CidadeRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var cidadesOutput []*cidadedto.CidadeOutputDTO
	for _, cidade := range cidades {
		cidadesOutput = append(cidadesOutput, &cidadedto.CidadeOutputDTO{
			ID:        cidade.ID,
			Descricao: cidade.Descricao,
			EstadoID:  cidade.EstadoID,
			CodIbge:   cidade.CodIbge,
		})
	}

	output := cidadedto.CidadesOutputDTO{
		Cidades: cidadesOutput,
	}

	return []*cidadedto.CidadesOutputDTO{&output}, nil
}
