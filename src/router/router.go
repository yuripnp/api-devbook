package router

import "github.com/gorilla/mux"

// Gerar retorna uma rota configurada
func Gerar() *mux.Router {
	return mux.NewRouter()
}
