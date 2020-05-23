package dao

import (
	"database/sql"
	"fmt"
	"log"

	// !
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "teste123"
	dbname   = "wow_api_data"
)

var db *sql.DB

func init() {
	var err error
	var connectionStr string
	connectionStr = getConnectionString()
	if err != nil {
		log.Fatalf("Could not get db connection string. Erro: %s", err)
	}
	db, err = sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatalf("Could not get db connection. Erro: %s", err)
	}
}

func getConnectionString() string {
	var connString string
	connString = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return connString
}
