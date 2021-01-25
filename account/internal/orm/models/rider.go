package models

import (
	"errors"
	"strings"
	"time"

	"github.com/Bendomey/goutilities/pkg/hashpassword"
	"gorm.io/gorm"
)

// Rider defines a riders who are to accept trips from customers
type Rider struct {
	BaseModelSoftDelete
	LastName     string `gorm:"not null;"`
	OtherNames   string `gorm:"not null;"`
	Email        string `gorm:"not null;unique"`
	Phone        string `gorm:"not null;unique"`
	Password     string `gorm:"not null;"`
	AuthorisedAt *time.Time
	AuthorisedBy *Admin `gorm:"foreignKey:ID"`
	ProfilePhoto string `gorm:"not null;"`
	LicenseBack  string `gorm:"not null;"`
	LicenseFront string `gorm:"not null;"`
	Bike         string `gorm:"not null;"`
}

// BeforeSave hook is called before the data is persisted to db
func (rider *Rider) BeforeSave(tx *gorm.DB) (err error) {
	//remove 0 from phone and adds 233
	trimmedPhone := strings.TrimSpace(rider.Phone)
	ghCodeAdded := "233" + trimmedPhone[1:10]
	rider.Phone = ghCodeAdded

	//hashes password
	hashed, err := hashpassword.HashPassword(rider.Password)
	rider.Password = hashed
	if err != nil {
		err = errors.New("CannotHashRiderPassword")
	}
	return
}
