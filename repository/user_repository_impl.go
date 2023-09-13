package repository

import (
	"context"
	"database/sql"
	"errors"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/helper/jwt"
	"nandes007/blog-post-rest-api/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
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

func (repository *UserRepositoryImpl) Find(ctx context.Context, db *sql.DB, token string) (domain.User, error) {
	userId, err := jwt.ParseUserToken(token)

	if err != nil {
		return domain.User{}, errors.New("invalid credential")
	}

	query := "SELECT id, name, email FROM users WHERE id = $1"
	var user domain.User
	err = db.QueryRowContext(ctx, query, userId).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return domain.User{}, errors.New("invalid credential")
	}

	return user, nil
}
