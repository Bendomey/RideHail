package models

import (
	"errors"
	"strings"

	"github.com/Bendomey/goutilities/pkg/hashpassword"
	"gorm.io/gorm"
)

// Admin defines a administrators who are to manage the app
type Admin struct {
	BaseModelSoftDelete
	FullName string `gorm:"not null;"`
	Email    string `gorm:"not null;unique_index:idx_email"`
	Password string `gorm:"not null;"`
	Phone    *string
	Role     string `gorm:"not null;"`
}

// BeforeSave hook is called before the data is persisted to db
func (admin *Admin) BeforeSave(tx *gorm.DB) (err error) {
	//remove 0 from phone and adds 233
	if admin.Phone != nil {
		trimmedPhone := strings.TrimSpace(*admin.Phone)
		ghCodeAdded := "233" + trimmedPhone[1:10]
		admin.Phone = &ghCodeAdded
	}

	//hashes password
	hashed, err := hashpassword.HashPassword(admin.Password)
	admin.Password = hashed
	if err != nil {
		err = errors.New("CannotHashPassword")
	}
	return
}
