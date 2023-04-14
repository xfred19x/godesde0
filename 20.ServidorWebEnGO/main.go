package main

import (
	"fmt"

	"github.com/xfred19x/godesde0/20.ServidorWebEnGO/webserver"
)

func main() {

	fmt.Println("-------------")
	webserver.MiWebServer()

	//cuando ejecutes este main.go quedara pegado y escuchando
	//debemos ir a "http://localhost:3000/" para ver el contenido de index.hmtl

}
