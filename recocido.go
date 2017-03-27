package main

import (
	"fmt"
	"math/rand"
	con "github.com/andreagonz/recocido/conexion"
	heu "github.com/andreagonz/recocido/heuristica"
	imp "github.com/andreagonz/recocido/implementacion"
)

func RutaAleatoria(r *rand.Rand, n int, ciudades *[]imp.Ciudad) imp.Ruta {
	//var s imp.Ruta
	ciud := make([]int, n)
	for i := 0; i < n; i++ {
		a := r.Intn(len(*ciudades))
		ciud[i] = (*ciudades)[a].Id
	}
	return imp.Ruta{Ciudades : ciud}
}

func main() {

	seed := int64(9)
	numCiudades := 30
	tLote := 200
	t := 8.0
	p := 0.85
	ep := 0.1
	et := 0.1
	e := 5.5
	phi := 0.2
	
	ciudades := con.LeeCiudades(277)
	distanciasI, sum := con.LeeConexiones(277)
	distancias := make([][]float64, len(distanciasI))

	for i := 0; i < 277; i++ {
		distancias[i] = make([]float64, 277)
	}
	
	for i := 0; i < 277; i++ {
		for j := 0; j < 277; j++ {
			if distanciasI[i][j] > 0.0 {
				distancias[i][j] = distanciasI[i][j]
			} else {
				distancias[i][j] = sum * 2
			}
		}
	}

	imp.SetDistancias(&distancias)
	imp.SetDistanciasI(&distanciasI)
	imp.SetCiudades(&ciudades)
	
	r := rand.New(rand.NewSource(seed))

	sol := RutaAleatoria(r, numCiudades, &ciudades)
	
	//fmt.Println(sol.Str())
	
	lote, _, _ := heu.CalculaLote(00.0, &sol, tLote, r)
	fmt.Println("cl")
	t = heu.TemperaturaInicial(&sol, t, p, ep, et, numCiudades, r)
	fmt.Println("ti")
	lote, _ = heu.AceptacionPorHumbrales(t, &sol, e, numCiudades, r, phi)
	fmt.Println("aph")
	fmt.Println(lote)
	/*
	fact := 0
	for i := 0; i < len(lote.Soluciones); i++ {
		fmt.Println(lote.Soluciones[i].Str())				
		if lote.Soluciones[i].EsFactible() {
			fact++
		}
	}
	fmt.Println(fact)
*/
}
