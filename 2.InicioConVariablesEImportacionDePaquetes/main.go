package main

import (
	"fmt"

	"github.com/xfred19x/godesde0/2.InicioConVariablesEImportacionDePaquetes/variables" //aqui importamos el paquete variables
)

func main() {

	fmt.Println("-------------")
	//Aqui invocamos la funcion Publica del paquete "variables"
	variables.MostrarEnteros()
	//Lo bueno de esto es que se puede distribuir codigo reutilizable en carpetas
	//ante nuevos desarrollos se puede importar dichos paquetes

}

//Comandos GITHUB
//"git status" para saber que falta subir al repositorio
//"git add ." es para agregar los archivos a subir
//"git commit -m "descripcion del commit" define y prepara los archivos a subir al repo
//"git push" es para subir los archivos definidos sean nuevos o actualizados
