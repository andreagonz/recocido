package recocido

import (
	_ "fmt"
	"strconv"
	"math/rand"
	"github.com/andreagonz/recocido/heuristica"
)

var distancias *[][]float64
var distanciasI *[][]float64

func SetDistancias(p *[][]float64) {
	distancias = p
}

func SetDistanciasI(p *[][]float64) {
	distanciasI = p
}

type Ciudad struct {
	Id int
	Nombre string
	Pais string
	Poblacion int
	Latitud float64
	Longitud float64	
}

type Ruta struct {
	Ciudades []Ciudad
	fun float64
}

type Conexiones struct {
	Distancias []float64
}

func (r Ruta) Str() string {
	s := ""
	s += "{"
	for i := 0; i < len(r.Ciudades); i++ {
		s += strconv.Itoa(r.Ciudades[i].Id) + " "
	}
	s += "}"
	return s
}

func (r Ruta) ObtenFun() float64 {
	return r.fun
}

func (r Ruta) AsignaFun(f float64) {
	r.fun = f
}

func (ruta Ruta) ObtenVecino(rand *rand.Rand) recocido.Solucion {
	//fmt.Println(ruta.Str())
	i := rand.Intn(len(ruta.Ciudades))
	j := rand.Intn(len(ruta.Ciudades))
	var nruta Ruta
	nruta.Ciudades = make([]Ciudad, len(ruta.Ciudades))
	for i := 0; i < len(ruta.Ciudades); i++ {
		nruta.Ciudades[i] = ruta.Ciudades[i]
	}
	a := nruta.Ciudades[i]
	nruta.Ciudades[i] = nruta.Ciudades[j]
	nruta.Ciudades[j] = a
	//fmt.Println(ruta.Str())
	return &nruta
}

func (r *Ruta) CalculaFun() {
	f := float64(0.0)
	for i := 0; i < len(r.Ciudades) - 1; i++ {
		f += (*distancias)[r.Ciudades[i].Id][r.Ciudades[i + 1].Id]
	}
	r.fun = f
}

func(r Ruta) EsFactible() bool {
	bool := true
	for j := 0; j < len(r.Ciudades) - 1; j++ {
		if (*distanciasI)[r.Ciudades[j].Id][r.Ciudades[j + 1].Id] == 0.0 {
			bool = false
		}
	}
	return bool
}
