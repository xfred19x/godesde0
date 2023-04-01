package variables

//aqui el paquete lleva el nombre de la carpeta

import "fmt"

//Si una funcion la quieres hacer publica y otros paquetes lo importen
//La primera Letra debe estar en mayusculas
func MostrarEnteros() { //en go la primera debe estar unicamente en 1 sola linea

	//2 formas de declarar una variable
	//1ra forma de declarar una variable por declaracion
	var intcomun int
	//Lo bueno de go es que al no inicializar la variable, Go lo hace con el valor mas minimo
	//Por ejemplo si es "int" y no inicializa no lo hace con nulo sino 0

	//2da forma de declarar una variable por asignacion
	intde32 := int32(10)
	intde64 := int64(100)

	//intde64 = intcomun
	//Go es tipeado, cuando intentas igualar variables de diferente tipo Go no lo deja.

	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)

} //en go la llave fin debe estar unicamente en 1 sola linea
