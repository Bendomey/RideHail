package services

import (
	"context"
	"errors"

	"github.com/Bendomey/RideHail/account/internal/orm"
	"github.com/Bendomey/RideHail/account/internal/orm/models"
	"github.com/Bendomey/RideHail/account/pkg/utils"
	"github.com/Bendomey/goutilities/pkg/signjwt"
	"github.com/Bendomey/goutilities/pkg/validatehash"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// AdminService inteface holds the user-databse transactions of this controller
type AdminService interface {
	CreateAdmin(ctx context.Context, name string, email string, password string, role string, createdBy *string, phone *string) (*models.Admin, error)
	LoginAdmin(ctx context.Context, email string, password string) (*LoginResult, error)
}

//ORM gets orm connection
type ORM struct {
	DB *orm.ORM
}

//LoginResult is the typing for returning login successful data to user
type LoginResult struct {
	Token string       `json:"token"`
	Admin models.Admin `json:"admin"`
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

// LoginAdmin checks if the email is having valid credentials and returns them a unique, secured token to help them get resources from app
func (orm *ORM) LoginAdmin(ctx context.Context, email string, password string) (*LoginResult, error) {
	var _Admin models.Admin

	//check if email is in db
	_Result := orm.DB.DB.Where("email = ?", email).Preload("CreatedBy").Find(&_Admin)
	if errors.Is(_Result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("AdminNotFound")
	}

	//since email in db, lets validate hash and then send back
	isSame := validatehash.ValidateCipher(password, _Admin.Password)
	if isSame == false {
		return nil, errors.New("PasswordIncorrect")
	}

	//sign token
	token, signTokenErrr := signjwt.SignJWT(jwt.MapClaims{
		"id":   _Admin.ID,
		"role": _Admin.Role,
	}, utils.MustGet("SECRET"))

	if signTokenErrr != nil {
		return nil, signTokenErrr
	}

	return &LoginResult{
		Token: token,
		Admin: _Admin,
	}, nil
}
