package models

import (
	"strings"

	"gorm.io/gorm"
)

// Customer defines a customer who requests for a ride hail service
type Customer struct {
	BaseModelSoftDelete
	LastName   string `gorm:"not null;"`
	OtherNames *string
	Email      *string `gorm:"unique"`
	Phone      string  `gorm:"not null;unique"`
}

// BeforeSave hook is called before the data is persisted to db
func (customer *Customer) BeforeSave(tx *gorm.DB) (err error) {

	//remove 0 from phone and adds 233
	trimmedPhone := strings.TrimSpace(customer.Phone)
	ghCodeAdded := "233" + trimmedPhone[1:10]
	customer.Phone = ghCodeAdded

	return
}
