package funciones

import "fmt"

//Esta funcion tabla va devolver una funcion anonima que retorna un valor entero
func tabla(valor int) func() int {

	numero := valor
	secuencia := 0

	//aqui se ve que el return es una funcion que retorna un entero
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClosure() {

	tabladel := 2

	//creamos una variable que inicializara la secuencia pero con la tabla del 2
	MiTabla := tabla(tabladel)

	//este ciclo va repetir 10 veces la funcion Tabla
	//si te das cuenta la secuencia va ir guardando su posicion en memoria
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}

}
