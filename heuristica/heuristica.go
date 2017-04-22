package recocido

import (
	"math/rand"
	"math"
	"container/list"
)

// Solucion es una interfaz para una solucion del problema
type Solucion interface {
	ObtenVecino(*rand.Rand, float64, bool) bool
	CalculaFun()
	ObtenFun() float64
	ObtenFunObj() float64
	Str() string
	EsFactible() bool
	Copia() Solucion
}

// CalculaLote calcula las soluciones de un lote.
// Regresa la última solución obtenida junto con la mejor y el promedio de los costos.
func CalculaLote(t float64, solucion Solucion, mejor Solucion, l int, rand *rand.Rand, lista *list.List, bGrafica bool) (float64, Solucion, Solucion) {
	var c = 0
	var r = 0.0
	i := 0
	for c < l {
		if solucion.ObtenVecino(rand, t, i > l * l) || (i > l * l) {
			if bGrafica {
				(*lista).PushBack(solucion.ObtenFun())
			}
			c++
			r += solucion.ObtenFun()
			if mejor.ObtenFun() > solucion.ObtenFun() {
				mejor = solucion.Copia()
			}
		}
		i++
	}
	return r/float64(l), solucion, mejor
}

// AceptacionPorUmbrales ejecuta el algoritmo de aceptación por umbrales a partir de una
// temperatura y una solución dadas.
// Regresa la mejor solución y el costo promedio del último lote.
func AceptacionPorUmbrales(t float64, s Solucion, mejor Solucion, e float64, ep float64, l int, rand *rand.Rand, phi float64, lista *list.List, bGrafica bool) (Solucion, float64) {
	p := math.MaxFloat64
	for t > e {
		r := 0.0
		i := 0
		for math.Abs(p - r) > ep && i < l * l {
			r = p
			p, s, mejor = CalculaLote(t, s, mejor, l, rand, lista, bGrafica)
			i++
		}
		t *= phi
	}
	return mejor, p
}

// TemperaturaInicial calcula la temperatura con la que se ejecutara la aceptación por umbrales.
// Regresa la temperatura.
func TemperaturaInicial(s Solucion, t float64, p float64, ep float64, et float64, n int, rand *rand.Rand) float64 {
	r := PorcentajeAceptados(s, t, n, rand)
	t1 := 0.0
	t2 := 0.0
	if math.Abs(p - r) <= ep {
		return t
	}
	if r < p {
		for r < p {
			t = 2 * t
			r = PorcentajeAceptados(s, t, n, rand)
		}
		t1 = t / 2
		t2 = t
	} else {
		for r > p {
			t = t / 2
			r = PorcentajeAceptados(s, t, n, rand)
		}
		t1 = t
		t2 = 2 * t
	}
	return BusquedaBinaria(s, t1, t2, p, ep, et, n, rand)
}

// PorcentajeAceptados calcula el porcentaje de soluciones aceptadas en un rango.
func PorcentajeAceptados(sol Solucion, t float64, n int, rand *rand.Rand) float64 {
	c := 0.0
	s := sol.Copia()
	for i := 1; i <= n; i++ {
		if s.ObtenVecino(rand, t, true) {
			c++
		}
	}
	return c/float64(n)
}

// BusquedaBinaria ayuda a calcular la temperatura inicial haciendo una búsqueda binaria.
func BusquedaBinaria(s Solucion, t1 float64, t2 float64, p float64, ep float64, et float64, n int, rand *rand.Rand) float64 {
	tm := (t1 + t2) / 2
	if t2 - t1 < et {
		return tm
	}
	r := PorcentajeAceptados(s, tm, n, rand)
	if math.Abs(p - r) < ep {
		return tm
	}
	if r > p {
		return BusquedaBinaria(s, t1, tm, p, ep, et, n, rand)
	}
	return BusquedaBinaria(s, tm, t2, p, ep, et, n, rand)
}

// Recocido ejecuta el algoritmo de recocido simulado con aceptación por humbrales.
func Recocido(s Solucion, e float64, ep float64, et float64, p float64, l int, rand *rand.Rand, phi float64, lista *list.List, bGrafica bool) (Solucion, float64) {
	t := TemperaturaInicial(s, 8, p, ep, et, l, rand)
	return AceptacionPorUmbrales(t, s.Copia(), s.Copia(), e, ep, l, rand, phi, lista, bGrafica)
}
