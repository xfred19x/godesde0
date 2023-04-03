package teclado

import (
	"bufio" //libreria que tiene que ver con lectura de teclado, archivo, etc.
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error //permite manejar los posibles errores de datos

func IngresoNumeros() {

	//inicializa una variable que leera los datos del estandar input (Stind)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese néumero 1 : ")

	//este if asegura que solo se ejecutara los pasos si ingresaron algun dato.
	if scanner.Scan() {
		//convertimos un string a un entero con "strconv.Atoi", ya que el input siempre es un string
		numero1, err = strconv.Atoi(scanner.Text())

		//se valida en caso de algun error
		if err != nil {
			//si quieren abortar la aplicacion, lo hacen con "panic"
			panic("El dato ingresado es incorrecto " + err.Error())
		}

	}

	fmt.Println("Ingrese néumero 2 : ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}

	}

	fmt.Println("Ingrese leyenda : ")
	if scanner.Scan() {
		leyenda = scanner.Text()

	}

	fmt.Println(leyenda, numero1*numero2)
}
