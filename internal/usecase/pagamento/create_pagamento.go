package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	pagamentodto "github.com/jefersonsr05/integrador_pos/internal/usecase/pagamento/dto"
)

type CreatePagamentoUseCase struct {
	PagamentoRepository entity.PagamentoRepository
}

func NewCreatePagamentoUseCase(pagamentoRepository entity.PagamentoRepository) *CreatePagamentoUseCase {
	return &CreatePagamentoUseCase{PagamentoRepository: pagamentoRepository}
}

func (u *CreatePagamentoUseCase) Execute(input pagamentodto.PagamentoInputDTO) (*pagamentodto.PagamentoOutputDTO, error) {
	pagamento := entity.NewPagamento(input.EmpresaID, input.Descricao, input.TpPagamento, input.IndPagamento, input.PosExclusivo, input.IDPos)
	err := u.PagamentoRepository.Create(pagamento)
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
