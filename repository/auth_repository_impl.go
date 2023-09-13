package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/helper/hash"
	"nandes007/blog-post-rest-api/helper/jwt"
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/auth"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (repository AuthRepositoryImpl) Login(ctx context.Context, db *sql.DB, request auth.LoginRequest) (string, error) {
	stmt, err := db.Prepare("SELECT id, email, password FROM users WHERE email = $1 LIMIT 1")
	helper.PanicIfError(err)

	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, request.Email)

	var id int
	var (
		email    string
		password string
	)

	err = row.Scan(&id, &email, &password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("credential mismatch")
		} else {
			helper.PanicIfError(err)
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(request.Password))
	if err != nil {
		return "", errors.New("credential mismatch")
	}

	tokenString, err := jwt.CreateToken(id)
	helper.PanicIfError(err)

	return tokenString, nil
}

func (repository AuthRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	currentDate := helper.GetCurrentTime()
	passwordHashed := hash.PasswordHash(user.Password)
	stmt := "INSERT INTO users(name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	var id int
	err := tx.QueryRowContext(ctx, stmt, user.Name, user.Email, passwordHashed, currentDate, currentDate).Scan(&id)
	helper.PanicIfError(err)

	user.Id = id
	return user, nil
}
