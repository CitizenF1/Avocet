package config

import (
	"avocet/internal/db"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	OdooLogin    string
	OdooPassword string
	OdooDatabase string
	OdooUrl      string
)

func init() {
	err := godotenv.Overload(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	OdooLogin = os.Getenv("ODOOLOGIN")
	OdooPassword = os.Getenv("ODOOPASSWORD")
	OdooDatabase = os.Getenv("ODOODATABASE")
	OdooUrl = os.Getenv("ODOOURL")

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println(err)
	}
	db.Server = os.Getenv("SERVER")
	db.Password = os.Getenv("PASSWORD")
	db.Port = port
	db.User = os.Getenv("USER")
	db.Database = os.Getenv("DATABASE")
}
