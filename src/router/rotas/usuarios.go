package rotas

import "net/http"

var rotasUsuarios = []Rotas{
	{
		URI:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios",
		Metodo: http.MethodGet,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodPut,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	},
}
