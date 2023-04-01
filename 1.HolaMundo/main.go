package main

//el paquete lleva el nombre de main, solo si esta en la carpeta principal
//solo puede existir unocls

import (
	"fmt"
)

//en el import puedes invocar una sola libreria entre ""
//en caso de ser varias entre parentesis y dentro de ellas cada paquete en comillas ""

func main() {
	//Sirve para imprimir por consola
	fmt.Println("Hola mundo")
}

//se ejecuta con el comando "go run main.go"
//esto te ejecuta por el terminal

//se compila con el comando "go build main.go"
//genera un ejecutable que ejecutado desde CMD se ejecuta sin problema
