package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	clientedto "github.com/jefersonsr05/integrador_pos/internal/usecase/cliente/dto"
)

type GetClienteUseCase struct {
	ClienteRepository entity.ClienteRepository
}

func NewGetClienteUseCase(clienteRepository entity.ClienteRepository) *GetClienteUseCase {
	return &GetClienteUseCase{ClienteRepository: clienteRepository}

}

func (u *GetClienteUseCase) GetClienteByID(id string) (*clientedto.ClienteOutputDTO, error) {
	cliente, err := u.ClienteRepository.GetClienteByID(id)
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

// func (u *GetClienteUseCase) GetClienteByCodigoMC(codigomc string) (*clientedto.ClienteOutputDTO, error) {
// 	cliente, err := u.ClienteRepository.GetClienteByCodigoMC(codigomc)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &clientedto.ClienteOutputDTO{
// 		ID:          cliente.ID,
// 		EmpresaID:   cliente.EmpresaID,
// 		CodigoMc:    cliente.CodigoMc,
// 		Nome:        cliente.Nome,
// 		Cep:         cliente.Cep,
// 		CidadeID:    cliente.CidadeID,
// 		Endereco:    cliente.Endereco,
// 		Numero:      cliente.Numero,
// 		Complemento: cliente.Complemento,
// 	}, nil
// }

func (u *GetClienteUseCase) GetClienteByEmpresa(empresa string) ([]*clientedto.ClienteOutputDTO, error) {
	clientes, err := u.ClienteRepository.GetClienteByEmpresa(empresa)
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
