package main

import (
	"fmt"
	"math/rand"
	con "github.com/andreagonz/recocido/conexion"
	heu "github.com/andreagonz/recocido/heuristica"
	imp "github.com/andreagonz/recocido/implementacion"
)

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
		fmt.Println(l.Soluciones[i])
	}
}

func main() {

	seed := int64(11)
	numCiudades := 277
	//tProblema := 78
	tLote := 500
	p := 0.9
	ep := 0.01
	et := 0.01
	e := 0.01
	phi := 0.9
	c := 5
	
	r := rand.New(rand.NewSource(seed))
	ciudades := con.LeeCiudades(numCiudades)
	distancias, _ := con.LeeConexiones(numCiudades)
	//problema := ProblemaAleatorio(tProblema, &ciudades, &distancias, r)

	/*
	problema := []int{1, 5, 9, 12, 16, 22, 23, 29, 30, 31, 39, 48, 52, 56, 58, 62, 65, 66, 70, 75, 80, 84, 86, 90, 92, 94, 95, 101, 107, 117, 119, 122, 133, 135, 143, 144, 146, 147, 150, 158, 159, 160, 166, 167, 176, 178, 179, 185, 186, 188, 190, 191, 194, 198, 200, 203, 207, 209, 213, 215, 216, 220, 221, 224, 227, 232, 233, 235, 238, 241, 244, 248, 250, 254, 264, 266, 274, 276}
	
	for i := 0; i < len(problema); i++ {
		problema[i] = problema[i] - 1
	}
*/	
	problema := []int{25,36,13,6,0,26,30,1,32}
	
	imp.SetDistancias(&distancias)
	imp.SetCiudades(&ciudades)
	imp.SetProblema(&problema)
	imp.SetC(c)
	imp.MaxAvg()
	
	sol := imp.Ruta{Ciudades : problema}
	fmt.Println(sol.Str())
	fmt.Println()	
	sol.CalculaFun()

	lote, mejor, p := heu.Recocido(&sol, e, ep, et, p, tLote, r, phi)
	
	//fmt.Println("APH calculada")
	fmt.Println("Mejor sol")
	fmt.Println(mejor.Str())
	fmt.Println("Mejor dist")
	fmt.Println(mejor.ObtenFunObj())
	fmt.Println("Mejor costo")
	fmt.Println(mejor.ObtenFun())
	fmt.Println("Costo prom")
	fmt.Println(p)
	fmt.Println("% factibles")
	fmt.Println(lote.PorcentajeFactibles())
	ImprimeLote(lote)
}
