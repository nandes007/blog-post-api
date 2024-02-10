package repository

import (
	"context"
	"database/sql"
	"errors"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/helper/hash"
	"nandes007/blog-post-rest-api/helper/jwt"
	"nandes007/blog-post-rest-api/model/web/auth"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type authRepositoryImpl struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepositoryImpl{
		db: db,
	}
}

func (r *authRepositoryImpl) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	stmt, err := r.db.Prepare("SELECT id, email, password FROM users WHERE email = $1 LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, req.Email)

	var id int
	var (
		email    string
		password string
	)

	err = row.Scan(&id, &email, &password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("credential mismatch")
		} else {
			return nil, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("credential mismatch")
	}

	tokenString, err := jwt.CreateToken(id)
	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		Token: tokenString,
	}, nil
}

func (r *authRepositoryImpl) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	var id int
	passwordHashed := hash.PasswordHash(req.Password)

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)
	stmt := "INSERT INTO users(name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err = tx.QueryRowContext(ctx, stmt, req.Name, req.Email, passwordHashed, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{
		Id:        id,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
