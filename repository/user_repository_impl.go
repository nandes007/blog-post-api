package repository

import (
	"context"
	"database/sql"
	"log"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	currentDate := helper.GetCurrentTime()
	SqlQuery := "INSERT INTO users(name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SqlQuery, user.Name, user.Email, user.Password, currentDate, currentDate).Scan(&id)
	//helper.PanicIfError(err)

	if err != nil {
		log.Fatal(err)
	}

	//id, err := result.LastInsertId()
	//helper.PanicIfError(err)

	if err != nil {
		log.Fatal(err)
	}

	user.Id = id
	return user
}

func (repository *UserRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SqlQuery := "SELECT id, name, email FROM users"
	rows, err := tx.QueryContext(ctx, SqlQuery)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}
