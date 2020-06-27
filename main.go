package main

import (
	"ACCPostminister/bot"
	"log"
)

func main() {
	session, err := bot.Startup()
	if err != nil {
		log.Fatal("while setting up bot:", err)
	}

	err = bot.Run(session)
	if err != nil {
		log.Fatal("while running bot:", err)
	}

	err = bot.Shutdown(session)
	if err != nil {
		log.Fatal("while shutting down bot:", err)
	}
}
