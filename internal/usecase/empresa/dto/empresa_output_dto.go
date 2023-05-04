package empresadto

type EmpresaOutputDTO struct {
	ID            string `json:"id"`
	Descricao     string `json:"descricao"`
	Cnpj          string `json:"cnpj"`
	ChaveRegistro string `json:"chave_registro"`
}

type EmpresasOutputDTO struct {
	Empresas EmpresaOutputDTO `json:"empresas"`
}
