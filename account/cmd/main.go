package main

import (
	"context"

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
	adminSvc := services.NewAdminSvc(orm)
	loginRes, err := adminSvc.LoginAdmin(context.TODO(), "domeybenjami@gmail.com", "akankobateng1")
	if err != nil {
		log.Error(err)
	}

	log.NewLogger().Print(loginRes.Admin.CreatedBy)
	log.NewLogger().Print("Hello world")
}
