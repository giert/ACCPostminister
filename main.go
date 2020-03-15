package main

import (
	"ACCPostminister/bot"
	"log"
)

func main() {
	session, err := bot.Startup()
	if err != nil {
		log.Fatal("While setting up bot:", err)
	}

	err = bot.Run(session)
	if err != nil {
		log.Fatal("While running bot:", err)
	}

	err = bot.Shutdown(session)
	if err != nil {
		log.Fatal("While shutting down bot:", err)
	}
}
