package arreglos_slices

import "fmt"

//se declara la variable tabla que es un slice por no tener dimensiones "[]"
var tablaS []int = []int{2, 5, 4}

//creamos un arreglo como modelo para crear un slice que se ajuste a sus domensiones
//Y no solo eso que tambien me permite tomar algunos elemenos para crear el slice
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestroSlice() {
	//lo que hace el slice ajusto la dimencion en funcion a los valores que cuenta
	fmt.Println(tablaS)

	porcion := arreglo[3:]   //slice creado desde un vector, desde la posicion 3
	porcion2 := arreglo[:5]  //slice creado desde un vector, desde la posicion 0 hasta la 5
	porcion3 := arreglo[6:7] //slice creado desde un vector, desde la posicion 6 hasta la 7

	fmt.Println("-------------")
	fmt.Println(porcion)

	fmt.Println("-------------")
	fmt.Println(porcion2)

	fmt.Println("-------------")
	fmt.Println(porcion3)
}

func Capacidad() {

	//el "make" permite crear un slice vacio y darle una capacidad
	//con la definicion de 5 elementos y 20 de capacidad
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d\n", len(elementos), cap(elementos))

	fmt.Println("-------------")
	//crearemos un slice donde definimos que no tiene largo ni capacidad
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		//se adiciona al slice elementos
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d\n", len(nums), cap(nums))
}
