package main

import (
	"fmt"
	"github.com/andreagonz/recocido/lote"
	_ "github.com/mattn/go-sqlite3"
	"log"
	_ "os"
	"database/sql"
)

func main() {
	fmt.Printf(recocido.Reverse("lolah"))
	
	db, err := sql.Open("sqlite3", "./db/ciudades")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, name from cities")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
