package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//REVISAR TODO O FLUXO PARA CRIAR USUARIO APENAS REVISAR PARA ENTENDER

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- Inicia função de criar usuario na controller")
	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	repositorio.Criar((usuario))

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando todos usuário"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando um usuário"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("alterando um usuário"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deletar um usuário"))
}
