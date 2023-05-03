package maquininhaposdto

type MaquininhaPosInputDTO struct {
	EmpresaID      string `json:"empresa_id"`
	Descricao      string `json:"descricao"`
	Administradora string `json:"administradora"`
	Cnpj           string `json:"cnpj"`
}
