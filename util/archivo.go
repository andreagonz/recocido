package recocido

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
	"os"
)

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

func LeeArchivo(nom string) string {
	if _, err := os.Stat(nom); os.IsNotExist(err) {
		return ""
	}
	dat, err := ioutil.ReadFile(nom)
	check(err)
	return string(dat)
}

func CadenaARuta(ruta string) []int {
	r := strings.Replace(ruta, ", ", " ", -1)
	l := strings.Fields(r)
	res := make([]int, len(l))
	for i := 0; i < len(l); i++ {
		ind, err := strconv.Atoi(l[i])
		check(err)
		res[i] = ind - 1
	}
	return res
}

func CadenaParametros(seed int64, tLote int, p float64, ep float64, et float64, e float64, phi float64, c int) string {
	s := ""
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
