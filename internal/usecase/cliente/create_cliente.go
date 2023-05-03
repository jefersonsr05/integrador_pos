package usecase

import (
	"log"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	clientedto "github.com/jefersonsr05/integrador_pos/internal/usecase/cliente/dto"
)

type CreateClienteUseCase struct {
	ClienteRepository entity.ClienteRepository
}

func NewCreateClienteUseCase(clienteRepository entity.ClienteRepository) *CreateClienteUseCase {
	return &CreateClienteUseCase{ClienteRepository: clienteRepository}
}

func (u *CreateClienteUseCase) Execute(input clientedto.ClienteInputDTO) (*clientedto.ClienteOutputDTO, error) {
	cliente := entity.NewCliente(input.EmpresaID, input.CodigoMc, input.Nome, input.Cep, input.CidadeID, input.Endereco, input.Numero, input.Complemento)
	err := u.ClienteRepository.Create(cliente)
	if err != nil {
		log.Printf("Erro Create UseCase" + err.Error())
		return nil, err
	}
	return &clientedto.ClienteOutputDTO{
		ID:          cliente.ID,
		EmpresaID:   cliente.EmpresaID,
		CodigoMc:    cliente.CodigoMc,
		Nome:        cliente.Nome,
		Cep:         cliente.Cep,
		CidadeID:    cliente.CidadeID,
		Endereco:    cliente.Endereco,
		Numero:      cliente.Numero,
		Complemento: cliente.Complemento,
	}, nil
}
