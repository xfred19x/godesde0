package arreglos_slices

import "fmt"

//un slice es un tipo de arreglo
//declaramos una variable que se convierte en una coleccion de un tipo de dato
//cuando no tiene una dimension "[]" es una slice, caso contrario "[10]" seria un arreglo
var tabla [10]int

//Declaramos una matriz, que es un arreglo vector de vectores
var matriz [20][30]int

//otro tipo de declaracion e inicializacion de un vector
//podemos definir los valores y los que no, lo rellena con "0"
var tabla2 [10]int = [10]int{10, 0, 59}

func MuestroArreglos() {

	//grabamos valores fijos en la posicion 7 y 2
	tabla[7] = 33
	tabla[2] = 54

	fmt.Println(tabla)

	fmt.Println("-------------")
	tabla2[7] = 33
	tabla2[2] = 54
	fmt.Println(tabla2)

	fmt.Println("-------------")
	tabla3 := [10]string{"Pablo", "Juan"}
	fmt.Println(tabla3)

	fmt.Println("-------------")
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	fmt.Println("-------------")
	matriz[7][24] = 15
	fmt.Println(matriz)

}
