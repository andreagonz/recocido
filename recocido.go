package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"io/ioutil"
	con "github.com/andreagonz/recocido/conexion"
	heu "github.com/andreagonz/recocido/heuristica"
	imp "github.com/andreagonz/recocido/implementacion"
)

func Mapa(ruta string, ciudades *[]imp.Ciudad) string {
	r := strings.Replace(ruta, ", ", " ", -1)
	l := strings.Fields(r)
	s := "<!DOCTYPE html>\n<html>\n<head>\n<meta name='viewport'content='initial-scale=1.0, user-scalable=no'>\n<meta charset='utf-8'>\n<title>Mapa</title>\n<style>\n#map {\nheight: 100%;\n}\nhtml, body {\nheight: 100%;\nmargin: 0;\npadding: 0;\n}\n</style>\n</head>\n<body>\n<div id='map'></div>\n<script>\nfunction initMap() {\n var map = new google.maps.Map(document.getElementById('map'), {\nzoom: 3,\ncenter: {lat: 0, lng: -180},\nmapTypeId: 'terrain'\n});\nvar flightPlanCoordinates = [\n"
	for i := 0; i < len(l); i++ {
		ind, err := strconv.Atoi(l[i])
		check(err)
		s += "{lat: " + strconv.FormatFloat((*ciudades)[ind - 1].Latitud, 'f', -1, 64) + ", lng: " + strconv.FormatFloat((*ciudades)[ind - 1].Longitud, 'f', -1, 64) + "}"
		if i < len(l) - 1 {
			s += ",\n"
		}
	}
	s += "      ];\n      var flightPath = new google.maps.Polyline({\n      path: flightPlanCoordinates,\n      geodesic: true,\n      strokeColor: '#FF0000',\n      strokeOpacity: 1.0,\n      strokeWeight: 2\n      });      \n      flightPath.setMap(map);\n      }\n    </script>\n    <script async defer\n            src='https://maps.googleapis.com/maps/api/js?key=AIzaSyDdBAKYa4kQqUStHeV39ngfUVZwRAl84bk&callback=initMap'>\n    </script>\n  </body>\n</html>"
	return s
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func EscribeArchivo(s string, nombre string) {
	d1 := []byte(s)
	err := ioutil.WriteFile(nombre, d1, 0644)
	check(err)
}

func CadenaParametros(seed int64, tLote int, p float64, ep float64, et float64, e float64, phi float64, c int) string {
	s := "Parámetros\n"
	s += fmt.Sprintf("Semilla: %d\n", seed)
	s += fmt.Sprintf("Tamaño lote: %d\n", tLote)
	s += fmt.Sprintf("p: %f\n", p)
	s += fmt.Sprintf("ep: %f\n", ep)
	s += fmt.Sprintf("et: %f\n", et)
	s += fmt.Sprintf("e: %f\n", e)
	s += fmt.Sprintf("phi: %f\n", phi)
	s += fmt.Sprintf("C: %d\n", c)
	return s
}

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

/*
//Carlos Main
func main() {
	numCiudades := 277
	ciudades := con.LeeCiudades(numCiudades)
	s := "213, 90, 185, 227, 159, 133, 147, 250, 233, 276, 66, 254, 143, 176, 209, 58, 146, 22, 23, 84, 266, 31, 158, 95, 144, 188, 1, 9, 216, 122, 248, 62, 5, 274, 80, 224, 29, 203, 235, 52, 215, 92, 48, 12, 179, 241, 186, 244, 221, 150, 107, 232, 75, 198, 220, 70, 190, 16, 238, 191, 264, 86, 135, 166, 119, 30, 94, 101, 194, 39, 160, 167, 117, 178, 65, 56, 200, 207"
	EscribeArchivo(Mapa(s, &ciudades), "carlos.html")
}
*/

func main() {

	seed := int64(1)
	numCiudades := 277
	// tProblema := 50
	tLote := 1000
	p := 0.9
	ep := 0.001
	et := 0.001
	e := 0.001
	phi := 0.9
	c := 5
	direc := "archivero/4/"
	
	ciudades := con.LeeCiudades(numCiudades)
	distancias := con.LeeConexiones(numCiudades)

	r := rand.New(rand.NewSource(seed))
//	problema := ProblemaAleatorio(tProblema, &ciudades, &distancias, r)


	problema := []int{1, 5, 9, 12, 16, 22, 23, 29, 30, 31, 39, 48, 52, 56, 58, 62, 65, 66, 70, 75, 80, 84, 86, 90, 92, 94, 95, 101, 107, 117, 119, 122, 133, 135, 143, 144, 146, 147, 150, 158, 159, 160, 166, 167, 176, 178, 179, 185, 186, 188, 190, 191, 194, 198, 200, 203, 207, 209, 213, 215, 216, 220, 221, 224, 227, 232, 233, 235, 238, 241, 244, 248, 250, 254, 264, 266, 274, 276}
	
	for i := 0; i < len(problema); i++ {
		problema[i] = problema[i] - 1
	}

	
	//problema := []int{25,36,13,6,0,26,30,1,32}
	
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
	
	EscribeArchivo(CadenaParametros(seed, tLote, p, ep, et, e, phi, c), direc + "parametros.txt")
	
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
		fmt.Println("Factible:" + strconv.FormatBool(mejor.EsFactible()))
		fmt.Println("Costo: " + strconv.FormatFloat(mejor.ObtenFun(), 'f', -1, 64))
		fmt.Println("Distancia: " + strconv.FormatFloat(mejor.ObtenFunObj(), 'f', -1, 64))
		if mejor.EsFactible() {
			EscribeArchivo(mejor.Str(), fmt.Sprintf(direc + "%02d.tsp", i))
			EscribeArchivo(fmt.Sprintf("Costo: %f", mejor.ObtenFun()), fmt.Sprintf(direc + "%02d.fun", i))
		}
	}

	EscribeArchivo(fmt.Sprintf("%02d.tsp", mejorGI), direc + "mejor.txt")
	EscribeArchivo(Mapa(mejorG.Str(), &ciudades), direc + "mejor.html")
	
	fmt.Println("Promedio: " + strconv.FormatFloat(prom, 'f', -1, 64))
	//fmt.Println(mejorG.Str())
	fmt.Println("Factible:" + strconv.FormatBool(mejorG.EsFactible()))
	fmt.Println("Costo: " + strconv.FormatFloat(mejorG.ObtenFun(), 'f', -1, 64))
	fmt.Println("Distancia: " + strconv.FormatFloat(mejorG.ObtenFunObj(), 'f', -1, 64))
	fmt.Println()
	
}
