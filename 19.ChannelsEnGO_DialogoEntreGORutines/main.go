package main

import (
	"fmt"

	"github.com/xfred19x/godesde0/19.ChannelsEnGO_DialogoEntreGORutines/goroutines"
)

func main() {

	fmt.Println("-------------")
	//como declarar una variable channel de tipo booleano
	//el uso del channel es para esperar hasta que termine la go rutine
	canal1 := make(chan bool)

	//puedes guardar el estado booleano en una variable
	//estado <-canal1

	//en caso no quieras guardar en ninguna variable
	//<-canal1

	go goroutines.MiNombreLento("Freddy Gonzales", canal1)
	//en caso no quieras tener "<-canal1" en cualquier lugar lo puedes guardar en un "defer"
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy aqui")

}
