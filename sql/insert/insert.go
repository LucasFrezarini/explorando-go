package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3307)/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO usuarios (nome) VALUES(?)")

	stmt.Exec("Lucas")
	stmt.Exec("João")
	res, _ := stmt.Exec("Maria")

	id, _ := res.LastInsertId()

	fmt.Println("LastInsertId:", id)

	rowsAffected, _ := res.RowsAffected()
	fmt.Println("Rows Affected:", rowsAffected)
}
