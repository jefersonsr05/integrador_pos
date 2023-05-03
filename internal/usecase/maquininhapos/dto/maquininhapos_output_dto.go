package maquininhaposdto

type MaquininhaPosOutputDTO struct {
	ID             string `json:"id"`
	EmpresaID      string `json:"empresa_id"`
	Descricao      string `json:"descricao"`
	Administradora string `json:"administradora"`
	Cnpj           string `json:"cnpj"`
}
