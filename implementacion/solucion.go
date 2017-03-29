package recocido

import (
	_ "fmt"
	"strconv"
	"math/rand"
	"github.com/andreagonz/recocido/heuristica"
)

var distancias *[][]float64
var ciudades *[]Ciudad
var problema *[]int
var max float64
var avg float64
var c int

func SetDistancias(p *[][]float64) {
	distancias = p
}

func SetCiudades(p *[]Ciudad) {
	ciudades = p
}

func SetC(i int) {
	c = i
}

func SetProblema(p *[]int) {
	problema = p
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
	Ciudades []int
	funObj float64
	fun float64
}

type Conexiones struct {
	Distancias []float64
}

func (r Ruta) Str() string {
	s := ""
	s += "{"
	for i := 0; i < len(r.Ciudades); i++ {
		s += "(" + strconv.Itoa(r.Ciudades[i]) + ": " + (*ciudades)[r.Ciudades[i]].Nombre + ") "
		if i < len(r.Ciudades) - 1 {
			s += strconv.FormatFloat((*distancias)[r.Ciudades[i]][r.Ciudades[i + 1]], 'f', -1, 64) + " "
		}
	}
	s += "}"
	return s
}

func (r Ruta) ObtenFun() float64 {
	return r.fun
}


func (r Ruta) ObtenFunObj() float64 {
	return r.funObj
}

func (r Ruta) AsignaFun(f float64) {
	r.fun = f
}

func (ruta Ruta) ObtenVecino(rand *rand.Rand) recocido.Solucion {
	//fmt.Println(ruta.Str())
	i := rand.Intn(len(ruta.Ciudades))
	j := rand.Intn(len(ruta.Ciudades))
	var nruta Ruta
	nruta.Ciudades = make([]int, len(ruta.Ciudades))
	for i := 0; i < len(ruta.Ciudades); i++ {
		nruta.Ciudades[i] = ruta.Ciudades[i]
	}
	a := nruta.Ciudades[i]
	nruta.Ciudades[i] = nruta.Ciudades[j]
	nruta.Ciudades[j] = a
	//fmt.Println(ruta.Str())
	return &nruta
}

func(r Ruta) EsFactible() bool {
	bool := true
	for j := 0; j < len(r.Ciudades) - 1; j++ {
		if (*distancias)[r.Ciudades[j]][r.Ciudades[j + 1]] == 0.0 {
			bool = false
		}
	}
	return bool
}

func MaxAvg() {
	n := 0.0
	p := 0.0
	for i := 1; i < len(*problema); i++ {
		if (*distancias)[(*problema)[i - 1]][(*problema)[i]] > 0.0 {
			if (*distancias)[(*problema)[i - 1]][(*problema)[i]] > max {
				max = (*distancias)[(*problema)[i - 1]][(*problema)[i]]
			}
			p += (*distancias)[(*problema)[i - 1]][(*problema)[i]]
			n++
		}
	}
	avg = p / n
}

func (r *Ruta) CalculaFun() {
	f := 0.0
	for i := 1; i < len(r.Ciudades) - 1; i++ {
		if (*distancias)[(r.Ciudades)[i - 1]][(r.Ciudades)[i]] > 0.0 {
			f += (*distancias)[r.Ciudades[i - 1]][r.Ciudades[i]]
		} else {
			f += max * float64(c)
		}
	}
	r.funObj = f
	r.fun = f / (avg * float64((len(r.Ciudades)) - 1))
}
