package modelos

//Definimos el modelo "Mujer", pero hacemos referencia "Hombre" para que herede todas las propiedades.
type Mujer struct {
	Hombre
}

func (m *Mujer) Respirar()    { m.respirando = true }
func (m *Mujer) Comer()       { m.comiendo = true }
func (m *Mujer) Pensar()      { m.pensando = true }
func (m *Mujer) Sexo() string { return "Mujer" }

func (mujer *Mujer) AddMujer(edad int, altura float32, peso float32, respirando bool, pensando bool, comiendo bool, vivo bool) {
	mujer.edad = edad
	mujer.altura = altura
	mujer.peso = peso
	mujer.respirando = respirando
	mujer.pensando = pensando
	mujer.comiendo = comiendo
	mujer.vivo = vivo
}
