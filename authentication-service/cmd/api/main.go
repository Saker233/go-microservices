package main

import (
	"authentication/data"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
)

const webPort = "8000"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting Authentication Service.......")

	// conenct to DB
	conn := connectToDB()
	if conn == nil {
		log.Panic("Cant connect to postgres")
	}

	// setup config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	router := app.routes()

	err := router.Run(":" + webPort)
	if err != nil {
		log.Fatal(err)
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgress!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for 2 seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}
