package usecase

import (
	"log"

	"github.com/jefersonsr05/integrador_pos/internal/entity"
	pagamentodto "github.com/jefersonsr05/integrador_pos/internal/usecase/pagamento/dto"
)

type UpdatePagamentoUseCase struct {
	PagamentoRepository entity.PagamentoRepository
}

func NewUpdatePagamentoUseCase(pagamentoRepository entity.PagamentoRepository) *UpdatePagamentoUseCase {
	return &UpdatePagamentoUseCase{PagamentoRepository: pagamentoRepository}
}

func (u *UpdatePagamentoUseCase) Execute(id string, input pagamentodto.PagamentoInputDTO) (*pagamentodto.PagamentoOutputDTO, error) {
	pagamento, err := u.PagamentoRepository.GetPagamentoByID(id)
	if err != nil {
		log.Printf("não achei resultado")
		return nil, err
	}

	pagamento.EmpresaID = input.EmpresaID
	pagamento.Descricao = input.Descricao
	pagamento.TpPagamento = input.TpPagamento
	pagamento.IndPagamento = input.IndPagamento
	pagamento.PosExclusivo = input.PosExclusivo
	pagamento.IDPos = input.IDPos

	err = u.PagamentoRepository.Update(pagamento)
	if err != nil {
		return nil, err
	}

	return &pagamentodto.PagamentoOutputDTO{
		ID:           pagamento.ID,
		EmpresaID:    pagamento.EmpresaID,
		Descricao:    pagamento.Descricao,
		TpPagamento:  pagamento.TpPagamento,
		IndPagamento: pagamento.IndPagamento,
		PosExclusivo: pagamento.PosExclusivo,
		IDPos:        pagamento.IDPos,
	}, nil
}
