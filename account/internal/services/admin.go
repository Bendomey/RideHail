package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Bendomey/RideHail/account/internal/orm"
	"github.com/Bendomey/RideHail/account/internal/orm/models"
	"github.com/Bendomey/RideHail/account/pkg/utils"
	"github.com/Bendomey/goutilities/pkg/generatecode"
	"github.com/Bendomey/goutilities/pkg/signjwt"
	"github.com/Bendomey/goutilities/pkg/validatehash"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// AdminService inteface holds the user-databse transactions of this controller
type AdminService interface {
	CreateAdmin(ctx context.Context, name string, email string, password string, role string, createdBy *string, phone *string) (*models.Admin, error)
	LoginAdmin(ctx context.Context, email string, password string) (*LoginResult, error)
	UpdateAdminRole(ctx context.Context, adminID string, role string) (bool, error)
	UpdateAdmin(ctx context.Context, adminID string, fullname *string, email *string, phone *string) (bool, error)
	UpdateAdminPassword(ctx context.Context, adminID string, oldPassword string, newPassword string) (bool, error)
	DeleteAdmin(ctx context.Context, adminID string) (bool, error)
}

//ORM gets orm connection
type ORM struct {
	DB  *orm.ORM
	rdb *redis.Client
}

//LoginResult is the typing for returning login successful data to user
type LoginResult struct {
	Token string       `json:"token"`
	Admin models.Admin `json:"admin"`
}

// NewAdminSvc exposed the ORM to the admin functions in the module
func NewAdminSvc(db *orm.ORM, rdb *redis.Client) AdminService {
	return &ORM{db, rdb}
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
	_Result := orm.DB.DB.Joins("CreatedBy").First(&_Admin, "admins.email = ?", email)
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

// UpdateAdminRole updates role of an admin
func (orm *ORM) UpdateAdminRole(ctx context.Context, adminID string, role string) (bool, error) {
	var _Admin models.Admin

	err := orm.DB.DB.First(&_Admin, "id = ?", adminID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errors.New("AdminNotFound")
	}
	_Admin.Role = role
	orm.DB.DB.Save(&_Admin)
	return true, nil

}

// UpdateAdmin updates data of an admin
func (orm *ORM) UpdateAdmin(ctx context.Context, adminID string, fullname *string, email *string, phone *string) (bool, error) {
	var _Admin models.Admin

	err := orm.DB.DB.First(&_Admin, "id = ?", adminID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errors.New("AdminNotFound")
	}

	if fullname != nil {
		_Admin.FullName = *fullname
	}
	if email != nil {
		_Admin.Email = *email
	}
	if phone != nil {
		_Admin.Phone = phone
	}
	orm.DB.DB.Save(&_Admin)
	return true, nil

}

// UpdateAdminPassword updates password of an admin
func (orm *ORM) UpdateAdminPassword(ctx context.Context, adminID string, oldPassword string, newPassword string) (bool, error) {
	var _Admin models.Admin

	err := orm.DB.DB.First(&_Admin, "id = ?", adminID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errors.New("AdminNotFound")
	}

	isSame := validatehash.ValidateCipher(oldPassword, _Admin.Password)
	if isSame == false {
		return false, errors.New("OldPasswordIncorrect")
	}

	_Admin.Password = newPassword
	orm.DB.DB.Save(&_Admin)
	return true, nil

}

// DeleteAdmin deletes an admin
func (orm *ORM) DeleteAdmin(ctx context.Context, adminID string) (bool, error) {
	var _Admin models.Admin

	err := orm.DB.DB.Delete(&_Admin, "id = ?", adminID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errors.New("AdminNotFound")
	}
	return true, nil
}

/**
* Forgot Password Fucntionality
* 1. Request with Email
* 2. Send Code And save code in redis
* 3. Compare code And send response
* 4. Enter new password and remove from redis
 */

// ForgotPasswordRequest is to start the process
func (orm *ORM) ForgotPasswordRequest(ctx context.Context, email string) (*models.Admin, error) {
	var _Admin models.Admin

	// check if admin exists
	err := orm.DB.DB.First(&_Admin, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("AdminNotFound")
	}

	//generate code
	code := generatecode.GenerateCode(6)

	// send code to email

	//save in redis and expire in an hours time
	redisErr := orm.rdb.Set(ctx, fmt.Sprintf("%s", _Admin.ID), code, 1*time.Hour).Err()
	if redisErr != nil {
		return nil, redisErr
	}

	return &_Admin, nil
}

//ResendCode helps to resend a new code
func (orm *ORM) ResendCode(ctx context.Context, adminID string) (*models.Admin, error) {
	var _Admin models.Admin

	// check if admin exists
	err := orm.DB.DB.First(&_Admin, "id = ?", adminID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("AdminNotFound")
	}

	//generate code
	code := generatecode.GenerateCode(6)

	// send code to email

	//save in redis and expire in an hours time
	redisErr := orm.rdb.Set(ctx, fmt.Sprintf("%s", _Admin.ID), code, 1*time.Hour).Err()
	if redisErr != nil {
		return nil, redisErr
	}

	return &_Admin, nil
}

// CompareAdminCodes compares the admin code sent by user
func (orm *ORM) CompareAdminCodes(ctx context.Context, adminID string, code string) (bool, error) {
	//check in redis to see if its the same and not expired
	value, err := orm.rdb.Get(ctx, fmt.Sprintf("%s", adminID)).Result()
	if err == redis.Nil {
		return false, errors.New("CodeHasExpired")
	} else if err != nil {
		return false, err
	}

	if value != code {
		return false, errors.New("CodeIncorrect")
	}
	return true, nil
}

// ResetPassword updates the admins new password
func (orm *ORM) ResetPassword(ctx context.Context, adminID string, password string) (bool, error) {
	var _Admin models.Admin

	err := orm.DB.DB.First(&_Admin, "id = ?", adminID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errors.New("AdminNotFound")
	}

	//update password
	_Admin.Password = password
	orm.DB.DB.Save(&_Admin)

	//invalidate the redis data pertaining to this admin

	return true, nil
}
