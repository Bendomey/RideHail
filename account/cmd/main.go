package main

import (
	log "github.com/Bendomey/RideHail/account/internal/logger"
	"github.com/Bendomey/RideHail/account/internal/orm"
	"github.com/Bendomey/RideHail/account/internal/services"
)

func main() {
	// creates a new ORM instance to send it to our server
	orm, err := orm.Factory()
	if err != nil {
		log.Panic(err)
	}

	//call service
	services.NewAdminSvc(orm)

	log.NewLogger().Print("Hello world")
}
