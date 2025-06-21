package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	query := fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", name)
	fmt.Println("Executing query: ", query)

	db, err := sql.Open("mysql", "dbuser:dbpass@tcp(localhost:3306)/dbname")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Fprintf(w, "Query executed successfully")
}

func main() {
	http.HandleFunc("/users", handler)
	http.ListenAndServe(":8080", nil)
}
