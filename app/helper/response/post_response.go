package response

import (
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/post"
)

func ToPostResponse(postDomain domain.Post) post.PostResponse {
	return post.PostResponse{
		ID:        postDomain.ID,
		Title:     postDomain.Title,
		Content:   postDomain.Content,
		CreatedAt: postDomain.CreatedAt,
		UpdatedAt: postDomain.UpdatedAt,
	}
}

func ToPostsResponse(posts []domain.Post) []post.PostResponse {
	var postsResponse []post.PostResponse
	for _, post := range posts {
		postsResponse = append(postsResponse, ToPostResponse(post))
	}

	return postsResponse
}
