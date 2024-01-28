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
		User:      postDomain.User,
	}
}

func ToPostsResponse(posts []domain.Post) []post.Response {
	var postsResponse []post.Response
	for _, post := range posts {
		postsResponse = append(postsResponse, ToPostResponse(post))
	}

	return postsResponse
}
