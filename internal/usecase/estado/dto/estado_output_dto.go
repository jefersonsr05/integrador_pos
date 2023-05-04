package estadodto

type EstadoOutputDTO struct {
	ID        string `json:"id"`
	Descricao string `json:"descricao"`
	UF        string `json:"uf"`
	CodIbge   int32  `json:"codigo_ibge"`
}

type EstadosOutputDTO struct {
	Estados EstadoOutputDTO `json:"estados"`
}
