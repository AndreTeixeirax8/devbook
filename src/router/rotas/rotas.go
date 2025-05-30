package rotas

import "net/http"

type Rotas struct {
	URI                string `json:"uri"`
	Metodo             string `json:"metodo"`
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}
