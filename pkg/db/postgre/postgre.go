package postgre

import (
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST = "localhost"
	PORT = 5432
)

type Database struct {
	Conn *sql.DB
}

func Initialize(username, password, database string) *gorm.DB {
	address := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", HOST, username, password, database, PORT)
	db, err := gorm.Open(postgres.Open(address), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
		return nil
	}
	if db == nil {
		log.Printf("cannot connect to postgres")
		return nil
	}

	return db
}
