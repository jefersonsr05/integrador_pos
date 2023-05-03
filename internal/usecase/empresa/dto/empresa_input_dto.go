package empresadto

type EmpresaInputDTO struct {
	Descricao     string `json:"descricao"`
	Cnpj          string `json:"cnpj"`
	ChaveRegistro string `json:"chave_registro"`
}
