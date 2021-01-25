package main

import (
	log "github.com/Bendomey/RideHail/account/internal/logger"
	"github.com/Bendomey/RideHail/account/internal/orm"
)

func main() {
	// creates a new ORM instance to send it to our server
	if _, err := orm.Factory(); err != nil {
		log.Panic(err)
	}

	log.NewLogger().Print("Hello world")
}
