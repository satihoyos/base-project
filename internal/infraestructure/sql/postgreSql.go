package sql

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type postgresConnection struct{}

func NewPostgresConnection() Connection {
	return &postgresConnection{}
}

func (psc *postgresConnection) Get() *sql.DB {
	//todo manejar parametros de conexion en la struct y concatenar
	//todo mejorar el manejo del error error
	dsn := "postgres://postgres:123qweasd@rest-test.cqhdjci6rbay.us-east-2.rds.amazonaws.com:5432/crud-test?sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
