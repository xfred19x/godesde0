package variables

import (
	"fmt"
	"time" //importa pacquete tipo time para poder obtener funciones y propiedades de tiempo
)

// Se declara una variable de forma publica por tener la primera letra en mayuscula
// Al no inicializar esta se crea no con nulo sino vacio
var Nombre string

var Estado bool
var Sueldo float32

// Con esta variable podemos obtener la hora actual, adicionar fechas, etc.
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()

	//Porque el Println te permite imprimir un booleano y no un Texto?
	//es que esta funcion recibe variable del tipo Any = Cualquier dato
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}
