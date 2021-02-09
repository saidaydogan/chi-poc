package postgre

import (
	"database/sql"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

const (
	HOST = "localhost"
	PORT = 5432
)

type Database struct {
	Conn *sql.DB
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	query, _ := q.FormattedQuery()

	log.Debug().Msg(string(query))
	return nil
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
		return nil
	}

	db.AddQueryHook(dbLogger{})
	return db
}
