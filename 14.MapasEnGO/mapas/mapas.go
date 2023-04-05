package mapas

import "fmt"

//Un mapa tiene un parecido a los vectores
//pero un mapa es una coleccion de clave/valor
//no tiene que ver con la cantidad de elementos ni posicion
func MostrarMapas() {

	//1ra forma de crear un map con clave y valor de tipo string
	paises := make(map[string]string)
	//aqui el map se imprime en vacío
	fmt.Println(paises)

	fmt.Println("-------------")
	paises2 := make(map[string]string)
	//se agrega elementos clave/valor al map
	paises2["Mexico"] = "D.F"
	paises2["Argentina"] = "Buenos Aires"
	//aqui el map se mostrara completo
	fmt.Println(paises2)

	fmt.Println("-------------")
	//Aqui escogeremos un map en especifico con la llave y mostrará solo el valor
	fmt.Println(paises2["Argentina"])

	//2da forma de crear un map con clave y valor de tipo string, ademas de un largo determinado de "5"
	//paises := make(map[string]string, 5)

	fmt.Println("-------------")
	//3ra forma de crear un map con clave y valor de tipo string
	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}
	//cuando lo imprime lo muestra de forma desordenado
	fmt.Println(campeonato)

	fmt.Println("-------------")
	//vamos a recorrer con for/range toda la coleccion del map
	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	fmt.Println("-------------")
	//para eliminar un elemento del mapa se usa la funcion delete
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	fmt.Println("-------------")
	//Ejemplo cuando no existe un elemento del mapa
	puntaje, existe := campeonato["Juventus"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)

	fmt.Println("-------------")
	//Ejemplo cuando existe un elemento del mapa
	puntaje, existe = campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
}
