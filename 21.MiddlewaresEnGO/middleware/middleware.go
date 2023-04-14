package middleware

import "fmt"

//creando y agrupando funciones del mismo tipo, 2 input y 1 output entero
func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

//Si tenemos funciones del mismo input/output podemos crear un middleware
func Middleware() {
	fmt.Println("Inicio de Midleware")
	fmt.Println("")

	//vamos a crear nuestra funcion middleware "operacionesMidd" que tendra la misma forma de las funciones operacionales
	//tendra como input la funcion y los valores a procesar
	result := operacionesMidd(sumar)(2, 3)
	fmt.Println(result)

	result = operacionesMidd(restar)(10, 6)
	fmt.Println(result)

	result = operacionesMidd(multiplicar)(2, 4)
	fmt.Println(result)
}

//creando funcion middleware
//definiremos que como input tendra un funcion "f", con 2 parametros de entrada y que retorna 1 entero
//como es middleware si reciben un tipo de dato o unos tipos de datos determinado, deve retornar el mismo
func operacionesMidd(f func(int, int) int) func(int, int) int {
	//este va retornar el return de una funcion
	return func(a, b int) int {
		//aqui agregar la logica que quieras
		fmt.Println("Inicio de Operacion")
		//aqui se va retornar la funcion del tipo que corresponde
		return f(a, b)
	}
}
