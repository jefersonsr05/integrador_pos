package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jefersonsr05/integrador_pos/internal/infra/db"
	api_routes "github.com/jefersonsr05/integrador_pos/internal/routes"
)

func main() {

	dbConn, _ := db.Conectar()
	defer dbConn.Close()

	log.Printf("iniciando serviços")

	r := chi.NewRouter()

	r.Route("/empresas", api_routes.RouterEmpresa)
	r.Route("/clientes", api_routes.RouterCliente)
	r.Route("/estados", api_routes.RouterEstado)
	r.Route("/cidades", api_routes.RouterCidade)
	r.Route("/itens_venda", api_routes.RouterItemVenda)
	r.Route("/maquininhas_pos", api_routes.RouterMaquininhaPOS)
	r.Route("/pagamentos", api_routes.RouterPagamento)
	r.Route("/pagamentos_venda", api_routes.RouterPagamentoVenda)
	r.Route("/produtos", api_routes.RouterProduto)
	r.Route("/vendas", api_routes.RouterVenda)

	log.Printf("serviços inicializados")
	http.ListenAndServe(":8000", r)

}
