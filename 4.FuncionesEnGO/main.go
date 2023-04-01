package main

import (
	"fmt"

	"github.com/xfred19x/godesde0/4.FuncionesEnGO/variables"
)

func main() {

	fmt.Println("-------------")
	//Al llamar la funcion debo declarar variables para que reciba la respuesta
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

}

//Para no estar escribiendo a cada momento uno por uno el comando git
//se creara un archivo subo.bat
//luego ejecutamos el comando ./subo.bat
