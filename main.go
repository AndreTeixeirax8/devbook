package main

import (
	router "api/src"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando Api")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":3000", r))
}
