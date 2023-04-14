package webserver

import "net/http" //libreria que tiene que ver con servidor http

//Esta funcion sera mi webserver
func MiWebServer() {
	//colocaremos un manejador que recibe la peticion
	//para todo el que llegue "/", que ejecute la funcion "home"
	http.HandleFunc("/", home)
	//vamos a escuchar el puerto "3000"
	http.ListenAndServe(":3000", nil)

}

//funcion que va mostrar contenido o procesar una peticion de http
func home(w http.ResponseWriter, r *http.Request) {
	//llamaremos la funcion "ServeFile"
	//permite servir a nuestra web un archivo
	http.ServeFile(w, r, "./webserver/index.html")

}
