package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {

	fmt.Println("Ponto de partida")

	buscaApp := app.Gerar()
	if erro := buscaApp.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
	
	
}