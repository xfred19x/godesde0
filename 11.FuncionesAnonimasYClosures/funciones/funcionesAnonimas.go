package funciones

import "fmt"

func Calculos() {

	var numero3, numero4 int = 32, 243
	//declarando funcion anonima
	//estamos asignando una funcion a una variable con 2 parametros de entrada y un tipo de parametro de salida
	suma := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}

	fmt.Println(suma(10, 25))

	//declarar otras variables con funciones anonimas es como sobrecargas de las funciones pero anonimas
	//es como reinventar el calculo pero desde una misma funcion
	fmt.Println("-------------")
	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}
	fmt.Println(calculo(10, 25))

	fmt.Println("-------------")
	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3
	}
	fmt.Println(calculo(10, 25))
}
