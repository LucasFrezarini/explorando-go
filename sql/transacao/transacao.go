package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3307)/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO usuarios(id, nome) VALUES (?, ?)")

	stmt.Exec(2000, "Mike")
	stmt.Exec(2001, "Thor")

	_, err = stmt.Exec(1, "Lucao") // ID duplicado

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
