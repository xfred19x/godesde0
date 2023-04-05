package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

// Ya habiendo definido las funciones, estas son detectadas por la interfaz que hace referencia a esta estructura
func (h *Hombre) Respirar()    { h.respirando = true }
func (h *Hombre) Comer()       { h.comiendo = true }
func (h *Hombre) Pensar()      { h.pensando = true }
func (h *Hombre) Sexo() string { return "Hombre" }

func (hombre *Hombre) AddHombre(edad int, altura float32, peso float32, respirando bool, pensando bool, comiendo bool, vivo bool) {
	hombre.edad = edad
	hombre.altura = altura
	hombre.peso = peso
	hombre.respirando = respirando
	hombre.pensando = pensando
	hombre.comiendo = comiendo
	hombre.vivo = vivo
}
