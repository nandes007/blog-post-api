package app

import (
	"database/sql"
	"nandes007/blog-post-rest-api/helper"
	"time"
)

func NewDB() *sql.DB {
	//connStr := "postgres://postgres:postgre@localhost/blog_post?sslmode=disable"
	// driver := os.Getenv("DB_DRIVER")
	// username := os.Getenv("DB_USERNAME")
	// password := os.Getenv("DB_PASSWORD")
	// host := os.Getenv("DB_HOST")
	// dbName := os.Getenv("DB_NAME")

	connStr := "host=localhost port=5433 user=postgres password=password dbname=blog_post sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
