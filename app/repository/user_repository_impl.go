package repository

import (
	"database/sql"
	"errors"
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

func (r *userRepositoryImpl) GetAll() ([]*user.UserResponse, error) {
	SqlQuery := "SELECT id, name, email, created_at, updated_at FROM users"
	rows, err := r.db.Query(SqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user.UserResponse
	for rows.Next() {
		user := &user.UserResponse{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepositoryImpl) GetByID(id int) (*user.UserResponse, error) {
	var name, email string
	var createdAt, updatedAt time.Time

	query := "SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&id, &name, &email, &createdAt, &updatedAt)
	if err != nil {
		return nil, errors.New("invalid credential")
	}

	return &user.UserResponse{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
