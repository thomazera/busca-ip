package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {

	fmt.Println("Ponto de partida")

	//Caso o app.go gere algum erro o codigo abaixo irá apresentar o que ocasionou o erro.
	buscaApp := app.Gerar()
	if erro := buscaApp.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
	
	
}
