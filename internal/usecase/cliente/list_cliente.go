package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	clientedto "github.com/jefersonsr05/integrador_pos/internal/usecase/cliente/dto"
)

type ListClienteUseCase struct {
	ClienteRepository entity.ClienteRepository
}

func NewListClienteUseCase(clienteRepository entity.ClienteRepository) *ListClienteUseCase {
	return &ListClienteUseCase{ClienteRepository: clienteRepository}

}

func (u *ListClienteUseCase) Execute() ([]*clientedto.ClienteOutputDTO, error) {
	clientes, err := u.ClienteRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var clientesOutput []*clientedto.ClienteOutputDTO
	for _, cliente := range clientes {
		clientesOutput = append(clientesOutput, &clientedto.ClienteOutputDTO{
			ID:          cliente.ID,
			EmpresaID:   cliente.EmpresaID,
			CodigoMc:    cliente.CodigoMc,
			Nome:        cliente.Nome,
			Cep:         cliente.Cep,
			CidadeID:    cliente.CidadeID,
			Endereco:    cliente.Endereco,
			Numero:      cliente.Numero,
			Complemento: cliente.Complemento,
		})
	}
	return clientesOutput, nil
}
