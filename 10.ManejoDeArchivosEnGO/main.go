package main

import (
	"fmt"

	"github.com/xfred19x/godesde0/10.ManejoDeArchivosEnGO/files"
)

func main() {

	fmt.Println("-------------")
	files.GrabaTabla()

	fmt.Println("-------------")
	files.SumaTabla()

	//fmt.Println("-------------")
	//files.LeoArchivos()

	fmt.Println("-------------")
	files.LeoArchivos2()

}
