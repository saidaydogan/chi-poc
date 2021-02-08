package postgre

import (
	"database/sql"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
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

type User struct {
	Id     int64
	Name   string
	Emails []string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

type Story struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *User
}

func (s Story) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}

// createSchema creates database schema for User and Story models.
func CreateSchema(db *pg.DB) error {
	var err error

	models := []interface{}{
		(*User)(nil),
		(*Story)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}

	user1 := &User{
		Name:   "admin",
		Emails: []string{"admin1@admin", "admin2@admin"},
	}
	_, err = db.Model(user1).Insert()
	if err != nil {
		panic(err)
	}

	_, err = db.Model(&User{
		Name:   "root",
		Emails: []string{"root1@root", "root2@root"},
	}).Insert()
	if err != nil {
		panic(err)
	}

	story1 := &Story{
		Title:    "Cool story",
		AuthorId: user1.Id,
	}
	_, err = db.Model(story1).Insert()
	if err != nil {
		panic(err)
	}

	// Select user by primary key.
	user := &User{Id: user1.Id}
	err = db.Model(user).WherePK().Select()
	if err != nil {
		panic(err)
	}

	story := new(Story)
	err = db.Model(story).
		//Relation("Author").
		Where("story.id = ?", story1.Id).
		Select()
	if err != nil {
		panic(err)
	}

	return nil
}
