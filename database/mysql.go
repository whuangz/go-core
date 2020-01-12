package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(dataSource string) {
	DB, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err.Error())
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	defer DB.Close()

	log.Println("Conencted To Mysql")
}
