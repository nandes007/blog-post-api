package response

import (
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/post"
)

func ToPostResponse(postDomain domain.Post) post.Response {
	return post.Response{
		Id:        postDomain.Id,
		AuthorId:  postDomain.AuthorId,
		Title:     postDomain.Title,
		Content:   postDomain.Content,
		CreatedAt: postDomain.CreatedAt,
		UpdatedAt: postDomain.UpdatedAt,
	}
}
