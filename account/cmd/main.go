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
	a, svcErr := adminSvc.UpdateAdminPassword(context.TODO(), "bc00463c-c0ff-4560-a16e-ae222423c397", "domeybenjamin", "akankobateng1")
	if svcErr != nil {
		log.Error(svcErr)
	}

	log.NewLogger().Print(a)
	log.NewLogger().Print("Hello world")
}
