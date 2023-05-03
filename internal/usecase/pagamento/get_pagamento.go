package usecase

import (
	"github.com/jefersonsr05/integrador_pos/internal/entity"
	pagamentodto "github.com/jefersonsr05/integrador_pos/internal/usecase/pagamento/dto"
)

type GetPagamentoUseCase struct {
	PagamentoRepository entity.PagamentoRepository
}

func NewGetPagamentoUseCase(pagamentoRepository entity.PagamentoRepository) *GetPagamentoUseCase {
	return &GetPagamentoUseCase{PagamentoRepository: pagamentoRepository}

}

func (u *GetPagamentoUseCase) GetPagamentoByID(id string) (*pagamentodto.PagamentoOutputDTO, error) {
	pagamento, err := u.PagamentoRepository.GetPagamentoByID(id)
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

func (u *GetPagamentoUseCase) GetPagamentoByEmpresa(id string) ([]*pagamentodto.PagamentoOutputDTO, error) {
	pagamentos, err := u.PagamentoRepository.GetPagamentoByEmpresaID(id)
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
