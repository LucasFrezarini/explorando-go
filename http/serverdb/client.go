package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Usuario representa um usuário e seus respectivos atributos
type Usuario struct {
	ID   int64  `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler é utilizado para direcionar a operação a ser realizada de acordo com o método HTTP e a URL
// utilizada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/") // retira tudo até o prefixo e deixa o que está somente dps da barra
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == http.MethodGet && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == http.MethodGet:
		usuarioTodos(w, r)
	case r.Method == http.MethodPost:
		criarUsuario(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Página não encontrada :(")
	}
}

func getConexao() *sql.DB {
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3307)/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db := getConexao()
	defer db.Close()

	var u Usuario
	db.QueryRow("SELECT id, nome FROM usuarios WHERE id = ?", id).Scan(&u.ID, &u.Nome)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))

}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db := getConexao()
	defer db.Close()

	rows, _ := db.Query("SELECT id, nome FROM usuarios")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func criarUsuario(w http.ResponseWriter, r *http.Request) {
	var u Usuario
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)

	if err != nil {
		panic(err)
	}

	db := getConexao()
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO usuarios (nome) VALUES (?)")
	defer stmt.Close()

	result, _ := stmt.Exec(u.Nome)

	u.ID, _ = result.LastInsertId()

	response, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response))
}
