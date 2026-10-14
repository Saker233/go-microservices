package main

import (
	"authentication/data"
	"database/sql"
)

const webPort = "8000"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

}
