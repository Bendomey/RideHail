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
	fullname := "Domey"
	phone := "0545526661"
	a, svcErr := adminSvc.UpdateAdmin(context.TODO(), "ed4b567e-c397-4435-82e5-090d8b4bc58c", &fullname, nil, &phone)
	if svcErr != nil {
		log.Error(svcErr)
	}

	log.NewLogger().Print(a)
	log.NewLogger().Print("Hello world")
}
