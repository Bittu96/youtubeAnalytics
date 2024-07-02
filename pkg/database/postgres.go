package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	conn     *sql.DB
	mux      *sync.Mutex
}

// global db client variable
var dbClient *DB

// create new db
func New(host string, port int, user, password, dbname string) {
	dbClient = &DB{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
		mux:      &sync.Mutex{},
	}
	fmt.Println(dbClient)
}

// connect to database
func (db *DB) Connect() (err error) {
	// connection string
	psqlConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.host,
		db.port,
		db.user,
		db.password,
		db.dbname)

	// open database
	db.conn, err = sql.Open("postgres", psqlConnectionString)
	if err != nil {
		fmt.Println("postgres connection failed!", err)
		return err
	}

	// check db
	if err = db.conn.Ping(); err != nil {
		log.Println("postgres ping failed!", err)
		return err
	}

	log.Println("postgres connection success!")
	return err
}

func GetClient() *sql.DB {
	if dbClient.conn == nil {
		dbClient.mux.Lock()
		if dbClient.conn == nil {
			for dbClient.Connect() != nil {
				time.Sleep(time.Second)
			}
		}
		dbClient.mux.Unlock()
	}

	return dbClient.conn
}

func CloseClient() {
	if dbClient.conn == nil {
		log.Println("postgres connection already closed!")
		return
	}

	if err := dbClient.conn.Close(); err != nil {
		log.Println("postgres connection close failed!")
	}
}
