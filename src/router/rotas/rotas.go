package rotas

import "net/http"

// Rota representa todas as rotas da api
type Rotas struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}
