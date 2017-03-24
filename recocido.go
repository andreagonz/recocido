package main

import (
	"fmt"
	"github.com/andreagonz/recocido/conexion"
)

func main() {
	ciudades := recocido.LeeCiudades(277)
	for i := 0; i < len(ciudades); i++ {
		fmt.Println(ciudades[i].Id, ciudades[i].Nombre)
	}
	conexiones := recocido.LeeConexiones(277)
	fmt.Println(conexiones[260][265])
}
