package services

import (
	"context"

	"github.com/Bendomey/RideHail/account/internal/orm"
	"github.com/Bendomey/RideHail/account/internal/orm/models"
)

// AdminService inteface holds the user-databse transactions of this controller
type AdminService interface {
	CreateAdmin(ctx context.Context, name string, email string, password string, role string, createdBy *string, phone *string) (*models.Admin, error)
}

//ORM gets orm connection
type ORM struct {
	DB *orm.ORM
}

// NewAdminSvc exposed the ORM to the admin functions in the module
func NewAdminSvc(db *orm.ORM) AdminService {
	return &ORM{db}
}

// CreateAdmin creates an admin when invoked
func (orm *ORM) CreateAdmin(ctx context.Context, name string, email string, password string, role string, createdBy *string, phone *string) (*models.Admin, error) {
	_Admin := models.Admin{
		FullName:    name,
		Email:       email,
		Password:    password,
		Phone:       phone,
		Role:        role,
		CreatedByID: createdBy,
	}
	_Result := orm.DB.DB.Select("FullName", "Email", "Password", "Phone", "Role", "CreatedByID").Create(&_Admin)
	if _Result.Error != nil {
		return nil, _Result.Error
	}
	return &_Admin, nil
}
