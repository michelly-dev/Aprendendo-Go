package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar vai retornar a aplicação de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes do Servidor sa internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "gogoole.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPS de endereços da internet",
			Flags:  flags,
			Action: bucarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidor,
		},
	}

	return app
}

func bucarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidor(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) //name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
