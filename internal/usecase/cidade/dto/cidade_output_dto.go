package cidadedto

type CidadeOutputDTO struct {
	ID        string `json:"id"`
	Descricao string `json:"descricao"`
	EstadoID  string `json:"estado_id"`
	CodIbge   string `json:"codigo_ibge"`
}
