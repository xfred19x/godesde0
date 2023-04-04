package funciones

import "fmt"

//La recurcion es una funcion que se llama a si misma n veces
func Exponencia(valor int) {
	//va parar cuando llegue a los 1000000000
	if valor > 1000000000 {
		return
	}

	fmt.Println(valor)
	//aqui se vuelve llamar n veces con un input "valor*2"
	Exponencia(valor * 2)
}
