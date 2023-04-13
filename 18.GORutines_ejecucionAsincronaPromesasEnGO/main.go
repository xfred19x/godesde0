package main

import (
	"fmt"

	"github.com/xfred19x/godesde0/18.GORutines_ejecucionAsincronaPromesasEnGO/goroutines"
)

func main() {

	fmt.Println("-------------")
	goroutines.MiNombreLento("Freddy Gonzales")

	fmt.Println("-------------")
	//si quieres volver asincrono debes agregar la palabra "go"
	go goroutines.MiNombreLento("Freddy Gonzales")
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)
	//cuando ejecutes el go rutines espera un poco e ingresa cualquier letra y veras que pausa la ejecucion
}
