package service

import (
	"context"
	"fmt"
	"nandes007/blog-post-rest-api/helper/jwt"
	"nandes007/blog-post-rest-api/model/web/post"
	"nandes007/blog-post-rest-api/repository"

	"github.com/go-playground/validator/v10"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewPostService(postRepository repository.PostRepository, userRepository repository.UserRepository, validate *validator.Validate) PostService {
	return &PostServiceImpl{
		PostRepository: postRepository,
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (s *PostServiceImpl) Create(ctx context.Context, req *post.PostRequest, token string) (*post.PostResponse, error) {
	err := s.Validate.Struct(req)
	if err != nil {
		fmt.Println("Error validate : ", err)
		return nil, err
	}

	tokenFormatted := jwt.FormatToken(token)
	user, err := s.UserRepository.Find(ctx, tokenFormatted)
	if err != nil {
		fmt.Println("Error when get user : ", err)
		return nil, err
	}

	post, err := s.PostRepository.Save(ctx, user, req)
	if err != nil {
		fmt.Println("Error when create post : ", err)
		return nil, err
	}
	return post, nil
}

func (s *PostServiceImpl) FindAll(ctx context.Context, token string) ([]*post.PostResponse, error) {
	tokenFormatted := jwt.FormatToken(token)
	user, err := s.UserRepository.Find(ctx, tokenFormatted)
	if err != nil {
		fmt.Println("Error when get user : ", err)
		return nil, err
	}

	posts, err := s.PostRepository.GetAll(ctx, user)
	return posts, nil
}

func (s *PostServiceImpl) Find(ctx context.Context, token string, id int) (*post.PostResponse, error) {
	tokenFormatted := jwt.FormatToken(token)
	user, err := s.UserRepository.Find(ctx, tokenFormatted)
	if err != nil {
		fmt.Println("Error when get user : ", err)
		return nil, err
	}

	post, err := s.PostRepository.Find(ctx, user, id)
	if err != nil {
		fmt.Println("Error when get post : ", err)
	}
	return post, nil
}

func (s *PostServiceImpl) Update(ctx context.Context, req *post.UpdatePostRequest, token string) (*post.PostResponse, error) {
	err := s.Validate.Struct(req)
	if err != nil {
		fmt.Println("Error validate : ", err)
		return nil, err
	}

	tokenFormatted := jwt.FormatToken(token)
	user, err := s.UserRepository.Find(ctx, tokenFormatted)
	if err != nil {
		fmt.Println("Error when get user : ", err)
		return nil, err
	}

	post, err := s.PostRepository.Update(ctx, req, user)
	if err != nil {
		fmt.Println("Error when update post : ", err)
	}

	return post, nil
}

func (s *PostServiceImpl) Delete(ctx context.Context, id int) error {
	err := s.PostRepository.Delete(ctx, id)
	if err != nil {
		fmt.Println("Error when delete post : ", err)
		return err
	}

	return nil
}
