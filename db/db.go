package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "host=localhost port=5433 user=postgres dbname=fitness sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := DB.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='public'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// log.Println("--- LISTA TABELA ---")
	// for rows.Next() {
	// 	var name string
	// 	rows.Scan(&name)
	// 	log.Println("Pronađena tabela:", name)
	// }
	// log.Println("--------------------")

	var dbname string
	DB.QueryRow("SELECT current_database()").Scan(&dbname)
	log.Println("Connected to DB:", dbname)
}
