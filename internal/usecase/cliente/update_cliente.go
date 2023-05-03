package usecase

import (
	"log"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	clientedto "github.com/jefersonsr05/integrador_pos/internal/usecase/cliente/dto"
)

type UpdateClienteUseCase struct {
	ClienteRepository entity.ClienteRepository
}

func NewUpdateClienteUseCase(clienteRepository entity.ClienteRepository) *UpdateClienteUseCase {
	return &UpdateClienteUseCase{ClienteRepository: clienteRepository}
}

func (u *UpdateClienteUseCase) Execute(id string, input clientedto.ClienteInputDTO) (*clientedto.ClienteOutputDTO, error) {
	cliente, err := u.ClienteRepository.GetClienteByID(id)
	if err != nil {
		log.Printf("não achei resultado")
		return nil, err
	}

	cliente.EmpresaID = input.EmpresaID
	cliente.CodigoMc = input.CodigoMc
	cliente.Nome = input.Nome
	cliente.Cep = input.Cep
	cliente.CidadeID = input.CidadeID
	cliente.Endereco = input.Endereco
	cliente.Numero = input.Numero
	cliente.Complemento = input.Complemento

	err = u.ClienteRepository.Update(cliente)
	if err != nil {
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
