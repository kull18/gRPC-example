package core

import (
	"database/sql"
	"log"
	"sync"
	_ "github.com/go-sql-driver/mysql"
)

var (
	conn *sql.DB
	once sync.Once
)

func GetConn() *sql.DB {
	once.Do(func ()  {
		dsn := "root:root@tcp(127.0.0.1:3306)/testdb"

		var err error
		conn, err = sql.Open("mysql",dsn)

		if err != nil {
			log.Fatal("Error to connect to database")
		}

		if err := conn.Ping(); err != nil {
			log.Fatal("Erro to make ping to database")
		}
	})

	return conn; 
}