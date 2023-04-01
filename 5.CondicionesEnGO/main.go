package main

import (
	"fmt"
	"runtime" //paquete que tiene toda la informacion del equipo que esta corriendo nuestro sistema
)

func main() {

	fmt.Println("-------------")
	//1era declaracion de IF
	//se pone la condicion y tambien se declara una variable inicializada
	//operadores de comparacion igual "==", mayor ">", menor "<", mayorigual">=", menorigual "<=", and "&&" y or "||"
	if os := runtime.GOOS; os == "Linux." {
		fmt.Println("Esto no es Windows, es", os)

	} else {
		fmt.Println("Esto es Windows")

	}

	fmt.Println("-------------")
	//1ra declaracion de SWITCH
	//se pone la variable a evaluar y tambien se declara una variable inicializada
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		//Permite que formateemos el texto
		fmt.Printf("%s \n", os)
	}

}
