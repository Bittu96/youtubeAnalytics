package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DBClient *sql.DB

type database struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func New(host string, port int, user, password, dbname string) database {
	return database{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
	}
}

func (db database) Connect() (*sql.DB, error) {
	_, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	fmt.Println("postgres creds", db)
	// if DBClient != nil {
	// 	// db.Connect()
	// 	// todo
	// 	// add nil test
	// 	// add ping test
	// 	// add auto retries
	// }

	// connection string
	psqlConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.host,
		db.port,
		db.user,
		db.password,
		db.dbname)
	fmt.Println("psqlConnectionString:", psqlConnectionString)

	var err error
	// open database
	DBClient, err = sql.Open("postgres", psqlConnectionString)
	if err != nil {
		fmt.Println("postgres connection failed!", err)
		log.Fatal("postgres connection failed!", err)
		return DBClient, err
	}

	// check db
	if err = DBClient.Ping(); err != nil {
		fmt.Println("postgres connection failed!", err)
		return DBClient, err
	}

	fmt.Println("postgres connection success!")
	return DBClient, err
}
