package recocido

import (
	"github.com/andreagonz/recocido/implementacion"
	_"github.com/mattn/go-sqlite3"
	"log"
	"database/sql"
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
			id - 1, name,
			country, population,
			latitude, longitude}
		ciudades[i] = ciudad
		i++
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return ciudades
}

func LeeConexiones(numCiudades int) [][]float64 {
	conexiones := make([][]float64, numCiudades)
	for i := 0; i < numCiudades; i++ {
		conexiones[i] = make([]float64, numCiudades)
	}
	
	db, err := sql.Open("sqlite3", "./db/ciudades")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id_city_1, id_city_2, distance from connections")
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
		conexiones[id1 - 1][id2 - 1] = distancia

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return conexiones

}
