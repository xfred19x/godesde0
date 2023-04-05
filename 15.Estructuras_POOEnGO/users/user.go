package users

import (
	"fmt"
	"time"

	"github.com/xfred19x/godesde0/15.Estructuras_POOEnGO/modelos"
)

func AltaUsuario() {

	//creamos un objeto usuario
	u := new(modelos.User)
	//agregamos un nuevo usuario con la estrctura y datos que corresponde
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
