package repository

import (
	"context"
	"database/sql"
	"errors"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/helper/jwt"
	"nandes007/blog-post-rest-api/model/web/user"
	"time"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) Find(ctx context.Context, token string) (*user.UserResponse, error) {
	userId, err := jwt.ParseUserToken(token)
	var id int
	var name, email string
	var createdAt, updatedAt time.Time

	if err != nil {
		return nil, errors.New("invalid credential")
	}

	query := "SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1"
	err = r.db.QueryRowContext(ctx, query, userId).Scan(&id, &name, &email, &createdAt, &updatedAt)
	if err != nil {
		return nil, errors.New("invalid credential")
	}

	return &user.UserResponse{
		Id:        id,
		Name:      name,
		Email:     email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (r *userRepositoryImpl) GetAll(ctx context.Context) ([]*user.UserResponse, error) {
	SqlQuery := "SELECT id, name, email FROM users"
	rows, err := r.db.QueryContext(ctx, SqlQuery)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []*user.UserResponse
	for rows.Next() {
		user := &user.UserResponse{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
