package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	pagamentodto "github.com/jefersonsr05/integrador_pos/internal/usecase/pagamento/dto"
)

type ListPagamentoUseCase struct {
	PagamentoRepository entity.PagamentoRepository
}

func NewListPagamentoUseCase(pagamentoRepository entity.PagamentoRepository) *ListPagamentoUseCase {
	return &ListPagamentoUseCase{PagamentoRepository: pagamentoRepository}

}

func (u *ListPagamentoUseCase) Execute() ([]*pagamentodto.PagamentoOutputDTO, error) {
	pagamentos, err := u.PagamentoRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var pagamentosOutput []*pagamentodto.PagamentoOutputDTO
	for _, pagamento := range pagamentos {
		pagamentosOutput = append(pagamentosOutput, &pagamentodto.PagamentoOutputDTO{
			ID:           pagamento.ID,
			EmpresaID:    pagamento.EmpresaID,
			Descricao:    pagamento.Descricao,
			TpPagamento:  pagamento.TpPagamento,
			IndPagamento: pagamento.IndPagamento,
			PosExclusivo: pagamento.PosExclusivo,
			IDPos:        pagamento.IDPos,
		})
	}
	return pagamentosOutput, nil
}
