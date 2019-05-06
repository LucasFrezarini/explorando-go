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

	stmt, _ := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")
	stmt.Exec("Xeroque Rolmes", 1)
	stmt.Exec("Inbonha", 2)

	stmt2, _ := db.Prepare("DELETE FROM usuarios WHERE ID = ?")
	stmt2.Exec(3)
}
