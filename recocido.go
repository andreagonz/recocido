package main

import (
	"fmt"
	"math/rand"
	con "github.com/andreagonz/recocido/conexion"
	heu "github.com/andreagonz/recocido/heuristica"
	imp "github.com/andreagonz/recocido/implementacion"
)

var distancias [][]float64

func main() {

	seed := int64(9)
	numCiudades := 30
	tLote := 200
	
	ciudades := con.LeeCiudades(277)
	distanciasI, sum := con.LeeConexiones(277)
	distancias = make([][]float64, len(distanciasI))
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
	s := rand.NewSource(seed)
	r := rand.New(s)
	
	ciud := make([]imp.Ciudad, numCiudades)
	for i := 0; i < numCiudades; i++ {		
		a := r.Intn(len(ciudades))
		ciud[i] = ciudades[a]
	}
	
	sol := imp.Ruta{Ciudades : ciud}
	//fmt.Println(sol.Str())
	
	lote, _, _ := heu.CalculaLote(00.0, &sol, tLote, r)

	fact := 0
	for i := 0; i < len(lote.Soluciones); i++ {
		fmt.Println(lote.Soluciones[i].Str())				
		if lote.Soluciones[i].EsFactible() {
			fact++
		}
	}
	fmt.Println(fact)
}
