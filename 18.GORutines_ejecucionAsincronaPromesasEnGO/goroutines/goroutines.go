package goroutines

import (
	"fmt"
	"strings" //libreria que tiene todo que ver con strings
	"time"
)

// vamos a mostrar un nombre de manera lenta en pantalla, caracter por caracter
func MiNombreLento(nombre string) {
	//crearemos un vector y usaremos una cadena que se generara a partir de la funcion "split"
	letras := strings.Split(nombre, "")

	//vamos a recorrer letra por letra
	for _, letra := range letras {
		//vamos hacer una pausa en el sistema de 1 segundo = 1000* time.Millisecond
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
