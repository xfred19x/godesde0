package interfaces

//Al igual que las estructuras una interface se crea con la palabra "type"
type Humano interface {
	//No se puede colocar variables, propiedades ni tipos de datos
	//solo una definicion de funciones
	Respirar()
	Pensar()
	Comer()
	Sexo() string
}
