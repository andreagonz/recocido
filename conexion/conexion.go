package recocido

import (
	"github.com/andreagonz/recocido/implementacion"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// LeeCiudades lee las ciudades de la base de datos
func LeeCiudades(numCiudades int) []recocido.Ciudad {
	ciudades := make([]recocido.Ciudad, numCiudades)
	var i = 0
	db, err := sql.Open("sqlite3", "./db/ciudades")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from cities")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		
		var id int
		var name string
		var country string
		var population int
		var latitude float64
		var longitude float64
		
		err = rows.Scan(&id, &name, &country, &population, &latitude, &longitude)
		if err != nil {
			log.Fatal(err)
		}
		ciudad := recocido.Ciudad{
			id - 1,
			name,
			country,
			population,
			latitude,
			longitude}
		ciudades[i] = ciudad
		i++
	}
	
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return ciudades
}

// LeeConexiones lee las conexiones de las ciudades de la base de datos y
// regresa la matriz de adyacencias.
func LeeConexiones(numCiudades int) [][]float64 {
	sum := 0.0
	distancias := make([][]float64, numCiudades)
	for i := 0; i < numCiudades; i++ {
		distancias[i] = make([]float64, numCiudades)
	}
	
	db, err := sql.Open("sqlite3", "./db/ciudades")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from connections")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		
		var id1 int
		var id2 int
		var distancia float64
		
		err = rows.Scan(&id1, &id2, &distancia)
		if err != nil {
			log.Fatal(err)
		}
		distancias[id1 - 1][id2 - 1] = distancia
		distancias[id2 - 1][id1 - 1] = distancia
		sum += distancia

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return distancias
}
