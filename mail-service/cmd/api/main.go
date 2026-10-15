package main

import (
	"log"
	"os"
	"strconv"
)

const webPort = "8000"

type Config struct {
	Mailer Mail
}

func main() {
	app := Config{
		Mailer: createMail(),
	}

	log.Println("Starting Mail Service on port:", webPort)

	router := app.routes()

	err := router.Run(":" + webPort)
	if err != nil {
		log.Fatal(err)
	}
}

func createMail() Mail {
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	m := Mail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        port,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORF"),
		Encryption:  os.Getenv("MAIL_ENCRYPTION"),
		FromName:    os.Getenv("FromName"),
		FromAddress: os.Getenv("FromAddress"),
	}

	return m
}
