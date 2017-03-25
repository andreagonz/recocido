package recocido

import (
	"fmt"
	"math/rand"
)

type Solucion interface {
	ObtenVecino(*rand.Rand) Solucion
	CalculaFun()
	ObtenFun() float64
	AsignaFun(float64)
	Str() string
	EsFactible() bool
}

type Temperatura interface {
	TemperaturaInicial(solucion Solucion, temperatura Temperatura, porcentaje float64)
	PorcentajeAceptados(solucion Solucion, temperatura Temperatura)
}

type Recocido interface {
	AceptacionPorHumbrales(temperatura Temperatura, solucion Solucion)
}

type Lote struct {
	Soluciones []Solucion
}

func (lote Lote) CalculaFunciones() {
	max := 0.0
	min := 100000000.0
	for i := 0; i < len(lote.Soluciones); i++ {
		if max < lote.Soluciones[i].ObtenFun() {
			max = lote.Soluciones[i].ObtenFun()
		}
		if min > lote.Soluciones[i].ObtenFun() {
			min = lote.Soluciones[i].ObtenFun()
		}
	}
	for i := 0; i < len(lote.Soluciones); i++ {
		lote.Soluciones[i].AsignaFun((lote.Soluciones[i].ObtenFun() - min) / (max - min))
	}
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
		s := solucion.ObtenVecino(rand)
		s.CalculaFun()
		fmt.Println("s' " + s.Str())
		fmt.Println("s " + solucion.Str())
		fmt.Println(c)
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
