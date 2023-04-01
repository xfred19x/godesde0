package variables

import (
	"strconv" //importa paquete que nos permite agregar funcionalidad
)

// 2da forma declarar una funcion, con parametros de entrada y salida
// En este ejemplo el input es un parametro entero
// Y el output son 2 tipo de variable bool y string
func ConviertoaTexto(numero int) (bool, string) {

	var texto string

	//texto = numero -> No deja setear directamente 2 tipos de datos distintos

	//si usamos para convertir a texto string(numero)
	//este se encuentra limitado a solo a 8 bytes
	//si un numero supera el resultado no es el esperado
	//Para probar descomente la linea de abajo
	//texto = string(numero)

	//Se llamara la funcion Itoa que convierte el entero a string
	texto = strconv.Itoa(numero)

	return true, texto

}
