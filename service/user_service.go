package service

import (
	"context"
	"errors"

	"github.com/ibrahimker/latihan-register/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, user *entity.User) (*entity.User, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

type UserSvc struct {
	userRepo UserRepository
}

func NewUserSvc(userRepo UserRepository) UserService {
	return &UserSvc{
		userRepo: userRepo,
	}
}

func (u *UserSvc) Register(ctx context.Context, user *entity.User) (*entity.User, error) {
	// validate user data
	if err := validateUser(user); err != nil {
		return nil, err
	}

	// encrypt password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	// create user in database
	return u.userRepo.CreateUser(ctx, user)
}

func validateUser(user *entity.User) error {
	// validasi field field user
	if user.Username == "" {
		return errors.New("username cannot be empty")
	}
	if user.Email == "" {
		return errors.New("email cannot be empty")
	}
	if len(user.Password) < 6 {
		return errors.New("password must be minimum 6 characters")
	}
	if user.Age < 8 {
		return errors.New("age must be greater than 8")
	}
	return nil
}
