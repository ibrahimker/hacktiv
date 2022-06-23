package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ibrahimker/latihan-register/entity"
	"github.com/ibrahimker/latihan-register/service"
	mock_service "github.com/ibrahimker/latihan-register/test/mock/service"
	"github.com/stretchr/testify/require"
)

func TestNewUserSvc(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("Test Initiate New User Service", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		require.NotNil(t, userService)
	})
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty Username", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		res, err := userService.Register(context.Background(), &entity.User{
			Username: "",
		})
		require.Error(t, err)
		require.Equal(t, errors.New("username cannot be empty"), err)
		require.Nil(t, res)
	})

	t.Run("database down", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		user := &entity.User{
			Username: "abc 123",
			Password: "password",
			Email:    "email@email.com",
			Age:      25,
		}
		mockUserRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil, errors.New("db down"))

		res, err := userService.Register(context.Background(), user)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, "db down", err.Error())
	})

	t.Run("Successfullly insert to db", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		user := &entity.User{
			Username: "abc 123",
			Password: "password",
			Email:    "email@email.com",
			Age:      25,
		}
		userRes := &entity.User{
			Id:       1,
			Username: "abc 123",
			Password: "password",
			Email:    "email@email.com",
			Age:      25,
		}
		mockUserRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(userRes, nil)

		res, err := userService.Register(context.Background(), user)
		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, 1, res.Id)
	})
}
