package router

import (
	"api-devbook/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar retorna uma rota configurada
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
