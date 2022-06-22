package repository

import (
	"context"
	"fmt"

	"github.com/ibrahimker/latihan-register/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserIface interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

type UserRepo struct {
	postgrespool *pgxpool.Pool
}

func NewUserRepo(postgrespool *pgxpool.Pool) UserIface {
	return &UserRepo{postgrespool: postgrespool}
}

func (u *UserRepo) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	rows, err := u.postgrespool.Query(ctx, "insert into public.user values('')"
	if err != nil {
		fmt.Println("query row error", err)
	}
	defer rows.Close()

	users := []*entity.User{}
	for rows.Next() {
		var user entity.User
		if serr := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.CreatedAt, &user.UpdatedAt); serr != nil {
			fmt.Println("Scan error", serr)
		}
		users = append(users, &user)
	}
	return nil, nil
}
