package main

import (
	"fmt"

	"github.com/xfred19x/godesde0/16.InterfacesEnGO/ejer_interfaces"
	"github.com/xfred19x/godesde0/16.InterfacesEnGO/modelos"
)

// interfaces son modelos de comportamientos que podemos agregar a nuestras estructuras
func main() {

	fmt.Println("-------------")
	Pedro := new(modelos.Hombre)
	//como ven se envia una estructura de tipo hombre, es porque tanto humano y hombre comparten las mismas funciones
	ejer_interfaces.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	ejer_interfaces.HumanosRespirando(Maria)

}
