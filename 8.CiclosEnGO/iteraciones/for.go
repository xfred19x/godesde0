package iteraciones

import (
	"fmt"
)

func Iterar() {

	//1ra declaracion de For hasta un limite, saltando 1 en 1
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("-------------")
	//2da declaracion de For hasta un limite, saltando de 5 en 5
	for j := 0; j < 100; j += 5 {
		fmt.Println(j)
	}

	fmt.Println("-------------")
	//3ra declaracion de For hasta un limite, reste de 5 en 5
	for k := 100; k > 10; k -= 5 {
		fmt.Println(k)
	}

	fmt.Println("-------------")
	//4ta declaracion de For hasta un limite, reste de 1 en 1
	for l := 10; l > 1; l-- {
		fmt.Println(l)
	}

	fmt.Println("-------------")
	//Podemos usar el break para abortar el for
	for m := 10; m > 1; m-- {
		if m == 6 {
			break
		}
		fmt.Println(m)
	}

	fmt.Println("-------------")
	//Podemos usar el continue para que siga los pasos siguiens del for y continue con la siguiente iteracion
	for m := 10; m > 1; m-- {
		if m == 6 {
			continue
		}
		fmt.Println(m)
	}
}
