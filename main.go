package main

import (
	"log"
)

func main() {
	session, err := startup()
	if err != nil {
		log.Fatal("While setting up bot:", err)
	}

	err = run(session)
	if err != nil {
		log.Fatal("While running bot:", err)
	}

	err = shutdown(session)
	if err != nil {
		log.Fatal("While shutting down bot:", err)
	}
}
