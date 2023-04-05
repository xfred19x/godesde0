package modelos

import (
	"time"
)

// para crear estructura tiene la palabra reservada "type" y "struct"
type User struct {
	Id       int
	Name     string
	CreateAt time.Time
	Status   bool
}

// para crear funcionalidad con la estructura, en la funcion debemos referirlo con "usuario *NombreEstructura"
// la palabra "usuario" no es obligatorio puede ser cualquiera, lo que importa que apunte a la estructura "User"
func (usuario *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	//Aqui relacionamos los parametros de entrada con la estructura "User"
	usuario.Id = id
	usuario.Name = name
	usuario.CreateAt = createdAt
	usuario.Status = status
}
