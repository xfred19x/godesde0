package defer_panic

import (
	"fmt"
	"log"
)

// Instruccion de Go que nos indica cual va ser la ultima instruccion cuando salga de la funcion
// es bueno usarlo cuando se origina un error abrupto que no permita terminar ciertas instrucciones
// solo acepta 1 instruccion
func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")
	fmt.Println("Este es el segundo mensaje")
}

// Instruccion aborta el programa con un mensaje
// Lo hace cuando detecta alguna falla
func EjemploPanic() {
	defer func() {
		//se dice que el recover siempre se usa con el "defer"
		//en este punto si el recover detecta un panic o no se ejecuta
		//sino detecta un panic llena una variable "reco" con "nil"
		reco := recover()

		//validamos si reco es "nil"
		if reco != nil {
			//usaremos la libreria "log" para poder definir la fecha y horadel error
			//Usaremos la funcion "Fatalf" para salir de ejcucion
			log.Fatalf("Ocurrio un error que genero Panic \n %v", reco)
		}

	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
