package clientedto

type ClienteOutputDTO struct {
	ID          string `json:"id"`
	EmpresaID   string `json:"empresa_id"`
	CodigoMc    int64  `json:"codigo_mc"`
	Nome        string `json:"nome"`
	Cep         string `json:"cep"`
	CidadeID    string `json:"cidade_id"`
	Endereco    string `json:"endereco"`
	Numero      string `json:"numero"`
	Complemento string `json:"complemento"`
}
