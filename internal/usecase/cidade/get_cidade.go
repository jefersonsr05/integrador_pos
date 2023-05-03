package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	cidadedto "github.com/jefersonsr05/integrador_pos/internal/usecase/cidade/dto"
)

type GetCidadeUseCase struct {
	CidadeRepository entity.CidadeRepository
}

func NewGetCidadeUseCase(cidadeRepository entity.CidadeRepository) *GetCidadeUseCase {
	return &GetCidadeUseCase{CidadeRepository: cidadeRepository}

}

func (u *GetCidadeUseCase) GetCidadeByID(id string) (*cidadedto.CidadeOutputDTO, error) {
	cidade, err := u.CidadeRepository.GetCidadeByID(id)
	if err != nil {
		return nil, err
	}
	return &cidadedto.CidadeOutputDTO{
		ID:        cidade.ID,
		Descricao: cidade.Descricao,
		EstadoID:  cidade.EstadoID,
		CodIbge:   cidade.CodIbge,
	}, nil
}

func (u *GetCidadeUseCase) GetCidadeByIBGE(ibge string) (*cidadedto.CidadeOutputDTO, error) {
	cidade, err := u.CidadeRepository.GetCidadeByIBGE(ibge)
	if err != nil {
		return nil, err
	}
	return &cidadedto.CidadeOutputDTO{
		ID:        cidade.ID,
		Descricao: cidade.Descricao,
		EstadoID:  cidade.EstadoID,
		CodIbge:   cidade.CodIbge,
	}, nil
}

func (u *GetCidadeUseCase) GetCidadeByUF(uf string) ([]*cidadedto.CidadeOutputDTO, error) {
	cidades, err := u.CidadeRepository.GetCidadeByUF(uf)
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
	return cidadesOutput, nil
}

func (u *GetCidadeUseCase) GetCidadeByIbgeUF(ibgeuf int32) ([]*cidadedto.CidadeOutputDTO, error) {
	cidades, err := u.CidadeRepository.GetCidadeByIbgeUF(ibgeuf)
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
	return cidadesOutput, nil
}
