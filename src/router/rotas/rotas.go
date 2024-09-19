package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da api
type Rotas struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar, coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotasUsuario := rotasUsuarios

	for _, rota := range rotasUsuario {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
