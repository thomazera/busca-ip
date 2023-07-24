package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)




//Gerando um busca ip
func Gerar() *cli.App {

app := cli.NewApp()
app.Usage = "Busca ips na internet"
app.Name = "Aplicação de linha de comando"

flags := []cli.Flag{
	
	cli.StringFlag{
		Name: "host", 
		Value: "google.com.br",
	},
}
	 

app.Commands = []cli.Command {
		{
			Name: "ip",
			Usage: "Busca Ips de endereços na internet",
			Flags: flags,
			Action: buscarIp,
		}, 
		{
			Name: "servidor",
			Usage: "Busca servidores na internet",
			Flags: flags,
			Action: buscarServidor,
		},
	}
		
	return app

}

// A usando a funçao buscar ip
func buscarIp(c *cli.Context) {
	
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

} 



func buscarServidor (c *cli.Context) {

	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

}