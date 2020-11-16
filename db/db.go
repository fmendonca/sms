package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Conexaodb() *sql.DB {

	//conexao := "user=postgres dbname=sms password=password host=localhost port=5432 sslmode=disable"
	//conexao := os.Getenv("DBURI")
	dbDriver := "mysql"
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := os.Getenv("DBHOST")
	dbName := os.Getenv("DBNAME")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbHost+"/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
