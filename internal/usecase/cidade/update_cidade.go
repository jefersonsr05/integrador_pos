package usecase

import (
	"log"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	cidadedto "github.com/jefersonsr05/integrador_pos/internal/usecase/cidade/dto"
)

type UpdateCidadeUseCase struct {
	CidadeRepository entity.CidadeRepository
}

func NewUpdateCidadeUseCase(cidadeRepository entity.CidadeRepository) *UpdateCidadeUseCase {
	return &UpdateCidadeUseCase{CidadeRepository: cidadeRepository}
}

func (u *UpdateCidadeUseCase) Execute(id string, input cidadedto.CidadeInputDTO) (*cidadedto.CidadeOutputDTO, error) {
	cidade, err := u.CidadeRepository.GetCidadeByID(id)
	if err != nil {
		log.Printf("não achei resultado")
		return nil, err
	}

	cidade.Descricao = input.Descricao
	cidade.EstadoID = input.EstadoID
	cidade.CodIbge = input.CodIbge

	err = u.CidadeRepository.Update(cidade)
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
