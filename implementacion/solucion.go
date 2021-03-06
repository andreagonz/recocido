package recocido

import (
	"strconv"
	"math/rand"
	"math"
	"github.com/andreagonz/recocido/heuristica"
)

var distancias *[][]float64
var ciudades *[]Ciudad
var problema *[]int
var max float64
var avg float64
var c int

// SetDistancias recibe un apuntador a la matríz de adyacencias de
// las ciudades y lo guarda.
func SetDistancias(p *[][]float64) {
	distancias = p
}

// SetCiudades recibe un apuntador a la lista de ciudades
// y lo guarda.
func SetCiudades(p *[]Ciudad) {
	ciudades = p
}

// SetC recibe guarda el valor de la C a utilizar
func SetC(i int) {
	c = i
}

// SetProblema recibe un apuntador a la instancia de ciudades
// a tomar en cuenta para el problema actual.
func SetProblema(p *[]int) {
	problema = p
}

// Ciudad es una estructura que define a una ciudad para el problema
// del agente viajero.
type Ciudad struct {
	Id int
	Nombre string
	Pais string
	Poblacion int
	Latitud float64
	Longitud float64	
}

// Ruta es una solución posible al problema del agente viajero.
type Ruta struct {
	Ciudades []int
	funObj float64
	fun float64
}

// Str devuelve una representación en cadena de la ruta.
func (r Ruta) Str() string {
	s := ""
	for i := 0; i < len(r.Ciudades); i++ {
		s += strconv.Itoa(r.Ciudades[i] + 1)
		if i < len(r.Ciudades) - 1 {
			s += ", "
		}
	}
	return s
}

// ObtenFun devuelve el costo de la ruta r.
func (r Ruta) ObtenFun() float64 {
	return r.fun
}

// ObtenFunObj devuelve la distancia de la ruta r.
func (r Ruta) ObtenFunObj() float64 {
	return r.funObj
}

// ObtenVecino devuelve una ruta vecina a la ruta r.
func (r *Ruta) ObtenVecino(rand *rand.Rand, t float64, cond bool) bool {
	esMejor := false
	i := rand.Intn(len(r.Ciudades))
	j := rand.Intn(len(r.Ciudades))
	for j == i {
		j = rand.Intn(len(r.Ciudades))
	}
	
	fo := r.funObj

	may := int(math.Max(float64(i), float64(j)))
	men := int(math.Min(float64(i), float64(j)))

	if may - men > 1 {
		fo -= r.getDistancia(men, men + 1)
		fo -= r.getDistancia(may - 1, may)
		fo += r.getDistancia(may, men + 1)
		fo += r.getDistancia(may - 1, men)	
	}		
	if may < len(r.Ciudades) - 1 {
		fo -= r.getDistancia(may, may + 1)
		fo += r.getDistancia(men, may + 1)
	}
	if men > 0 {
		fo -= r.getDistancia(men - 1, men)
		fo += r.getDistancia(men - 1, may)
	}

	f := fo / (avg * float64((len(r.Ciudades)) - 1))

	if f <= r.fun + t {
		esMejor = true
	}

	if esMejor || cond {
		a := r.Ciudades[i]
		r.Ciudades[i] = r.Ciudades[j]
		r.Ciudades[j] = a
		r.fun = f
		r.funObj = fo
	}

	return esMejor
}

// EsFactible dice si la ruta r es factible.
func(r Ruta) EsFactible() bool {
	bool := true
	for j := 0; j < len(r.Ciudades) - 1; j++ {
		if (*distancias)[r.Ciudades[j]][r.Ciudades[j + 1]] == 0.0 {
			bool = false
		}
	}
	return bool
}

// MaxAvg calcula la distancia máxima en el problema y
// el promedio de las distancias.
func MaxAvg() {
	n := 0.0
	p := 0.0
	for i := 0; i < len(*problema); i++ {
		for j := i + 1; j < len(*problema); j++ {
			if (*distancias)[(*problema)[i]][(*problema)[j]] > 0.0 {
				if (*distancias)[(*problema)[i]][(*problema)[j]] > max {
					max = (*distancias)[(*problema)[i]][(*problema)[j]]
				}
				p += (*distancias)[(*problema)[i]][(*problema)[j]]
				n++
			}
		}
	}
	avg = p / n
}

// Regresa la distancia* entre la ciudad i y la ciudad j
func (r Ruta) getDistancia(i int, j int) float64 {
	if (*distancias)[(r.Ciudades)[i]][(r.Ciudades)[j]] > 0.0 {
		return (*distancias)[(r.Ciudades)[i]][(r.Ciudades)[j]]
	} else {
		return max * float64(c)
	}
}

// CalculaFun calcula la función de costo y la distancia de
// la ruta r.
func (r *Ruta) CalculaFun() {
	f := 0.0
	for i := 1; i < len(r.Ciudades); i++ {
		f += r.getDistancia(i - 1, i)
	}
	r.funObj = f
	r.fun = f / (avg * float64((len(r.Ciudades)) - 1))
}

func (r Ruta) Copia() recocido.Solucion {
	c := make([]int, len(r.Ciudades))
	for i := 0; i < len(r.Ciudades); i++ {
		c[i] = r.Ciudades[i]
	}
	ruta := Ruta{c, r.funObj, r.fun}
	return &ruta
}
