package files

import (
	"bufio"
	"fmt"

	//"io/ioutil" //libreria deprecada pero si funciona, descomentar LeoArchivos para probar
	"os"

	"github.com/xfred19x/godesde0/10.ManejoDeArchivosEnGO/ejercicios"
)

// inicializa la variable con el nombre del path y archivo que ha de crearse
var filename string = "./files/txt/tabla.txt"

// Funcion que graba la tabla en un archivo
func GrabaTabla() {

	var texto string = ejercicios.TabladeMultiplicar()

	//crea el archivo
	archivo, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return
	}

	//La funcion Fprintln graba en el archivo e texto
	fmt.Fprintln(archivo, texto)
	//Todo arhicvo que se habre para grabar tiene que cerrarse
	archivo.Close()

	//Ojo cada vez que se ejecute esta funcion va sobreescribir el archivo

}

// Va agarrar cualquier archivo que le infiquemos y agregara la tabla al final
func SumaTabla() {
	var texto string = ejercicios.TabladeMultiplicar()

	//se va crear una funcion custom que recibira 2 parametros de entrada y output booleano
	//luego validara si el archivo fue correctamente creado

	//Primera forma de validar el resultado
	/*if Append(filename, texto) == false {
		fmt.Println("Error al concatenar contenido")
	}*/

	//Segunda forma de validar el resultado
	if !Append(filename, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {

	//La funcion OpenFIle abre un archivo de lectura
	//El flag filename, os.O_WRONLY indica que lo abra de modo escritura
	//El flag filename, os.O_APPEND indica que lo abra de forma append
	//Por ultimo se le da los permisos de lectura y escritura "0644"
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}

	//Funcion que permite grabar texto en el archivo
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}

	arch.Close()
	return true
}

//Primera forma de leer archivos
/*func LeoArchivos() {

	//Funcion ReadFile para leer arhivos
	archivo, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}

	fmt.Println(string(archivo))
}*/

// Segunda forma de leer archivos
func LeoArchivos2() {

	//funcion Open para leer archivo
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Hubo un error leyendo el archivo " + err.Error())
		return
	}

	//usaremos el NewScanner para obtener los datos del archivo
	scanner := bufio.NewScanner(archivo)

	//Ejecuta el for mientras hace un scanner linea por linea mientras tenga registros
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}

	archivo.Close()
}
