package db

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	Server   string
	Port     int
	User     string
	Password string
	Database string
	Conn     *sql.DB
)

func DatabaseConnection() *sql.DB {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		Server, User, Password, Port, Database)

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	return conn
}
