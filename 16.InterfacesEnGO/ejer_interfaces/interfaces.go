package ejer_interfaces

import (
	"fmt"

	"github.com/xfred19x/godesde0/16.InterfacesEnGO/interfaces"
)

// Creamos una funcion que recibe un parametro Humano, pero este puede ser Hombre o Mujer
// ya que todas esas interfaces comparten mismas funciones quiere decir que hombre y mujer son humanos
func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy una/a %s, y estoy respirando \n", hu.Sexo())
}
