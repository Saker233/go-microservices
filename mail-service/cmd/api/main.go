package main

import "log"

const webPort = "8000"

type Config struct {
}

func main() {
	app := Config{}

	log.Println("Starting Mail Service on port:", webPort)

	router := app.routes()

	err := router.Run(":" + webPort)
	if err != nil {
		log.Fatal(err)
	}
}
