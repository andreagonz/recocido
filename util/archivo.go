package recocido

import (
	"io/ioutil"
	"fmt"
)

func EscribeArchivo(s string, nombre string) {
	d1 := []byte(s)
	err := ioutil.WriteFile(nombre, d1, 0644)
	if err != nil {
		panic(err)
	}
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
