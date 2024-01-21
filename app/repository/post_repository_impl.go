package repository

import (
	"context"
	"database/sql"
	"log"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/post"
)

type PostRepositoryImpl struct {
}

func NewPostRepository() PostRepository {
	return &PostRepositoryImpl{}
}

func (repository PostRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User, post domain.Post) domain.Post {
	//TODO implement me
	currentDate := helper.GetCurrentTime()
	sqlQuery := "INSERT INTO posts(author_id, title, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, sqlQuery, user.Id, post.Title, post.Content, currentDate, currentDate).Scan(&id)

	if err != nil {
		helper.PanicIfError(err)
	}

	post.Id = id
	return post
}

func (repository PostRepositoryImpl) GetAll(ctx context.Context, db *sql.DB) []domain.Post {
	sqlQuery := "SELECT id, author_id, title, content, created_at FROM posts"
	rows, err := db.QueryContext(ctx, sqlQuery)

	if err != nil {
		helper.PanicIfError(err)
	}

	defer rows.Close()

	var posts []domain.Post

	for rows.Next() {
		post := domain.Post{}
		if err := rows.Scan(&post.Id, &post.AuthorId, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}

	rerr := rows.Close()

	if rerr != nil {
		log.Fatal(rerr)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return posts
}

func (repository PostRepositoryImpl) Find(ctx context.Context, db *sql.DB, id int) domain.Post {
	//TODO implement me
	sqlQuery := "SELECT id, author_id, title, content, created_at FROM posts WHERE id = $1 LIMIT 1"
	row := db.QueryRowContext(ctx, sqlQuery, id)
	var post domain.Post

	err := row.Scan(&post.Id, &post.AuthorId, &post.Title, &post.Content, &post.CreatedAt)

	if err != nil {
		//helper.PanicIfError(err)
		log.Fatal(err)
	}

	return post
}

func (repository PostRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, post post.CreateRequest, id int) bool {
	//TODO implement me
	currentDate := helper.GetCurrentTime()
	sqlQuery := "UPDATE posts SET title = $1, content = $2, updated_at = $3 WHERE id = $4"
	result, err := tx.ExecContext(ctx, sqlQuery, post.Title, post.Content, currentDate, id)

	helper.PanicIfError(err)

	rows, err := result.RowsAffected()

	helper.PanicIfError(err)

	if rows != 1 {
		log.Fatal("Error")
		return false
	}

	return true
}

func (repository PostRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) bool {
	sqlQuery := "DELETE FROM posts WHERE id = $1"
	result, err := tx.ExecContext(ctx, sqlQuery, id)

	helper.PanicIfError(err)

	rows, err := result.RowsAffected()

	helper.PanicIfError(err)

	if rows != 1 {
		log.Fatal("error")
		return false
	}

	return true
}
