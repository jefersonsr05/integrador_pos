package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	cidadedto "github.com/jefersonsr05/integrador_pos/internal/usecase/cidade/dto"
)

type CreateCidadeUseCase struct {
	CidadeRepository entity.CidadeRepository
}

func NewCreateCidadeUseCase(cidadeRepository entity.CidadeRepository) *CreateCidadeUseCase {
	return &CreateCidadeUseCase{CidadeRepository: cidadeRepository}
}

func (u *CreateCidadeUseCase) Execute(input cidadedto.CidadeInputDTO) (*cidadedto.CidadeOutputDTO, error) {
	cidade := entity.NewCidade(input.Descricao, input.EstadoID, input.CodIbge)
	// log.Printf("Executando UseCase")
	// log.Println(input.Descricao)
	// log.Println(input.EstadoID)
	// log.Println(input.CodIbge)
	err := u.CidadeRepository.Create(cidade)
	if err != nil {
		// log.Printf("Erro Create UseCase" + err.Error())
		return nil, err
	}
	return &cidadedto.CidadeOutputDTO{
		ID:        cidade.ID,
		Descricao: cidade.Descricao,
		EstadoID:  cidade.EstadoID,
		CodIbge:   cidade.CodIbge,
	}, nil
}
