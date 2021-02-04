package postgre

import (
	"database/sql"
	"fmt"
	"github.com/go-pg/pg"
	"log"
)

const (
	HOST = "localhost"
	PORT = 5432
)

type Database struct {
	Conn *sql.DB
}

func Initialize(username, password, database string) *pg.DB {
	address := fmt.Sprintf("%s:%d", HOST, PORT)
	db := pg.Connect(&pg.Options{
		Addr:     address,
		User:     "postgres",
		Password: password,
		Database: database,
	})

	if db == nil {
		log.Printf("cannot connect to postgres")
	}
	return db
}
