package recocido

import (
	_"fmt"
	"math/rand"
	"math"
)

type Solucion interface {
	ObtenVecino(*rand.Rand) Solucion
	CalculaFun()
	ObtenFun() float64
	ObtenFunObj() float64
	AsignaFun(float64)
	Str() string
	EsFactible() bool
}

type Lote struct {
	Soluciones []Solucion
}

func (l Lote) PorcentajeFactibles() float64 {
	c := 0.0
	for i := 0; i < len(l.Soluciones); i++ {
		if l.Soluciones[i].EsFactible() {
			c++
		}
	}
	if c == 0.0 {
		return 0
	}
	return c / float64(len(l.Soluciones))
}

func CalculaLote(t float64, solucion Solucion, l int, rand *rand.Rand) (Lote, float64, Solucion) {
	var c = 0
	var r = 0.0
	var s Solucion
	var lote Lote
	lote.Soluciones = make([]Solucion, l)
	solucion.CalculaFun()
	i := 0
	for c < l {
		s = solucion.ObtenVecino(rand)
		s.CalculaFun()
		if s.ObtenFun() <= solucion.ObtenFun() + t {
			solucion = s
			lote.Soluciones[c] = solucion
			c++
			r += solucion.ObtenFun()
		}
		i++
	}
	return lote, r/float64(l), s
}

func AceptacionPorHumbrales(t float64, s Solucion, e float64, l int, rand *rand.Rand, phi float64) (Lote, Solucion) {
	p := 1000000.0
	var lote Lote
	for t > e {
		r := 0.0
		for math.Abs(p - r) > e {
			r = p
			lote, p, s = CalculaLote(t, s, l, rand)			
		}
		t *= phi
	}
	return lote, s
}

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
	//fmt.Println("bb")
	return BusquedaBinaria(s, t1, t2, p, ep, et, n, rand)
}

func PorcentajeAceptados(s Solucion, t float64, n int, rand *rand.Rand) float64 {
	c := 0.0
	for i := 1; i <= n; i++ {
		//fmt.Println(t)
		r := s.ObtenVecino(rand)
		r.CalculaFun()
		s.CalculaFun()
		if r.ObtenFun() <= s.ObtenFun() + t {
			c++
		}
		s = r
	}
	return c/float64(n)
}

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
	} else {
		return BusquedaBinaria(s, tm, t2, p, ep, et, n, rand)
	}
}
