package controllers

import (
	"MireaPR4/database/models"
	"MireaPR4/database/repositories"
	"MireaPR4/http/jwt"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type RegisterController interface {
	CreateUser(
		username string,
		password string,
		email string,
		roleID int,
	) (*models.User, error)
	AuthenticateUser(username string, password string) (*string, error)
}

type registerController struct {
	userRepo repositories.UserRepository
}

func NewRegisterController(UserRepo repositories.UserRepository) RegisterController {
	return &registerController{userRepo: UserRepo}
}

func (ac *registerController) CreateUser(
	username string,
	password string,
	email string,
	roleID int,
) (*models.User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil || hashedPassword == nil {
		return nil, err
	}

	user := models.User{
		Username: username,
		Password: *hashedPassword,
		Email:    email,
		RoleID:   roleID,
		StatusID: 1,
	}

	if userOld, e := ac.userRepo.GetByUsername(username); e == nil {
		if userOld != nil {
			return nil, errors.New("user already exist")
		}
	}

	if err := ac.userRepo.Create(&user); err != nil {
		return nil, errors.New("internal server error")
	}

	return &user, nil
}

func (ac *registerController) AuthenticateUser(username string, password string) (*string, error) {
	var user *models.User

	user, err := ac.userRepo.GetByUsername(username)
	if err != nil || user == nil {
		return nil, err
	}

	if !checkPassword(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := jwt.GenerateToken(user.ID, user.Role.Name)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return &token, nil
}

func hashPassword(password string) (*string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	strHash := string(hash)
	return &strHash, nil
}

func checkPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
