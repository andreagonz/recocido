package main

import (
	"os"
	"fmt"
	"math/rand"
	"strconv"
	"bytes"
	"strings"
	"container/list"
	con "github.com/andreagonz/recocido/conexion"
	heu "github.com/andreagonz/recocido/heuristica"
	imp "github.com/andreagonz/recocido/implementacion"
	util "github.com/andreagonz/recocido/util"
)

func main() {

	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Uso: ./recocido <archivo.tsp> <params.txt> [ops]")
	} else {

		numCiudades := 277
		params := strings.Fields(util.LeeArchivo(args[1]))
		var seed int64
		seed, _ = strconv.ParseInt(params[0], 10, 64)
		tLote, _ := strconv.Atoi(params[1])
		p, _ := strconv.ParseFloat(params[2], 64)
		ep, _ := strconv.ParseFloat(params[3], 64)
		et, _ := strconv.ParseFloat(params[4], 64)
		e, _ := strconv.ParseFloat(params[5], 64)
		phi, _ := strconv.ParseFloat(params[6], 64)
		c, _ := strconv.Atoi(params[7])
		bGrafica := false
		bMapa := false

		for i := 2; i < len(args); i++ {
			if args[i] == "-g" {
				bGrafica = true
			}
			if args[i] == "-m" {
				bMapa = true
			}
		}
		
		ciudades := con.LeeCiudades(numCiudades)
		distancias := con.LeeConexiones(numCiudades)

		r := rand.New(rand.NewSource(seed))

		problema := util.CadenaARuta(util.LeeArchivo(args[0]))
		
		imp.SetDistancias(&distancias)
		imp.SetCiudades(&ciudades)
		imp.SetProblema(&problema)
		imp.SetC(c)
		imp.MaxAvg()
		
		sol := imp.Ruta{Ciudades : problema}
		sol.CalculaFun()

		listaFun := list.New()		
		mejor, _ := heu.Recocido(&sol, e, ep, et, p, tLote, r, phi, listaFun, bGrafica)

		if bGrafica {
			var buffer bytes.Buffer
			for f := listaFun.Front(); f != nil; f = f.Next() {
				buffer.WriteString(fmt.Sprintf("E: %.9f\n", f.Value))
			}
			util.EscribeArchivo(buffer.String(), "costos.txt")
			util.GraficaCosto(listaFun)
		}

		if bMapa {
			mapa := util.Mapa(mejor.Str(), &ciudades)
			util.EscribeArchivo(mapa, "mapa.html")
		}

		fmt.Println("Mejor solución")
		fmt.Println("Factible: " + strconv.FormatBool(mejor.EsFactible()))
		fmt.Println("Costo: " + strconv.FormatFloat(mejor.ObtenFun(), 'f', -1, 64))
		fmt.Println("Distancia: " + strconv.FormatFloat(mejor.ObtenFunObj(), 'f', -1, 64))
		util.EscribeArchivo(mejor.Str(), "ruta.tsp")
	}
	
}
