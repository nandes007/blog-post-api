package service

import (
	"fmt"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/web/post"
	"nandes007/blog-post-rest-api/repository"

	"github.com/go-playground/validator/v10"
)

type postServiceImpl struct {
	PostRepository repository.PostRepository
	Validate       *validator.Validate
}

func NewPostService(postRepository repository.PostRepository, validate *validator.Validate) PostService {
	return &postServiceImpl{
		PostRepository: postRepository,
		Validate:       validate,
	}
}

func (s *postServiceImpl) CreatePost(req *post.PostRequest, userID int) (*post.PostResponse, error) {
	err := s.Validate.Struct(req)
	helper.PanicIfError(err)

	post, err := s.PostRepository.Create(req, userID)
	if err != nil {
		fmt.Println("Error when create post : ", err)
		return nil, err
	}
	return post, nil
}

func (s *postServiceImpl) GetAllPosts() ([]*post.PostResponse, error) {
	posts, err := s.PostRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *postServiceImpl) GetPostByID(id int) (*post.PostResponse, error) {
	post, err := s.PostRepository.GetByID(id)
	if err != nil {
		fmt.Println("Error when get post : ", err)
		return nil, err
	}
	return post, nil
}

func (s *postServiceImpl) UpdatePost(req *post.UpdatePostRequest) (*post.PostResponse, error) {
	post, err := s.PostRepository.Update(req)
	if err != nil {
		fmt.Println("Error when update post : ", err)
		return nil, err
	}

	return post, nil
}

func (s *postServiceImpl) DeletePost(id int) error {
	err := s.PostRepository.Delete(id)
	if err != nil {
		fmt.Println("Error when delete post : ", err)
		return err
	}

	return nil
}
