package repository

import (
	"context"
	"database/sql"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/domain"
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

func (repository PostRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Post {
	//TODO implement me
	panic("implement me")
}

func (repository PostRepositoryImpl) Find(ctx context.Context, tx *sql.DB, id int) domain.Post {
	//TODO implement me
	panic("implement me")
}

func (repository PostRepositoryImpl) Update(ctx context.Context, tx *sql.DB, post domain.Post, id int) bool {
	//TODO implement me
	panic("implement me")
}

func (repository PostRepositoryImpl) Delete(ctx context.Context, tx *sql.DB, id int) bool {
	//TODO implement me
	panic("implement me")
}
