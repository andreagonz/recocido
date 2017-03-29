package main

import (
	"fmt"
	"math/rand"
	con "github.com/andreagonz/recocido/conexion"
	heu "github.com/andreagonz/recocido/heuristica"
	imp "github.com/andreagonz/recocido/implementacion"
)

/*
func ProblemaAleatorio(t int, ciudades *[]imp.Ciudad, distancias *[][]float64, r *rand.Rand) []int {
	p := make([]int, t)
	x := r.Intn(len(*ciudades))
	for i := 0; i < t; i++ {
		g := 1
		y := r.Intn(len(*ciudades))
		c := 0
		for (*distancias)[x % len(*ciudades)][y % len(*ciudades)] == 0.0 {
			y++
			c++
			if c == len(*ciudades) {
				x = p[i - g]
				c = 0
				g--
			}
		}
		p[i] = x % len(*ciudades)
		x = y % len(*ciudades)
		fmt.Println(x)
	}
	return p
}
*/

func ProblemaAleatorio(t int, ciudades *[]imp.Ciudad, distancias *[][]float64, r *rand.Rand) []int {
	var pila imp.Pila
	p := make([]int, t)
	x := r.Intn(len(*ciudades))
	i := 0
	for i < t {
		y := r.Intn(len(*ciudades))
		for (*distancias)[x % len(*ciudades)][y % len(*ciudades)] == 0.0 {
			y++
			if x % len(*ciudades) == y % len(*ciudades) {
				// i--
				if pila, x = pila.Pop(); pila == nil {
					x = r.Intn(len(*ciudades))
					// i = 0
				}
			}
		}
		p[i] = x % len(*ciudades)
		x = y % len(*ciudades)
		pila.Push(x)
		i++
	}
	return p
}

func ImprimeLote(l heu.Lote) {
	for i := 0; i < len(l.Soluciones); i++ {
		fmt.Println(l.Soluciones[i].Str())
	}
}

func main() {

	seed := int64(7)
	numCiudades := 277
	tProblema := 9
	tLote := 200
	t := 8.0
	p := 0.9
	ep := 0.9
	et := 0.1
	e := 0.1
	phi := 0.95
	c := 10
	
	r := rand.New(rand.NewSource(seed))
	ciudades := con.LeeCiudades(numCiudades)
	distancias, _ := con.LeeConexiones(numCiudades)	
	//problema := ProblemaAleatorio(tProblema, &ciudades, &distancias, r)

	problema := []int{0,1,6,13,25,26,30,32,36}
	imp.SetDistancias(&distancias)
	imp.SetCiudades(&ciudades)
	imp.SetProblema(&problema)
	imp.SetC(c)
	imp.MaxAvg()

	sol := imp.Ruta{Ciudades : problema}
	fmt.Println(sol.Str())
	fmt.Println()
	lote, _, _ := heu.CalculaLote(00.0, &sol, tLote, r)	
	t = heu.TemperaturaInicial(&sol, t, p, ep, et, tProblema, r)
	lote, mejorS := heu.AceptacionPorHumbrales(t, &sol, e, tProblema, r, phi)
	fmt.Println(mejorS.Str())
	fmt.Println(mejorS.ObtenFunObj())
	fmt.Println(mejorS.ObtenFun())
	fmt.Println(lote.PorcentajeFactibles())
}
