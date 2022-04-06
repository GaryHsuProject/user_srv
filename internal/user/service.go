package user

import (
	"shop/dto"
	"shop/models"
	"shop/pkg/bcrypt"
	helper "shop/pkg/error"
	"shop/pkg/jwt"
	"time"
)

type Service interface {
	Insert(*dto.InsertUser) error
	Login(*dto.FindOneByUser) (*dto.UserLogin, error)
}

type UserService struct {
	UserRepo Repository
}

func NewUserService(userRepo Repository) Service {
	return &UserService{UserRepo: userRepo}
}

func (u *UserService) Insert(req *dto.InsertUser) error {
	has, err := u.UserRepo.CheckExistByUsername(&models.User{Username: req.Username})
	if err != nil {
		return helper.NewCustomError(500, err, "Check user by username error", req)
	}
	if has {
		return helper.NewCustomError(409, nil, "User already Exist", req)
	}
	password, err := bcrypt.HashAndSalt(req.Password)
	if err != nil {
		return helper.NewCustomError(500, err, "Gen user md5 password error", req)
	}
	now := time.Now()
	user := &models.User{
		Username:       req.Username,
		Password:       password,
		Status:         req.Status,
		RoleId:         req.RoleId,
		EmployeeID:     0,
		CreateTime:     &now,
		LastUpdateTime: &now,
	}
	if err := u.UserRepo.Insert(user); err != nil {
		return helper.NewCustomError(500, err, "Create user failed", req)
	}
	return nil
}
func (u *UserService) Login(req *dto.FindOneByUser) (*dto.UserLogin, error) {
	user, err := u.UserRepo.FindOne(&models.User{Username: req.Username})
	if err != nil {
		return nil, helper.NewCustomError(500, err, "User login error", req)
	}
	if user == nil {
		return nil, helper.NewCustomError(400, err, "Find user not found", req)
	}
	if !bcrypt.ComparePasswords(user.Password, req.Password) && user.Status == 0 {
		return nil, helper.NewCustomError(403, err, "User Login Failed", req)
	}
	token := jwt.GenJWTTOken()
	if token == "" {
		return nil, helper.NewCustomError(400, err, "User Login JWT TOKEN Failed", req)
	}
	return &dto.UserLogin{Token: token}, nil
}
