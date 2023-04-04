package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TabladeMultiplicar() string {

	scanner := bufio.NewScanner(os.Stdin)

	var numero int
	var err error
	var texto string //es la variable que tendra toda la informacion que se va grabar en el archivo

	for {
		fmt.Println("Ingrese un número : ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i < 10; i++ {
		texto += fmt.Sprintf("%d X %d = %d \n", numero, i, i*numero) //este formateo con la funcion Sprintf retorna un string

	}

	return texto
}
