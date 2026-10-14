package main

import (
	"authentication/data"
	"database/sql"
	"log"
)

const webPort = "8000"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting Authentication Service.......")

	// conenct to DB

	// setup config
	app := Config{}

	router := app.routes()

	err := router.Run(":" + webPort)
	if err != nil {
		log.Fatal(err)
	}

}
