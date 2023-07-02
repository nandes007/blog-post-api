package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/helper/hash"
	"nandes007/blog-post-rest-api/helper/jwt"
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/user"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	currentDate := helper.GetCurrentTime()
	passwordHashed := hash.PasswordHash(user.Password)
	SqlQuery := "INSERT INTO users(name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SqlQuery, user.Name, user.Email, passwordHashed, currentDate, currentDate).Scan(&id)

	if err != nil {
		helper.PanicIfError(err)
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

func (repository *UserRepositoryImpl) Login(ctx context.Context, db *sql.DB, request user.LoginRequest) (string, error) {
	stmt, err := db.Prepare("SELECT id, email, password FROM users WHERE email = $1 LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(request.Email)

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

	if err != nil {
		helper.PanicIfError(err)
	}

	return tokenString, nil
}

func (repository *UserRepositoryImpl) Find(ctx context.Context, db *sql.DB, token string) (domain.User, error) {
	//TODO implement me
	userId, err := jwt.ParseUserToken(token)
	defer db.Close()

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
