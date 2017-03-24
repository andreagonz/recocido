package recocido

// Interfaces

type Solucion interface {
	ObtenVecino(solucion Solucion)
}

type Lote interface {
	CalculaLote(temperatura Temperatura, solucion Solucion)	
}

type Temperatura interface {
	TemperaturaInicial(solucion Solucion, temperatura Temperatura, porcentaje float64)
	PorcentajeAceptados(solucion Solucion, temperatura Temperatura)
}

type Recocido interface {
	AceptacionPorHumbrales(temperatura Temperatura, solucion Solucion)
}
