package recocido

type Ciudad struct {
	Id int
	Nombre string
	Pais string
	Poblacion int
	Latitud float64
	Longitud float64
}

type Camino struct {
	Ciudades []Ciudad
}

type Conexiones struct {
	Conexiones []int
	Distancias []float64
}
