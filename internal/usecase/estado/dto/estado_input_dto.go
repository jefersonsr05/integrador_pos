package estadodto

type EstadoInputDTO struct {
	Descricao string `json:"descricao"`
	UF        string `json:"uf"`
	CodIbge   int32  `json:"codigo_ibge"`
}

type EstadosInputDTO struct {
	Estados EstadoInputDTO `json:"estados"`
}
