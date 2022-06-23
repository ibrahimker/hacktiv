package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/ibrahimker/latihan-register/entity"
	"github.com/ibrahimker/latihan-register/service"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepo struct {
	postgrespool *pgxpool.Pool
}

func NewUserRepo(postgrespool *pgxpool.Pool) service.UserRepository {
	return &UserRepo{postgrespool: postgrespool}
}

func (u *UserRepo) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	queryString := "insert into public.users " +
		"(username,email,password,age,created_at,updated_at)" +
		"values ($1,$2,$3,$4,$5,$5) returning id"
	rows, err := u.postgrespool.Query(ctx, queryString, user.Username, user.Email, user.Password, user.Age, time.Now())
	if err != nil {
		fmt.Println("query row error", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if serr := rows.Scan(&id); serr != nil {
			fmt.Println("Scan error", serr)
			return nil, serr
		}
		user.Id = id
	}
	return user, nil
}
