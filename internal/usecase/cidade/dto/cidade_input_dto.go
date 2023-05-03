package cidadedto

type CidadeInputDTO struct {
	Descricao string `json:"descricao"`
	EstadoID  string `json:"estado_id"`
	CodIbge   string `json:"codigo_ibge"`
}
