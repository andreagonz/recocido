package main

import (
	// "os"
	// "fmt"
	"math/rand"
	// "strconv"
	// "strings"
	con "github.com/andreagonz/recocido/conexion"
	// heu "github.com/andreagonz/recocido/heuristica"
	imp "github.com/andreagonz/recocido/implementacion"
	util "github.com/andreagonz/recocido/util"
)

// ProblemaAleatorio genera un subconjunto aleatorio de ciudades
func ProblemaAleatorio(t int, ciudades *[]imp.Ciudad, distancias *[][]float64, r *rand.Rand) []int {
	p := make([]int, t)
	x := r.Intn(len(*ciudades))
	i := 0
	for i < t {
		y := r.Intn(len(*ciudades))
		for (*distancias)[x][y % len(*ciudades)] == 0.0 {
			y++
		}
		p[i] = x
		x = y % len(*ciudades)
		i++
	}
	return p
}

func main() {
	numCiudades := 277
	ciudades := con.LeeCiudades(numCiudades)
	s := "1, 188, 95, 144, 233, 9, 80, 250, 146, 147, 276, 274, 58, 224, 158, 31, 213, 22, 185, 29, 266, 133, 227, 254, 84, 66, 159, 216, 5, 143, 248, 23, 176, 122, 209, 62, 90, 203, 235, 52, 215, 70, 75, 198, 220, 92, 179, 16, 241, 190, 244, 186, 232, 238, 48, 12, 107, 264, 191, 221, 150, 86, 166, 135, 119, 94, 207, 56, 30, 65, 101, 200, 167, 160, 178, 194, 39, 117"
	util.EscribeArchivo(util.Mapa(s, &ciudades), "carlos.html")
}

/*
func main() {

	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("Uso: ./recocido <archivo.tsp> <params.txt> <directorio>")
	} else {

		numCiudades := 277
		// tProblema := 50
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
		direc := args[2]
		
		ciudades := con.LeeCiudades(numCiudades)
		distancias := con.LeeConexiones(numCiudades)

		r := rand.New(rand.NewSource(seed))
		//	problema := ProblemaAleatorio(tProblema, &ciudades, &distancias, r)

		problema := util.CadenaARuta(util.LeeArchivo(args[0]))
		
		imp.SetDistancias(&distancias)
		imp.SetCiudades(&ciudades)
		imp.SetProblema(&problema)
		imp.SetC(c)
		imp.MaxAvg()
		
		sol := imp.Ruta{Ciudades : problema}
		sol.CalculaFun()
		//fmt.Println(sol.Str())

		var mejorG heu.Solucion
		var prom float64
		var mejor heu.Solucion
		mejorGI := 0
		
		util.EscribeArchivo(util.CadenaParametros(seed, tLote, p, ep, et, e, phi, c), direc + "parametros.txt")
		sAceptadas := util.LeeArchivo(direc + "aceptadas.txt")
		sFunciones := util.LeeArchivo(direc + "funciones.txt")
		sDistancias := util.LeeArchivo(direc + "distancias.txt")

		for i := 0; i < 30; i++ {
			r = rand.New(rand.NewSource(int64(i)))
			mejor, prom = heu.Recocido(&sol, e, ep, et, p, tLote, r, phi)
			if i == 0 {
				mejorG = mejor
			} else if mejorG.ObtenFun() > mejor.ObtenFun() {
				mejorG = mejor
				mejorGI = i
			}
			fmt.Println(i)
			sMejor := ""
			sMejor += "Factible: " + strconv.FormatBool(mejor.EsFactible()) + "\n"
			sAceptadas += strconv.FormatBool(mejor.EsFactible()) + " "
			sMejor += "Costo: " + strconv.FormatFloat(mejor.ObtenFun(), 'f', -1, 64) + "\n"
			sFunciones += fmt.Sprintf("%.9f ", mejor.ObtenFun())
			sMejor += "Distancia: " + strconv.FormatFloat(mejor.ObtenFunObj(), 'f', -1, 64) + "\n"
			sDistancias += fmt.Sprintf("%f ", mejor.ObtenFunObj())
			fmt.Println(sMejor)
			util.EscribeArchivo(mejor.Str(), fmt.Sprintf(direc + "%02d.tsp", i))
			util.EscribeArchivo(sMejor, fmt.Sprintf(direc + "%02d.fun", i))
		}

		util.EscribeArchivo(fmt.Sprintf("%02d.tsp", mejorGI), direc + "mejor.txt")
		util.EscribeArchivo(util.Mapa(mejorG.Str(), &ciudades), direc + "mejor.html")
		util.EscribeArchivo(sAceptadas, direc + "aceptadas.txt")
		util.EscribeArchivo(sFunciones, direc + "funciones.txt")
		util.EscribeArchivo(sDistancias, direc + "distancias.txt")
		
		fmt.Println("Promedio: " + strconv.FormatFloat(prom, 'f', -1, 64))
		//fmt.Println(mejorG.Str())
		fmt.Println("Factible:" + strconv.FormatBool(mejorG.EsFactible()))
		fmt.Println("Costo: " + strconv.FormatFloat(mejorG.ObtenFun(), 'f', -1, 64))
		fmt.Println("Distancia: " + strconv.FormatFloat(mejorG.ObtenFunObj(), 'f', -1, 64))
		fmt.Println()
	}	
}
*/
